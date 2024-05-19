package app

import (
	"context"
	"fmt"

	"github.com/Insid1/go-auth-user/gateway/internal/api"
	"github.com/gin-gonic/gin"
)

type App struct {
	provider   *Provider
	httpEngine *gin.Engine
}

func NewApp(ctx context.Context) (*App, error) {
	app := &App{}
	err := app.initDeps(ctx)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(port string) error {
	return a.httpEngine.Run(fmt.Sprintf(":%s", port))
}

func (a *App) initDeps(ctx context.Context) error {
	var initArr = []func(context.Context) error{a.initProvider, a.initHttpServer}

	for _, fn := range initArr {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initProvider(_ context.Context) error {
	a.provider = NewProvider()
	return nil
}

func (a *App) initHttpServer(_ context.Context) error {
	engine := gin.Default()
	a.httpEngine = engine

	api.UseRoutes(a.httpEngine, a.provider.UserHandler())

	return nil
}
