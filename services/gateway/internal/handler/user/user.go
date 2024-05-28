package user

import (
	"net/http"

	"github.com/Insid1/go-auth-user/gateway/internal/entity"
	"github.com/Insid1/go-auth-user/gateway/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	Service service.User
}

func (u Handler) Create(ctx *gin.Context) {
	var usr entity.User
	uid := uuid.New()

	if err := ctx.BindJSON(&usr); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usr.Id = uid.String()

	userID, err := u.Service.Create(&usr)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, userID)
}

func (u Handler) Get(ctx *gin.Context) {
	rawUserID := ctx.Param("userID")

	usr, err := u.Service.Get(rawUserID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"auth_v1": usr})
}

func (u Handler) Update(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
}

func (u Handler) Delete(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": "not implemented"})
}
