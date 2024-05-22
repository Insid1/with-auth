package service

import (
	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/Insid1/go-auth-user/user-service/internal/service/user"
)

type Service struct {
	UserService
}

type UserService interface {
	Get(string) (*model.User, error)
	Create(*model.User) (string, error)
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: user.NewService(repo),
	}
}
