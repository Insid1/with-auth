package service

import (
	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/service/user"
)

type Service struct {
	UserService
}

type UserService interface {
	Get(string) *model.User
}

func NewService() *Service {
	return &Service{
		UserService: user.NewService(),
	}
}
