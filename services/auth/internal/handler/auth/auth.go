package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/auth-service/internal/converter"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"
	"github.com/Insid1/go-auth-user/auth-service/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	auth_v1.UnimplementedAuthV1Server

	Ctx         context.Context
	AuthService service.Auth
}

func (h *Handler) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	token, err := h.AuthService.Login(converter.ToLoginModelFromReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.LoginResponse{
		Token: token,
	}, nil
}

func (h *Handler) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {
	userID, err := h.AuthService.Register(converter.ToRegisterModelFromReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.RegisterResponse{
		UserId: userID,
	}, nil
}
