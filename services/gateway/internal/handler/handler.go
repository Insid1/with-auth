package handler

import (
	"github.com/Insid1/go-auth-user/gateway/internal/handler/user"
	"github.com/Insid1/go-auth-user/gateway/internal/service"
	"github.com/gin-gonic/gin"
)

type User interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func NewUserHandler() User {
	return &user.Handler{Service: service.NewUserService()}
}
