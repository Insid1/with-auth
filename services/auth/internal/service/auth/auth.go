package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/model"
)

type Service struct {
	Ctx context.Context
}

func (s *Service) Login(*model.Login) (string, error) {
	return "", nil
}
func (s *Service) Register(*model.Register) (string, error) {
	return "", nil
}
func (s *Service) Logout(string) (bool, error) {
	return false, nil
}
