package app

import (
	"github.com/Insid1/go-auth-user/gateway/internal/handler"
	"github.com/Insid1/go-auth-user/gateway/internal/repository"
	"github.com/Insid1/go-auth-user/gateway/internal/service"
)

type Provider struct {
	UserProvider
}

type UserProvider struct {
	userHandler handler.User
	userService service.User
	userRepo    repository.User
}

func NewProvider() *Provider {
	return &Provider{}
}

func (p *Provider) UserRepo() repository.User {
	if p.userRepo == nil {
		p.userRepo = repository.NewRepository()
	}
	return p.userRepo
}

func (p *Provider) UserService() service.User {
	if p.userService == nil {
		p.userService = service.NewUserService()
	}
	return p.userService
}

func (p *Provider) UserHandler() handler.User {
	if p.userHandler == nil {
		p.userHandler = handler.NewUserHandler()
	}
	return p.userHandler
}
