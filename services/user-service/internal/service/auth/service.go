package auth

import (
	"time"

	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/internal/repository"
	"github.com/golang-jwt/jwt"
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

	token, err := s.generateToken(usr.ID, usr.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) Register(register *model.Register) (string, error) {
	return "its actually some user new id", nil
}

func (s *Service) Logout(token string) (bool, error) {
	return true, nil
}

func (s *Service) generateToken(id string, email string) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := jwt.MapClaims{
		"sub":   id,
		"email": email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey))
}
