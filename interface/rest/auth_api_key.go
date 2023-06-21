package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func apiKeyAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader, headerExists := c.Request.Header["Authorization"]
		if !headerExists || len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Key passed in the request"})
			return
		}

		if authHeader[0] != os.Getenv("API_KEY") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid Key passed in the request"})
			return
		}

		c.Next()
	}
}
