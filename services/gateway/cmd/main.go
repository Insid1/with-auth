package main

import (
	"context"
	"log"

	"github.com/Insid1/go-auth-user/gateway/internal/app"
)

func main() {

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("failed to migrations app: %s", err.Error())
	}

	err = a.Run("8801")
	if err != nil {
		log.Fatalf("failed to run app: %s", err.Error())
	}
}
