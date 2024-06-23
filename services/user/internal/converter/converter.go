package converter

import (
	"github.com/Insid1/go-auth-user/user/internal/model"
	"github.com/Insid1/go-auth-user/user/pkg/user_v1"
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
		PassHash:  user.PassHash,
	}
}

func ToModelFromUser(user *user_v1.User) *model.User {
	return &model.User{
		Email:     user.GetEmail(),
		Name:      user.GetName().Value,
		Age:       user.GetAge(),
		PassHash:  user.GetPassHash(),
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
	}
}
