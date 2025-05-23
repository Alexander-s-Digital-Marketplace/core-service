package authmiddlewares

import (
	"net/http"

	authserviceclient "github.com/Alexander-s-Digital-Marketplace/core-service/internal/services/auth_service/auth_service_client"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
		}
		tokenString = tokenString[len("Bearer "):]

		code, id, role := authserviceclient.ValidAccessToken(tokenString)
		if code == 200 {
			c.Set("id", id)
			c.Set("role", role)
			c.Next()
		} else {
			c.Abort()
		}
	}
}
