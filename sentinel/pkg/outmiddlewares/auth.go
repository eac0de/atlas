package outmiddlewares

import (
	"net/http"
	"strings"

	pb "github.com/eac0de/atlas/sentinel/proto"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func NewAuthMiddleware(conn *grpc.ClientConn) gin.HandlerFunc {
	client := pb.NewAuthClient(conn)
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

		resp, err := client.AuthUser(c.Request.Context(), &pb.AuthUserRequest{Token: clearToken})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"detail": err.Error()})
			c.Abort()
			return
		}
		userID, err := uuid.Parse(resp.UserId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"detail": err.Error()})
			c.Abort()
			return
		}
		c.Set(gin.AuthUserKey, userID)
		c.Next()
	}
}

func NewAuthMiddlewareForTest(userID uuid.UUID) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(gin.AuthUserKey, userID)
		c.Next()
	}
}
