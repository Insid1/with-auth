package main

import (
	"github.com/Insid1/go-auth-user/internal/api"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	api.UseRoutes(engine)

	engine.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8080")
}
