package handler

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/handler/auth"
	"github.com/Insid1/go-auth-user/auth-service/pkg/auth_v1"
)

type Auth interface {
	auth_v1.AuthV1Server
}

func NewAuthHandler(ctx context.Context) Auth {
	return &auth.Handler{Ctx: ctx}
}
