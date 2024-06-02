package app

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/user-service/internal/config"
	"github.com/Insid1/go-auth-user/user-service/internal/handler"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
)

type Provider struct {
	ctx    context.Context
	config *config.Config
	db     *sql.DB

	userHandler    handler.User
	userService    service.User
	userRepository repository.User

	authHandler handler.Auth
	authService service.Auth
}

func newProvider(
	ctx context.Context,
	config *config.Config,
	db *sql.DB,
) *Provider {
	return &Provider{ctx: ctx, config: config, db: db}
}

func (p *Provider) UserHandler() handler.User {
	if p.userHandler == nil {
		p.userHandler = handler.NewUserHandler(p.ctx, p.UserService())
	}

	return p.userHandler
}

func (p *Provider) UserService() service.User {
	if p.userService == nil {
		p.userService = service.NewUserService(p.ctx, p.UserRepository())
	}

	return p.userService
}

func (p *Provider) UserRepository() repository.User {
	if p.userRepository == nil {
		p.userRepository = repository.NewUserRepository(p.ctx, p.db)
	}

	return p.userRepository
}

func (p *Provider) AuthHandler() handler.Auth {
	if p.userHandler == nil {
		p.authHandler = handler.NewAuthHandler(p.ctx, p.AuthService())
	}

	return p.authHandler
}

func (p *Provider) AuthService() service.Auth {
	if p.authService == nil {
		p.authService = service.NewAuthService(p.ctx, p.config.JWTKey, p.UserRepository())
	}

	return p.authService
}
