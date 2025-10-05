package routes

import (
	middlewares "pismo-assignment/middleware"
	"pismo-assignment/services"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/accounts", middlewares.ValidateAccountCreation(), services.CreateAccount)
	router.GET("/accounts/:id", services.GetAccountByID)
}

func RegisterTransactionRoutes(router *gin.Engine) {
	router.POST("/transactions", middlewares.ValidateAccountCreation(), services.CreateAccount)
}
