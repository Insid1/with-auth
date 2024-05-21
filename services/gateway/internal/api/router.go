package api

import (
	"github.com/Insid1/go-auth-user/gateway/internal/handler"
	"github.com/gin-gonic/gin"
)

func UseRoutes(httpEngine *gin.Engine, handler handler.User) {
	httpEngine.POST("/user_v1", handler.Create)
	httpEngine.GET("/user_v1/:userID", handler.Get)
	httpEngine.PUT("/user_v1/:userID", handler.Update)
	httpEngine.DELETE("/user_v1/:userID", handler.Delete)
}
