package middleware

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Example: Set userID based on session or JWT token
		// Replace this with your actual authentication logic
		userID := 3 // Replace with actual logic to get userID
		c.Set("userID", userID)
		c.Next()
	}
}
