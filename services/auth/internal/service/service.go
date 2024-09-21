package service

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository"
	"github.com/Insid1/go-auth-user/auth-service/internal/service/auth"
	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"
)

type Auth interface {
	Login(ctx context.Context, lgn *model.Login) (*auth.TokenPair, error)
	Register(ctx context.Context, rgst *model.Register) (*user_v1.User, error)
	Logout(ctx context.Context, refreshToken string) (bool, error)
	CheckTokens(ctx context.Context, tokens *model.Check) (*model.Check, error)
}

func NewAuthService(JWTKey string, userRepo repository.User, authRepo repository.Auth) Auth {
	return &auth.Service{
		JWTKey:         JWTKey,
		UserRepository: userRepo,
		AuthRepository: authRepo,
	}
}
