package inmiddlewares

import (
	"net/http"
	"strings"

	"github.com/eac0de/atlas/sentinel/internal/services"

	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(sessionService *services.SessionService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Authorization header is required"})
			c.Abort()
			return
		}

		clearToken, ok := strings.CutPrefix(authorizationHeader, "Bearer ")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Invalid Authorization header"})
			c.Abort()
			return
		}

		claims, err := sessionService.ParseToken(clearToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": "Invalid token"})
			c.Abort()
			return
		}
		c.Set(gin.AuthUserKey, claims.UserID)
		c.Next()
	}
}
