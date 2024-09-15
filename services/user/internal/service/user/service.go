package user

import (
	"errors"

	"github.com/Insid1/go-auth-user/user/internal/model"
	"github.com/Insid1/go-auth-user/user/internal/repository"
)

type Service struct {
	UserRepository repository.User
}

func (s *Service) Get(id string, email string) (*model.User, error) {
	if len(id) > 0 {
		return s.UserRepository.Get(id)
	}

	if len(email) > 0 {
		return s.UserRepository.GetBy("email", email)
	}

	return nil, errors.New("user not found") // actually no data provided for correct req
}

func (s *Service) Create(usr *model.User) (*model.User, error) {
	return s.UserRepository.Create(usr)
}

func (s *Service) Update(usr *model.User) (*model.User, error) {
	return s.UserRepository.Update(usr)
}

func (s *Service) Delete(id string) (string, error) {
	return s.UserRepository.Delete(id)
}
