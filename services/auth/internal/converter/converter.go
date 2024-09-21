package converter

import (
	"github.com/Insid1/go-auth-user/pkg/grpc/auth_v1"

	"github.com/Insid1/go-auth-user/auth-service/internal/model"
)

func ToLoginModelFromReq(req *auth_v1.LoginRequest) *model.Login {
	return &model.Login{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToRegisterModelFromReq(req *auth_v1.RegisterRequest) *model.Register {

	return &model.Register{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToCheckModelFromReq(req *auth_v1.CheckRequest) *model.Check {
	return &model.Check{
		AccessToken:  req.GetAuthToken(),
		RefreshToken: req.GetRefreshToken(),
	}
}
