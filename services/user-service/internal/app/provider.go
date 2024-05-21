package app

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/handler/user"
)

type Provider struct {
	userHandler *user.Handler
}

func newProvider(ctx context.Context) *Provider {
	return &Provider{}
}

func (p *Provider) UserHandler() *user.Handler {
	if p.userHandler == nil {
		p.userHandler = user.NewHandler()
	}

	return p.userHandler
}
