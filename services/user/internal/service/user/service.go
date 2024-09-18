package user

import (
	"errors"
	"fmt"

	"github.com/Insid1/go-auth-user/user/internal/model"
	"github.com/Insid1/go-auth-user/user/internal/repository"
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

	return nil, errors.New("user not found") // actually no data provided for correct req
}

func (s *Service) Create(usr *model.User, password string) (*model.User, error) {

	err := usr.UpdatePassHash(password)

	if err != nil {
		return nil, err
	}

	return s.UserRepository.Create(usr)
}

func (s *Service) Update(usr *model.User, password string) (*model.User, error) {

	usr.UpdatePassHash(password)

	return s.UserRepository.Update(usr)
}

func (s *Service) CheckPassword(id string, email, password string) (*model.User, error) {
	var usr *model.User
	var err error

	if id != "" {
		usr, err = s.UserRepository.GetBy("id", id)
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
		return fmt.Errorf("error: Password is invalid. %s", err)
	}

	return nil
}
