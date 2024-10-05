package app

import (
	"database/sql"

	"github.com/Insid1/go-auth-user/auth-service/internal/config"
	"github.com/Insid1/go-auth-user/auth-service/internal/handler"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"
	userPkg "github.com/Insid1/go-auth-user/user/pkg"
)

type Provider struct {
	config         *config.Config
	db             *sql.DB
	grpcUserClient *userPkg.GRPCInitializedUserClient

	authHandler    handler.Auth
	authService    service.Auth
	authRepository repository.Auth

	userRepository repository.User
}

func newProvider(
	config *config.Config,
	db *sql.DB,
	grpcUserClient *userPkg.GRPCInitializedUserClient,
) (*Provider, error) {
	return &Provider{config: config, db: db, grpcUserClient: grpcUserClient}, nil
}

func (p *Provider) AuthHandler() handler.Auth {
	if p.authHandler == nil {
		p.authHandler = handler.NewAuthHandler(p.AuthService())
	}

	return p.authHandler
}

func (p *Provider) AuthService() service.Auth {
	if p.authService == nil {
		p.authService = service.NewAuthService(p.config.JwtSecretKey, p.UserRepository(), p.AuthRepository())
	}

	return p.authService
}

func (p *Provider) UserRepository() repository.User {
	if p.userRepository == nil {
		p.userRepository = repository.NewUserRepository(p.grpcUserClient)
	}

	return p.userRepository
}

func (p *Provider) AuthRepository() repository.Auth {
	if p.authRepository == nil {
		p.authRepository = repository.NewAuthRepository(p.db)
	}

	return p.authRepository
}
