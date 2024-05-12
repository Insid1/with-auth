package api

import (
	handler "goAuth/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func UseRoutes(engine *gin.Engine) {
	userHandler := handler.NewUserHandler()

	engine.POST("/user", userHandler.Create)

	engine.GET("/user/:userID", userHandler.Get)
	engine.PUT("/user/:userID", userHandler.Update)
	engine.DELETE("/user/:userID", userHandler.Delete)
}
