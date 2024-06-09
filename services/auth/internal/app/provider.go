package app

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/auth-service/internal/config"
	"github.com/Insid1/go-auth-user/auth-service/internal/handler"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"
)

type Provider struct {
	ctx    context.Context
	config *config.Config
	db     *sql.DB

	// userHandler    handler.User
	// userService    service.User
	// userRepository repository.User

	authHandler handler.Auth
	authService service.Auth
}

func newProvider(ctx context.Context,
	config *config.Config,
	db *sql.DB) (*Provider, error) {
	return &Provider{ctx: ctx, config: config, db: db}, nil
}

func (p *Provider) AuthHandler() handler.Auth {
	if p.authHandler == nil {
		p.authHandler = handler.NewAuthHandler(p.ctx)
	}

	return p.authHandler
}
