package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Insid1/with-auth/url-shortener/internal/app"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	a := app.NewApp(ctx)

	if err := a.Run(); err != nil {
		a.Logger.Errorf("Failed to run application: %v", err)
	} else {
		a.Logger.Info("Application started")
	}

	gracefulShutdown(ctx, a)
}

func gracefulShutdown(ctx context.Context, application *app.App) {
	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop

	application.Logger.Info("Stopping application...", zap.String("signal", sign.String()))
	err := application.Stop(ctx)

	if err != nil {
		application.Logger.Error("Failed to stop application", zap.Error(err))
	} else {
		application.Logger.Info("Application stopped")
	}
}
