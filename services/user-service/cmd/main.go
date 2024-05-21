package main

import (
	"context"
	"log"

	"github.com/Insid1/go-auth-user/user-service/internal/app"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()
	application, err := app.NewApp(ctx)

	if err != nil {
		log.Fatalf("Failed to initialize application: %v", err)
	}

	if application.Run() != nil {
		log.Fatalf("Failed to run application: %v", err)
	}

	defer application.DB.Close()
}
