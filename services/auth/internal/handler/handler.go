package handler

import (
	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"

	"github.com/Insid1/go-auth-user/auth-service/internal/handler/auth"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"
)

type Auth interface {
	auth_v1.AuthV1Server
}

func NewAuthHandler(authService service.Auth) Auth {
	return &auth.Handler{AuthService: authService}
}
