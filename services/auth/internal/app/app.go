package app

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"
	"github.com/Insid1/go-auth-user/auth-service/internal/config"
	"github.com/Insid1/go-auth-user/auth-service/pkg/auth_v1"

	"github.com/Insid1/go-auth-user/user/pkg/user_v1"

	"github.com/Insid1/go-auth-user/pkg/utils"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

type App struct {
	config         *config.Config
	DB             *sql.DB
	Logger         *zap.SugaredLogger
	grpcServer     *grpc.Server
	grpcUserClient *common.GRPCClient[user_v1.UserV1Client]
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

	db, err := sql.Open("postgres", a.config.Db.DataSourceName())
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
	provider, err := newProvider(ctx, a.config, a.DB, a.grpcUserClient)
	if err != nil {
		return err
	}

	a.provider = provider
	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)
	auth_v1.RegisterAuthV1Server(a.grpcServer, a.provider.AuthHandler())

	return nil
}

func (a *App) initGRPCUserClient(ctx context.Context) error {
	// Загрузка TLS сертификата
	creds, err := utils.LoadTLSCredentials(&utils.CredentialParams{
		CAClientCertPath: a.config.Security.CAClientCertPath,
		CAServerCertPath: a.config.Security.CAServerCertPath,
		CertPath:         a.config.Security.ClientCertPath,
		CertKeyPath:      a.config.Security.ClientKeyPath,
	})
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to load TLS credentials. %s", err))
	}

	connection, err := grpc.NewClient(a.config.Global.Service.User.Server.Address(), grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}

	a.grpcUserClient = &common.GRPCClient[user_v1.UserV1Client]{
		Client:     user_v1.NewUserV1Client(connection),
		Connection: connection,
	}
	return nil
}

func (a *App) runGRPCServer() error {
	a.Logger.Infof("GRPC Auth server is running on %s", a.config.Global.Service.Auth.Server.Address())

	list, err := net.Listen("tcp", a.config.Global.Service.Auth.Server.Address())
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to listen GRPC Auth server. %s", err))
	}

	err = a.grpcServer.Serve(list)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to serve GRPC Auth server. %s", err))
	}

	return nil
}
