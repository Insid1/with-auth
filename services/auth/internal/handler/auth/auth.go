package auth

import (
	"context"

	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"

	"github.com/Insid1/go-auth-user/auth-service/internal/converter"
	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/internal/service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	auth_v1.UnimplementedAuthV1Server

	AuthService service.Auth
}

func (h *Handler) Login(ctx context.Context, req *auth_v1.LoginReq) (*auth_v1.LoginRes, error) {
	tokenPair, err := h.AuthService.Login(ctx, converter.ToLoginModelFromReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.LoginRes{
		TokenPair: &auth_v1.TokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}, nil
}

func (h *Handler) Register(ctx context.Context, req *auth_v1.RegisterReq) (*auth_v1.RegisterRes, error) {
	usr, err := h.AuthService.Register(ctx, converter.ToRegisterModelFromReq(req))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	tokenPair, err := h.AuthService.Login(ctx, &model.Login{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.RegisterRes{
		User:      usr,
		TokenPair: &auth_v1.TokenPair{AccessToken: tokenPair.AccessToken, RefreshToken: tokenPair.RefreshToken},
	}, nil
}

func (h *Handler) LogoutAll(ctx context.Context, req *auth_v1.LogoutAllReq) (*auth_v1.LogoutAllRes, error) {
	claims, err := h.AuthService.CheckAccessToken(ctx, req.GetAccessToken())

	if err != nil {
		return nil, status.Error(codes.PermissionDenied, err.Error())
	}

	err = h.AuthService.LogoutAll(ctx, claims.Subject)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.LogoutAllRes{
		Success: true,
	}, nil
}

func (h *Handler) Check(ctx context.Context, req *auth_v1.CheckReq) (*auth_v1.CheckRes, error) {

	_, err := h.AuthService.CheckAccessToken(ctx, req.GetAccessToken())

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	return &auth_v1.CheckRes{
		Success: true,
	}, nil
}

func (h *Handler) Refresh(ctx context.Context, req *auth_v1.RefreshReq) (*auth_v1.RefreshRes, error) {
	claims, err := h.AuthService.CheckRefreshToken(ctx, req.GetRefreshToken())

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	tokenPair, err := h.AuthService.GenerateTokenPair(ctx, claims.Subject, "")

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth_v1.RefreshRes{
		TokenPair: &auth_v1.TokenPair{
			AccessToken:  tokenPair.AccessToken,
			RefreshToken: tokenPair.RefreshToken,
		},
	}, nil
}
