package user

import (
	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
)

type Service struct {
	repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{repository: repo}
}

func (s *Service) Get(id string) (*model.User, error) {
	return s.repository.UserRepository.Get(id)
}

func (s *Service) Create(usr *model.User) (string, error) {
	return s.repository.UserRepository.Create(usr)
}
