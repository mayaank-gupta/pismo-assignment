package routes

import (
	"pismo-assignment/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/accounts", services.CreateAccount)
}
