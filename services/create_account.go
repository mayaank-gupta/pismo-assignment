package services

import (
	"net/http"
	"pismo-assignment/utils"

	"github.com/gin-gonic/gin"
)

var log = utils.GetLogger("service:create_account")

func CreateAccount(c *gin.Context) {
	var payload map[string]interface{}

	if err := c.ShouldBindJSON(&payload); err != nil {
		log.WithError(err).Warn("Invalid callback payload")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid payload"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "callback processed"})
}
