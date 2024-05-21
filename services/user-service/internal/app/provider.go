package app

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/handler/user"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
)

type Provider struct {
	userHandler *user.Handler
	service     *service.Service
}

func newProvider(ctx context.Context) *Provider {
	return &Provider{}
}

func (p *Provider) UserHandler() *user.Handler {
	if p.userHandler == nil {
		p.userHandler = user.NewHandler(p.Service())
	}

	return p.userHandler
}

func (p *Provider) Service() *service.Service {
	if p.service == nil {
		p.service = service.NewService()
	}

	return p.service
}
