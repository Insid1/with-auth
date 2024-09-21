package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"

	"github.com/Insid1/go-auth-user/auth-service/internal/converter"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	auth_v1.UnimplementedAuthV1Server

	AuthService service.Auth
}

func (h *Handler) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	tokenPair, err := h.AuthService.Login(ctx, converter.ToLoginModelFromReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.LoginResponse{
		AuthToken:    tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}

func (h *Handler) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	usr, err := h.AuthService.Register(ctx, converter.ToRegisterModelFromReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.RegisterResponse{
		User: usr,
	}, nil
}
