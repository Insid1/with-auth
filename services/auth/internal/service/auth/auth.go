package auth

import (
	"context"
	"encoding/json"
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

// Входа пользователя в систему
func (s *Service) Login(ctx context.Context, data *model.Login) (*TokenPair, error) {

	usr, err := s.UserRepository.CheckPassword(ctx, data.Email, data.Password)

	if err != nil {
		return nil, err
	}

	return s.GenerateTokenPair(ctx, usr.GetId(), usr.GetEmail())
}

// Метод регистрации пользователя
func (s *Service) Register(ctx context.Context, data *model.Register) (*user_v1.User, error) {

	usr, err := s.UserRepository.Create(ctx, data.Email, data.Password)
	if err != nil {
		return nil, fmt.Errorf("error: Unable to create user. %s", err)
	}

	return usr, nil
}

// Метод Выхода из всех устройств пользователя
func (s *Service) LogoutAll(ctx context.Context, userId string) error {

	_, err := s.AuthRepository.GenerateJWTUserKey(ctx, userId)

	return err
}

// Метод генерации Access и Refresh
func (s *Service) GenerateTokenPair(ctx context.Context, userId string, email string) (*TokenPair, error) {
	usr, err := s.UserRepository.Get(ctx, userId, email)

	if err != nil {
		return nil, err
	}

	// todo добавить генерацию JWTUserKey если такового не существует для такого пользователя
	jwtExtraKey, err := s.mustGetJWTExtraKey(ctx, usr.Id)

	if err != nil {
		return nil, err
	}

	token, err := s.generateAccessToken(usr)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.generateRefreshToken(usr.GetId(), jwtExtraKey)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  token,
		RefreshToken: refreshToken,
	}, nil
}

// Метод проверки Access токена
func (s *Service) CheckAccessToken(ctx context.Context, accessToken string) (*AccessTokenClaims, error) {
	_, err := s.validateToken(accessToken, "")

	if err != nil {
		return nil, err
	}

	claims, err := getTokenPayload[AccessTokenClaims](accessToken)

	if err != nil {
		return nil, err
	}

	return claims, nil
}

// Метод проверки Refresh токена.
func (s *Service) CheckRefreshToken(ctx context.Context, refreshToken string) (*RefreshTokenClaims, error) {
	// забираем пэйлоад у токена для получения информации о пользователе
	claims, err := getTokenPayload[RefreshTokenClaims](refreshToken)
	if err != nil {
		return nil, err
	}

	JWTUserKey, err := s.mustGetJWTExtraKey(ctx, claims.Subject)

	if err != nil {
		return nil, err
	}

	_, err = s.validateToken(refreshToken, JWTUserKey)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

// Метод генерации Access токена
func (s *Service) generateAccessToken(usr *user_v1.User) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: usr.GetId(),
			// 1 hour
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
		},
		Email: usr.GetEmail(),
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey))
}

// Метод генерации Refresh токена
func (s *Service) generateRefreshToken(id string, extraKeyData string) (string, error) {
	// Генерируем полезные данные, которые будут храниться в токене
	payload := RefreshTokenClaims{
		StandardClaims: jwt.StandardClaims{
			Subject: id,
			// 21 day
			ExpiresAt: time.Now().Add(time.Hour * 504).Unix(),
		},
	}

	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return token.SignedString([]byte(s.JWTKey + extraKeyData))
}

// Метод получения пэйлоада из токена
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

// Метод проверки токена
func (s *Service) validateToken(
	token string,
	extraJWTKey string,
) (*jwt.Token, error) {
	validToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Проверяем, что алгоритм подписи тот, что мы ожидаем
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// отдаем ключ подписи
		return []byte(s.JWTKey + extraJWTKey), nil
	})

	if err != nil || !validToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return validToken, nil
}

// Метод получения дополнительного ключа для JWT токена. Если ключа не существует он будет сгенерирован
func (s *Service) mustGetJWTExtraKey(ctx context.Context, userID string) (string, error) {
	jwtUserKey, err := s.AuthRepository.GetJWTUserKey(ctx, userID)

	if err == nil {
		return jwtUserKey, nil
	}
	// todo добавить проверку конкретной ошибки, что за текущим пользователем действительно нет JWTKey

	return s.AuthRepository.GenerateJWTUserKey(ctx, userID)
}
