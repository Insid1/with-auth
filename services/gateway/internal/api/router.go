package api

import (
	"github.com/Insid1/go-auth-user/gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(httpEngine *gin.Engine, handler handler.User) {
	httpEngine.POST("/auth_v1", handler.Create)
	httpEngine.GET("/auth_v1/:userID", handler.Get)
	httpEngine.PUT("/auth_v1/:userID", handler.Update)
	httpEngine.DELETE("/auth_v1/:userID", handler.Delete)
}
