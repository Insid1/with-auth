package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Insid1/with-auth/user/internal/app"
	"github.com/Insid1/with-auth/user/internal/config"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	config.MustLoad()

	application, err := app.NewApp(ctx)

	if err != nil {
		panic(fmt.Sprintf("Failed to initialize application: %v", err))
	}

	go func() {
		if application.Run() != nil {
			application.Logger.Errorf("Failed to run application: %v", err)
		}
	}()

	gracefulShutdown(application)
}

func gracefulShutdown(application *app.App) {
	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	application.Logger.Info("Stopping application...", zap.String("signal", sign.String()))
	err := application.Stop()

	if err != nil {
		application.Logger.Error("Failed to stop application", zap.Error(err))
	} else {
		application.Logger.Info("Application stopped")
	}
}
