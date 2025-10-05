package services

import (
	"net/http"
	"pismo-assignment/pkg"

	"github.com/gin-gonic/gin"
)

func createTransaction(c *gin.Context) {
	req, _ := c.Get("validatedRequest")

	transactionReq := req.(pkg.CreateTransactionRequest)

	c.JSON(http.StatusOK, gin.H{"status": "callback processed"})
}
