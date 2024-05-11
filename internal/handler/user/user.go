package handler

import (
	"net/http"
	"strconv"

	"goAuth/internal/handler"
	service2 "goAuth/internal/service"
	user2 "goAuth/internal/service/user"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service2.User
}

func (u UserHandler) Create(ctx *gin.Context) {

	// TODO implement me
	panic("implement me")
}

func (u UserHandler) Get(ctx *gin.Context) {
	rawUserID := ctx.Param("userID")

	userID, err := strconv.ParseUint(rawUserID, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	usr, err := u.service.Get(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"user": usr})

}

func (u UserHandler) Update(ctx *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func (u UserHandler) Delete(ctx *gin.Context) {
	// TODO implement me
	panic("implement me")
}

func NewUserHandler() handler.User {
	return &UserHandler{service: user2.NewUserService()}
}
