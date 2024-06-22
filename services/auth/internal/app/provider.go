package app

import (
	"context"
	"database/sql"

	"github.com/Insid1/go-auth-user/auth-service/internal/common"
	"github.com/Insid1/go-auth-user/auth-service/internal/config"
	"github.com/Insid1/go-auth-user/auth-service/internal/handler"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"

	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
)

type Provider struct {
	ctx            context.Context
	config         *config.Config
	db             *sql.DB
	grpcUserClient *common.GRPCClient[user_v1.UserV1Client]

	authHandler handler.Auth
	authService service.Auth

	userRepository repository.User
}

func newProvider(
	ctx context.Context,
	config *config.Config,
	db *sql.DB,
	grpcUserClient *common.GRPCClient[user_v1.UserV1Client],
) (*Provider, error) {
	return &Provider{ctx: ctx, config: config, db: db, grpcUserClient: grpcUserClient}, nil
}

func (p *Provider) AuthHandler() handler.Auth {
	if p.authHandler == nil {
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

func (p *Provider) UserRepository() repository.User {
	if p.userRepository == nil {
		p.userRepository = repository.NewUserRepository(p.ctx, p.grpcUserClient)
	}

	return p.userRepository
}
