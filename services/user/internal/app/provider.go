package app

import (
	"database/sql"

	"github.com/Insid1/go-auth-user/user/internal/config"
	"github.com/Insid1/go-auth-user/user/internal/handler"
	"github.com/Insid1/go-auth-user/user/internal/repository"
	"github.com/Insid1/go-auth-user/user/internal/service"
)

type Provider struct {
	config *config.Config
	db     *sql.DB

	userHandler    handler.User
	userService    service.User
	userRepository repository.User
}

func newProvider(
	config *config.Config,
	db *sql.DB,
) *Provider {
	return &Provider{config: config, db: db}
}

func (p *Provider) UserHandler() handler.User {
	if p.userHandler == nil {
		p.userHandler = handler.NewUserHandler(p.UserService())
	}

	return p.userHandler
}

func (p *Provider) UserService() service.User {
	if p.userService == nil {
		p.userService = service.NewUserService(p.UserRepository())
	}

	return p.userService
}

func (p *Provider) UserRepository() repository.User {
	if p.userRepository == nil {
		p.userRepository = repository.NewUserRepository(p.db)
	}

	return p.userRepository
}
