package converter

import (
	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/pkg/auth_v1"
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
