package app

import (
	"context"
	"database/sql"
	"fmt"
	"net"

	authPkg "github.com/Insid1/with-auth/auth/pkg"
	"github.com/Insid1/with-auth/pkg/grpc/user_v1"
	serverInterceptors "github.com/Insid1/with-auth/pkg/interceptors/server"
	"github.com/Insid1/with-auth/user/internal/config"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config *config.Config
	DB     *sql.DB
	Logger *zap.SugaredLogger

	grpcServer *grpc.Server
	authClient *authPkg.GRPCInitializedAuthClient
	provider   *Provider
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	err := a.runGRPCServer()
	if err != nil {
		return err
	}
	return nil
}

func (a *App) Stop() error {
	a.grpcServer.GracefulStop()

	if err := a.DB.Close(); err != nil {
		return err
	}

	a.Logger.Sync()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	arr := []func(ctx context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initDataBaseConnection,
		a.initProvider,
		a.initGRPCAuthClient,
		a.initGRPCServer,
	}

	for _, fn := range arr {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(ctx context.Context) error {
	a.config = config.MustLoad()
	return nil

}

func (a *App) initLogger(ctx context.Context) error {
	var logger *zap.Logger

	// todo Требует дополнительной доработки по необходимости
	if a.config.Env == "prod" {
		logger = zap.Must(zap.NewProduction())
	} else {
		logger = zap.Must(zap.NewDevelopment())
	}

	sugarLogger := logger.Sugar()
	a.Logger = sugarLogger

	return nil
}

func (a *App) initDataBaseConnection(ctx context.Context) error {

	db, err := sql.Open("postgres", a.config.GetDataBaseURL())
	if err != nil {
		return fmt.Errorf("unable to Open DB connection. %s", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("unable to connect to DB. %s", err)
	}

	a.Logger.Info("Connected to DataBase")
	a.DB = db

	return nil
}

func (a *App) initProvider(ctx context.Context) error {
	a.provider = newProvider(a.config, a.DB)
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {

	a.grpcServer = grpc.NewServer(grpc.ChainUnaryInterceptor(
		serverInterceptors.UnaryPanicInterceptor(a.Logger),
		serverInterceptors.UnaryLoggingInterceptor(a.Logger),
		authPkg.AuthUnaryInterceptor(a.authClient.Client, a.getAuthMethodNames()),
	))

	reflection.Register(a.grpcServer)

	user_v1.RegisterUserV1Server(a.grpcServer, a.provider.UserHandler())

	return nil
}

func (a *App) initGRPCAuthClient(ctx context.Context) error {
	println(a.config.GetAuthServiceAddress())
	client, err := authPkg.InitGRPCAuthClient(ctx, &authPkg.GRPCAuthClientConfig{
		ServerAddress:     a.config.GetAuthServiceAddress(),
		ClientServiceName: "user",
	})

	if err != nil {
		return err
	}

	a.authClient = client
	return nil
}

func (a *App) runGRPCServer() error {
	a.Logger.Infof("GRPC user server is running on %s", a.config.GetAppAddress())

	list, err := net.Listen("tcp", a.config.GetAppAddress())
	if err != nil {
		return fmt.Errorf("unable to listen GRPC user server. %s", err)
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return fmt.Errorf("unable to serve GRPC user server. %s", err)
	}

	return nil
}

// todo ПЕРЕНЕСИ интерсепторы в pkg

// Отдает список имен методов, для которых необходима проверка токена авторизации
func (a *App) getAuthMethodNames() []string {
	return []string{
		user_v1.UserV1_Get_FullMethodName,
	}
}
