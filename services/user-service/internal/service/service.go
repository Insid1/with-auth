package service

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/Insid1/go-auth-user/user-service/internal/service/user"
)

type User interface {
	Get(id string, email string) (*model.User, error)
	Create(*model.User) (string, error)
}

func NewUserService(ctx context.Context, repo repository.User) User {
	return &user.Service{
		UserRepository: repo,
	}
}
