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

	userHandler handler.User
	service     service.User
	repository  repository.User

	authHandler handler.Auth
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
		p.userHandler = handler.NewUserHandler(p.ctx, p.Service())
	}

	return p.userHandler
}

func (p *Provider) Service() service.User {
	if p.service == nil {
		p.service = service.NewUserService(p.ctx, p.Repository())
	}

	return p.service
}

func (p *Provider) Repository() repository.User {
	if p.repository == nil {
		p.repository = repository.NewUserRepository(p.ctx, p.db)
	}

	return p.repository
}

func (p *Provider) AuthHandler() handler.Auth {
	if p.userHandler == nil {
		p.authHandler = handler.NewAuthHandler(p.ctx)
	}

	return p.authHandler
}
