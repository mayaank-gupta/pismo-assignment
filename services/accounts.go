package services

import (
	"net/http"
	"pismo-assignment/pkg"
	"pismo-assignment/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

var log = utils.GetLogger("service:create_account")

func CreateAccount(c *gin.Context) {
	req, _ := c.Get("validatedRequest")

	accountReq := req.(pkg.CreateAccountRequest)

	c.JSON(http.StatusOK, gin.H{"status": "callback processed"})
}

func GetAccountByID(c *gin.Context) {
	var payload map[string]interface{}

	id, _ := strconv.Atoi(c.Param("id"))
	var account models.Account

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.WithError(err).Warn("Invalid callback payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "callback processed"})
}
