package app

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/user-service/internal/config"
	"github.com/Insid1/go-auth-user/user-service/internal/handler/user"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
)

type Provider struct {
	ctx    context.Context
	config *config.Config
	db     *sql.DB

	userHandler *user.Handler
	service     *service.Service
	repository  *repository.Repository
}

func newProvider(
	ctx context.Context,
	config *config.Config,
	db *sql.DB,
) *Provider {
	return &Provider{ctx: ctx, config: config, db: db}
}

func (p *Provider) UserHandler() *user.Handler {
	if p.userHandler == nil {
		p.userHandler = user.NewHandler(p.Service())
	}

	return p.userHandler
}

func (p *Provider) Service() *service.Service {
	if p.service == nil {
		p.service = service.NewService(p.Repository())
	}

	return p.service
}

func (p *Provider) Repository() *repository.Repository {
	if p.repository == nil {
		p.repository = repository.NewRepository(p.db)
	}

	return p.repository
}
