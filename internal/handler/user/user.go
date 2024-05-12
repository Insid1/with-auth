package handler

import (
	"net/http"

	"goAuth/internal/entity"
	"goAuth/internal/handler"
	service2 "goAuth/internal/service"
	user2 "goAuth/internal/service/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	service service2.User
}

func (u UserHandler) Create(ctx *gin.Context) {
	var user entity.User
	uid := uuid.New()

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Id = uid.String()

	userID, err := u.service.Create(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, userID)
}

func (u UserHandler) Get(ctx *gin.Context) {
	rawUserID := ctx.Param("userID")

	usr, err := u.service.Get(rawUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"user": usr})
}

func (u UserHandler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
}

func (u UserHandler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
}

func NewUserHandler() handler.User {
	return &UserHandler{service: user2.NewUserService()}
}
