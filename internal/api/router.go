package api

import (
	handler "github.com/Insid1/go-auth-user/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func UseRoutes(engine *gin.Engine) {
	userHandler := handler.NewUserHandler()

	engine.POST("/user", userHandler.Create)

	engine.GET("/user/:userID", userHandler.Get)
	engine.PUT("/user/:userID", userHandler.Update)
	engine.DELETE("/user/:userID", userHandler.Delete)
}
