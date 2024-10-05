package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"

	"github.com/Insid1/go-auth-user/auth-service/internal/config"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"

	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"
	userPkg "github.com/Insid1/go-auth-user/user/pkg"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type App struct {
	config         *config.Config
	DB             *sql.DB
	Logger         *zap.SugaredLogger
	grpcServer     *grpc.Server
	grpcUserClient *userPkg.GRPCInitializedUserClient
	provider       *Provider
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
	if err := a.grpcUserClient.Connection.Close(); err != nil {
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
		a.initGRPCUserClient,
		a.initProvider,
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
		return errors.New(fmt.Sprintf("Unable to Open DB Connection. %s", err))

	}
	err = db.Ping()
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to connect to DB. %s", err))
	}

	a.Logger.Info("Connected to DataBase")
	a.DB = db

	return nil
}

func (a *App) initProvider(ctx context.Context) error {
	provider, err := newProvider(a.config, a.DB, a.grpcUserClient)
	if err != nil {
		return err
	}

	a.provider = provider
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {

	// Логирование паники
	recoveryOpts := []recovery.Option{
		recovery.WithRecoveryHandler(func(p interface{}) (err error) {
			// Логируем информацию о панике с уровнем Error
			a.Logger.Errorf("Recovered from panic. panic: %s", p)

			// При паники возвращает internal error пользователю
			return status.Errorf(codes.Internal, "internal error")
		}),
	}

	a.grpcServer = grpc.NewServer(grpc.ChainUnaryInterceptor(
		// Для обработки паники внутри запросов к серверу
		recovery.UnaryServerInterceptor(recoveryOpts...),
	))

	reflection.Register(a.grpcServer)

	auth_v1.RegisterAuthV1Server(a.grpcServer, a.provider.AuthHandler())

	return nil
}

func (a *App) initGRPCUserClient(ctx context.Context) error {
	client, err := userPkg.InitGRPCUserClient(ctx, &userPkg.GRPCUserClientConfig{
		ServerAddress:     a.config.GetUserServiceAddress(),
		ClientServiceName: "auth",
	})
	if err != nil {
		return err
	}

	a.grpcUserClient = client

	return nil
}

func (a *App) runGRPCServer() error {
	a.Logger.Infof("GRPC Auth server is running on %s", a.config.GetAppAddress())

	list, err := net.Listen("tcp", a.config.GetAppAddress())
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to listen GRPC Auth server. %s", err))
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to serve GRPC Auth server. %s", err))
	}

	return nil
}
