package main

import (
	"context"
	"log"

	"github.com/Insid1/go-auth-user/user-service/internal/app"
	"github.com/Insid1/go-auth-user/user-service/internal/config"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	config.MustLoad()

	application, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	defer application.Stop()

	if application.Run() != nil {
		application.Logger.Errorf("Failed to run application: %v", err)
	}
}
