package converter

import (
	"github.com/Insid1/go-auth-user/auth-service/internal/model"
	"github.com/Insid1/go-auth-user/auth-service/pkg/auth_v1"
)

func toLoginModelFromReq(req *auth_v1.LoginRequest) *model.Login {
	email := req.GetEmail()
	return &model.Login{
		Email:    &email,
		Password: req.GetPassword(),
	}
}

func toRegisterModelFromReq(req *auth_v1.RegisterRequest) *model.Register {
	email := req.GetEmail()

	return &model.Register{
		Email:    &email,
		Password: req.GetPassword(),
	}
}
