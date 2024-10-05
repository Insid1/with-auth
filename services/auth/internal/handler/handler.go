package handler

import (
	"github.com/Insid1/with-auth/auth-service/internal/handler/auth"
	"github.com/Insid1/with-auth/auth-service/internal/service"
	"github.com/Insid1/with-auth/pkg/grpc/auth_v1"
)

type Auth interface {
	auth_v1.AuthV1Server
}

func NewAuthHandler(authService service.Auth) Auth {
	return &auth.Handler{
		UnimplementedAuthV1Server: auth_v1.UnimplementedAuthV1Server{},
		AuthService:               authService,
	}
}
