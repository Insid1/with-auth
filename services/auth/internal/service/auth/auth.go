package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/internal/repository"
	"github.com/Insid1/go-auth-user/pkg/grpc/user_v1"

	"github.com/golang-jwt/jwt"
)

type RefreshTokenClaims struct {
	jwt.StandardClaims
}

type AccessTokenClaims struct {
	jwt.StandardClaims

	Email string `json:"email"`
}

type Service struct {
	Ctx    context.Context
	JWTKey string

	UserRepository repository.User
	AuthRepository repository.Auth
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func (s *Service) Login(ctx context.Context, data *model.Login) (*TokenPair, error) {

	usr, err := s.UserRepository.CheckPassword(ctx, data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	token, err := s.generateAccessToken(usr.GetId(), usr.GetEmail())
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(usr.GetId())
	if err != nil {
		return nil, err
	}

	err = s.AuthRepository.SaveToken(ctx, refreshToken, usr.GetId())
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Register(ctx context.Context, data *model.Register) (*user_v1.User, error) {

	usr, err := s.UserRepository.Create(ctx, data.Email, data.Password)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error: Unable to create user. %s", err))
	}

	return usr, nil
}

func (s *Service) Logout(context.Context, string) (bool, error) {
	// удалять refresh токен из БД
	return false, nil
}

func (s *Service) CheckTokens(ctx context.Context, tokenPair *model.Check) (*model.Check, error) {
	var userId string

	if tokenPair.AccessToken != "" {
		// забираем пэйлоад у токена для получения информации о пользователе
		claims, err := getTokenPayload[AccessTokenClaims](tokenPair.AccessToken)
		if err != nil {
			return nil, err
		}
		userId = claims.Subject
	} else if tokenPair.RefreshToken != "" {
		claims, err := getTokenPayload[RefreshTokenClaims](tokenPair.RefreshToken)
		if err != nil {
			return nil, err
		}
		userId = claims.Subject
	} else {
		return nil, fmt.Errorf("no token provided")
	}

	// Забираем пользовательские данные из другого сервиса
	usr, err := s.UserRepository.Get(ctx, userId, "")
	if err != nil {
		return nil, fmt.Errorf("unable to find user. %s", err)
	}

	// Проверяем access токен и если тот не валиден, то refresh
	_, err = s.validateToken(tokenPair.AccessToken)
	if err != nil {
		_, err = s.validateToken(tokenPair.RefreshToken)
		if err != nil {
			return nil, err
		}
	}

	// Генерируем новые токены

	newAccessToken, err := s.generateAccessToken(usr.GetId(), usr.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("unable to create access token: %s", err)
	}
	newRefreshToken, err := s.generateRefreshToken(usr.GetId())
	if err != nil {
		return nil, fmt.Errorf("unable to create refresh token: %s", err)
	}

	return &model.Check{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil

}

func (s *Service) generateAccessToken(id string, email string) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(time.Hour * 3).Unix(),
		},
		Email: email,
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey))
}

func (s *Service) generateRefreshToken(id string) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   id,
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey))
}

func getTokenPayload[T interface{}](token string) (*T, error) {
	// Разделение токена на части
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	// Декодирование payload части
	payload, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		return nil, fmt.Errorf("unable to decode token")
	}

	// Парсинг payload части в структуру
	var claims T
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		return nil, fmt.Errorf("unable to parse payload")
	}
	return &claims, nil
}

func (s *Service) validateToken(
	token string,
) (*jwt.Token, error) {
	validToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи тот, что мы ожидаем
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// отдаем ключ подписи
		return []byte(s.JWTKey), nil
	})

	if err != nil || !validToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return validToken, nil
}
