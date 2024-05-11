package api

import (
	handler "goAuth/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func UseRoutes(engine *gin.Engine) {
	userHandler := handler.NewUserHandler()

	engine.GET("/user/:userID", userHandler.Get)
}
