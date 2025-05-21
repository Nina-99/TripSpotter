package middleware

import (
	"net/http"
	"strings"

	"github.com/Nina-99/TripSpotter/backend/utils"
	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		token := strings.Split(authHeader, "Bearer ")
		if len(token) < 2 {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		claims, err := utils.ValidateToken(token[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
