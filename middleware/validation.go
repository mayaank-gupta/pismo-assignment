package middlewares

import (
	"net/http"
	"pismo-assignment/pkg"
	"regexp"

	"github.com/gin-gonic/gin"
)

func ValidateAccountCreation() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req pkg.CreateAccountRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid document_number field in request body"})
			c.Abort() // Stop further processing
			return
		}

		isValid := regexp.MustCompile(`^[0-9]{11}$`).MatchString(req.DocumentNumber)
		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "document_number must contain exactly 11 digits"})
			c.Abort()
			return
		}

		// Store the validated data in context for later use by the handler
		c.Set("validatedRequest", req)
		c.Next() // Proceed to the next middleware or handler
	}
}

func ValidateTransactionRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req pkg.CreateTransactionRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid document_number field in request body"})
			c.Abort() // Stop further processing
			return
		}

		isValid := regexp.MustCompile(`^[0-9]{11}$`).MatchString(req.)
		if !isValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": "document_number must contain exactly 11 digits"})
			c.Abort()
			return
		}

		// Store the validated data in context for later use by the handler
		c.Set("validatedRequest", req)
		c.Next() // Proceed to the next middleware or handler
	}
}
