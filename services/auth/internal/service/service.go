package service

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository"
	"github.com/Insid1/go-auth-user/auth-service/internal/service/auth"
)

type Auth interface {
	Login(*model.Login) (*auth.TokenPair, error)
	Register(*model.Register) (string, error)
	Logout(string) (bool, error)
}

func NewAuthService(ctx context.Context, JWTKey string, userRepo repository.User, authRepo repository.Auth) Auth {
	return &auth.Service{
		Ctx:    ctx,
		JWTKey: JWTKey,

		UserRepository: userRepo,
		AuthRepository: authRepo,
	}
}
