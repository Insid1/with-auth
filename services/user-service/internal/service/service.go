package service

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/Insid1/go-auth-user/user-service/internal/service/auth"
	"github.com/Insid1/go-auth-user/user-service/internal/service/user"
)

type User interface {
	Get(string) (*model.User, error)
	Create(*model.User) (string, error)
}

type Auth interface {
	Login(*model.Login) (string, error)
	Register(*model.Register) (string, error)
	Logout(string) (bool, error)
}

func NewUserService(ctx context.Context, repo repository.User) User {
	return &user.Service{
		UserRepository: repo,
	}
}

func NewAuthService(ctx context.Context, JWTKey string, repo repository.User) Auth {
	return &auth.Service{
		UserRepository: repo,
		JWTKey:         JWTKey,
	}
}
