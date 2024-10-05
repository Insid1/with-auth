package user

import (
	userErrors "github.com/Insid1/with-auth/pkg/errors/user"
	"github.com/Insid1/with-auth/user/internal/model"
	"github.com/Insid1/with-auth/user/internal/repository"
	"golang.org/x/crypto/bcrypt"
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

	return nil, userErrors.ErrUserNotFound
}

func (s *Service) Create(usr *model.User, password string) (*model.User, error) {
	err := usr.UpdatePassHash(password)
	if err != nil {
		return nil, err
	}

	return s.UserRepository.Create(usr)
}

func (s *Service) Update(usr *model.User, password string) (*model.User, error) {
	err := usr.UpdatePassHash(password)
	if err != nil {
		return nil, err
	}

	return s.UserRepository.Update(usr)
}

func (s *Service) CheckPassword(
	userID string,
	email, password string,
) (*model.User, error) {
	var usr *model.User

	var err error

	if userID != "" {
		usr, err = s.UserRepository.GetBy("id", userID)
	}

	if usr == nil && email != "" {
		usr, err = s.UserRepository.GetBy("email", email)
	}

	if err != nil {
		return nil, err
	}

	err = s.CheckPasswordHash(usr.PassHash, password)
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func (s *Service) Delete(id string) (string, error) {
	return s.UserRepository.Delete(id)
}

func (s *Service) CheckPasswordHash(passwordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	if err != nil {
		return userErrors.ErrInvalidPassword
	}

	return nil
}
