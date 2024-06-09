package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/pkg/auth_v1"
)

type Handler struct {
	auth_v1.UnimplementedAuthV1Server

	Ctx context.Context
}
