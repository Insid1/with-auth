package service

import (
	"context"

	"github.com/Insid1/with-auth/auth-service/internal/model"
	"github.com/Insid1/with-auth/auth-service/internal/repository"
	"github.com/Insid1/with-auth/auth-service/internal/service/auth"
	"github.com/Insid1/with-auth/pkg/grpc/user_v1"
)

type Auth interface {
	Login(ctx context.Context, data *model.Login) (*auth.TokenPair, error)
	Register(ctx context.Context, data *model.Register) (*user_v1.User, error)
	LogoutAll(ctx context.Context, userId string) error
	GenerateTokenPair(ctx context.Context, userId string, email string) (*auth.TokenPair, error)
	CheckAccessToken(ctx context.Context, accessToken string) (*auth.AccessTokenClaims, error)
	CheckRefreshToken(ctx context.Context, refreshToken string) (*auth.RefreshTokenClaims, error)
}

func NewAuthService(JWTKey string, userRepo repository.User, authRepo repository.Auth) Auth {
	return &auth.Service{
		JWTKey:         JWTKey,
		UserRepository: userRepo,
		AuthRepository: authRepo,
	}
}
