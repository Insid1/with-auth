package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/user-service/internal/converter"
	"github.com/Insid1/go-auth-user/user-service/internal/service"
	"github.com/Insid1/go-auth-user/user-service/pkg/auth_v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	auth_v1.UnimplementedAuthV1Server

	Service service.Auth
}

const (
	emptyValue = 0
)

func (h *Handler) Login(ctx context.Context, req *auth_v1.LoginRequest) (*auth_v1.LoginResponse, error) {
	err := validateLogin(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	loginData := converter.ToModelFromLoginReq(req)
	token, err := h.Service.Login(loginData)
	if err != nil {
		// todo: можно дообработать ошибку
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.LoginResponse{Token: token}, nil
}

func (h *Handler) Register(ctx context.Context, req *auth_v1.RegisterRequest) (*auth_v1.RegisterResponse, error) {

	err := validateRegister(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	userID, err := h.Service.Register(converter.ToModelFromRegisterReq(req))
	if err != nil {
		// todo: можно дообработать ошибку
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.RegisterResponse{
		UserId: userID,
	}, nil
}

func (h *Handler) Logout(ctx context.Context, req *auth_v1.LogoutRequest) (*auth_v1.LogoutResponse, error) {
	isSuccess, err := h.Service.Logout(req.GetToken())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &auth_v1.LogoutResponse{
		Success: isSuccess,
	}, nil
}

func validateLogin(req *auth_v1.LoginRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	if req.GetAppId() == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}
	return nil
}

func validateRegister(req *auth_v1.RegisterRequest) error {
	if req.GetEmail() == "" {
		return status.Error(codes.InvalidArgument, "email is required")
	}
	if req.GetPassword() == "" {
		return status.Error(codes.InvalidArgument, "password is required")
	}
	return nil
}
