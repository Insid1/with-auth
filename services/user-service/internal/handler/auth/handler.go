package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/pkg/auth_v1"
)

type Handler struct {
	auth_v1.UnimplementedAuthV1Server
}

func (h *Handler) Login(context.Context, *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	return &auth_v1.LoginResponse{Token: "no token yet"}, nil
}

func (h *Handler) Register(context.Context, *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	return nil, nil
}

func (h *Handler) Logout(context.Context, *auth_v1.LogoutRequest) (*auth_v1.LogoutResponse, error) {
	return nil, nil
}
