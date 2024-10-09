package app

import (
	"context"
	"errors"
	"net"
	"net/http"
	"time"

	"github.com/Insid1/with-auth/url-shortener/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"go.uber.org/zap"
)

type App struct {
	config *config.Config
	DB     *mongo.Client
	Logger *zap.SugaredLogger

	httpServer *http.Server

	provider *Provider
}

const (
	dbTimeout    = 2 * time.Second
	readTimeout  = 12 * time.Second
	writeTimeout = 12 * time.Second
	idleTimeout  = 60 * time.Second
	stopTimeout  = 20 * time.Second
)

func NewApp(ctx context.Context) *App {
	a := App{
		config:     nil,
		DB:         nil,
		Logger:     nil,
		httpServer: nil,
		provider:   nil,
	}

	err := a.initDeps(ctx)
	if err != nil {
		panic(err)
	}

	return &a
}

func (a *App) Run() error {
	runFunctions := []func() error{
		a.runHTTPServer,
	}

	for _, fn := range runFunctions {
		err := fn()
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) Stop(ctx context.Context) error {
	if err := a.stopHTTPServer(ctx); err != nil {
		return err
	}

	if err := a.DB.Disconnect(ctx); err != nil {
		return err
	}

	// Игнорируем т.к. всегда возвращает ошибку
	_ = a.Logger.Sync()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	arr := []func(ctx context.Context) error{
		a.initConfig,
		a.initLogger,
		a.initDataBaseConnection,
		a.initProvider,
		a.initHTTPServer,
	}

	for _, fn := range arr {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	a.config = config.MustLoad()

	return nil
}

func (a *App) initLogger(_ context.Context) error {
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
	// Установка соединения
	client, err := mongo.Connect(
		options.Client().ApplyURI(
			a.config.GetMongoDataBaseURL(),
		))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, dbTimeout)
	// Чистим ресурсы
	defer cancel()

	// Попытка подключения
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}

	a.Logger.Info("Successfully connected to Mongo DB")

	a.DB = client

	return nil
}

func (a *App) initProvider(_ context.Context) error {
	a.provider = newProvider(a.config, a.DB, a.Logger)

	return nil
}

func (a *App) initHTTPServer(_ context.Context) error {
	server := &http.Server{
		Addr:                         a.config.GetAppAddress(),
		Handler:                      a.provider.getRouter(),
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  readTimeout,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 writeTimeout,
		IdleTimeout:                  idleTimeout,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
	}

	a.httpServer = server

	return nil
}

func (a *App) runHTTPServer() error {
	lis, err := net.Listen("tcp", a.httpServer.Addr)
	if err != nil {
		return err
	}

	// Run the server
	go func() {
		_ = a.httpServer.Serve(lis)
	}()

	a.Logger.Infof("HTTP url-shortener server is running on %s", a.config.GetAppAddress())

	return nil
}

func (a *App) stopHTTPServer(ctx context.Context) error {
	// Shutdown signal with grace period time
	shutdownCtx, cancel := context.WithTimeout(ctx, stopTimeout)
	defer cancel()

	go func() {
		<-shutdownCtx.Done()

		if errors.Is(shutdownCtx.Err(), context.DeadlineExceeded) {
			a.Logger.Info("graceful shutdown timed out.. forcing exit")
		}
	}()

	// Trigger graceful shutdown
	return a.httpServer.Shutdown(shutdownCtx)
}
