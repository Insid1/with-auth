package converter

import (
	"github.com/Insid1/with-auth/pkg/grpc/auth_v1"

	"github.com/Insid1/with-auth/auth-service/internal/model"
)

func ToLoginModelFromReq(req *auth_v1.LoginReq) *model.Login {
	return &model.Login{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}

func ToRegisterModelFromReq(req *auth_v1.RegisterReq) *model.Register {

	return &model.Register{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}
