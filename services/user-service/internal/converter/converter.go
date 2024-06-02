package converter

import (
	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/pkg/auth_v1"
	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromModel(user *model.User) *user_v1.User {
	return &user_v1.User{
		Id:        user.ID,
		Email:     user.Email,
		Name:      wrapperspb.String(user.Name),
		Age:       user.Age,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
}

func ToModelFromUser(user *user_v1.User) *model.User {
	return &model.User{
		Email:     user.GetEmail(),
		Name:      user.GetName().Value,
		Age:       user.GetAge(),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
	}
}

func ToModelFromLoginReq(login *auth_v1.LoginRequest) *model.Login {
	return &model.Login{
		AppID:    login.GetAppId(),
		Email:    login.GetEmail(),
		Password: login.GetPassword(),
	}
}

func ToModelFromRegisterReq(register *auth_v1.RegisterRequest) *model.Register {
	return &model.Register{
		Email:    register.GetEmail(),
		Password: register.GetPassword(),
	}
}
