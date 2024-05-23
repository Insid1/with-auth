package user

import (
	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
)

type Service struct {
	UserRepository repository.User
}

func (s *Service) Get(id string) (*model.User, error) {
	return s.UserRepository.Get(id)
}

func (s *Service) Create(usr *model.User) (string, error) {
	return s.UserRepository.Create(usr)
}
