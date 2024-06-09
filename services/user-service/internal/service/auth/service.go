package auth

import (
	"time"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	UserRepository repository.User
	JWTKey         string
}

func (s *Service) Login(data *model.Login) (string, error) {
	usr, err := s.UserRepository.GetBy("email", data.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(data.Password))
	if err != nil {
		return "", err
	}

	token, err := s.generateToken(usr.ID, usr.Email, usr.Password)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Register(data *model.Register) (string, error) {
	_, err := s.UserRepository.GetBy("email", data.Email)
	if err == nil {
		return "", errors.New("email already in use")
	}

	passHash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	userID, err := s.UserRepository.Create(&model.User{
		Email:    data.Email,
		Password: string(passHash),
	})

	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *Service) Logout(token string) (bool, error) {
	// invalidate refresh token ? delete from DB when it will be
	return true, nil
}

func (s *Service) generateToken(id string, email string, passHash string) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := jwt.MapClaims{
		"sub":   id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 3).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey + passHash))
}
