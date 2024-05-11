package handler

import (
	"github.com/gin-gonic/gin"
)

type User interface {
	Create(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}
