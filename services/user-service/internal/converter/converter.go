package converter

import (
	"github.com/Insid1/go-auth-user/user-service/internal/model"
	"github.com/Insid1/go-auth-user/user-service/pkg/user_v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromModel(user *model.User) *user_v1.User {
	return &user_v1.User{
		Id:    user.ID,
		Email: user.Email,
		Name:  wrapperspb.String(user.Name),
		Age:   user.Age,
	}
}
