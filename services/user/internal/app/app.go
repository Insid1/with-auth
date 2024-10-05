package app

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"strings"

	authPkg "github.com/Insid1/go-auth-user/auth/pkg"
	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"
	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
	"github.com/Insid1/go-auth-user/pkg/utils"
	"github.com/Insid1/go-auth-user/user/internal/config"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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
		a.initGRPCServer,
		a.initGRPCAuthClient,
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
		utils.GetUnaryPanicInterceptor(a.Logger),
		a.getAuthUnaryInterceptor(a.getAuthMethodNames()),
		a.getUnaryLoggingInterceptor(),
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
// Для логирования данных из запросов и ответов
func (a *App) getUnaryLoggingInterceptor() grpc.UnaryServerInterceptor {

	loggingOpts := []logging.Option{
		logging.WithLogOnEvents(
			logging.PayloadReceived,
			logging.PayloadSent,
		),
	}

	return logging.UnaryServerInterceptor(a.getInterceptorLogger(), loggingOpts...)
}

func (a *App) getInterceptorLogger() logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		a.Logger.Log(zap.InfoLevel, msg, fields)
	})
}

func (a *App) getAuthUnaryInterceptor(methodsName []string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {

		// check is method in methodsName
		if !utils.IsInList(info.FullMethod, methodsName) {
			return handler(ctx, req)
		}

		// Извлечение headers из запроса данных
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return "", status.Errorf(codes.Unauthenticated, "error retrieving headers")
		}

		// Проверка, что сервис имеет доступ к данным
		err = checkAuthService(md)
		if err == nil {
			return handler(ctx, req)
		}

		accessToken, err := retrieveAccessToken(md)
		if err != nil {
			return "", status.Errorf(codes.Unauthenticated, "error retrieving token")
		}

		_, err = a.authClient.Client.Check(ctx, &auth_v1.CheckReq{
			AccessToken: accessToken,
		})

		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "error validating token")
		} else {
			return handler(ctx, req)
		}
	}
}

func retrieveAccessToken(md metadata.MD) (string, error) {
	// Извлечение авторизационных данных
	authDataList, ok := md["authorization"]
	if !ok {
		return "", fmt.Errorf("unable to retrieve auth data")
	}

	// Извлечение токена
	var accessToken string
	prefix := "Bearer "
	for _, authData := range authDataList {
		accessToken, ok = strings.CutPrefix(authData, prefix)
		if !ok {
			return "", fmt.Errorf("unable to retrieve auth data")
		}
	}

	return accessToken, nil
}

func checkAuthService(md metadata.MD) error {
	// Извлечение авторизационных данных
	serviceDataList, ok := md["service-name"]
	if !ok {
		return fmt.Errorf("unable to retrieve service info")
	}

	// Извлечение токена
	for _, serviceName := range serviceDataList {
		if serviceName == "auth" {
			return nil
		}
	}

	return fmt.Errorf("not auth service")
}

// Отдает список имен методов, для которых необходима проверка токена авторизации
func (a *App) getAuthMethodNames() []string {
	return []string{
		user_v1.UserV1_Get_FullMethodName,
	}
}
