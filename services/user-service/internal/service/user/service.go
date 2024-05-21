package user

import "github.com/Insid1/go-auth-user/user-service/internal/model"

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Get(id string) *model.User {
	return &model.User{
		ID:    id,
		Name:  "",
		Email: "",
		Age:   0,
	}
}
