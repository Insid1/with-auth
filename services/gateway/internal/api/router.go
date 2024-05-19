package api

import (
	"github.com/Insid1/go-auth-user/gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(httpEngine *gin.Engine, handler handler.User) {
	httpEngine.POST("/user", handler.Create)
	httpEngine.GET("/user/:userID", handler.Get)
	httpEngine.PUT("/user/:userID", handler.Update)
	httpEngine.DELETE("/user/:userID", handler.Delete)
}
