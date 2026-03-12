package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkjwt "github.com/clerk/clerk-sdk-go/v2/jwt"
	"github.com/gin-gonic/gin"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error:": " Access denied token not found"})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied token not found"})
			c.Abort()
			return
		}
		tokenString := parts[1]
		clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
		claims, err := clerkjwt.Verify(c.Request.Context(), &clerkjwt.VerifyParams{
			Token: tokenString,
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Access not valid or expired"})
			c.Abort()
			return
		}

		c.Set("user_id", claims.Subject)
		c.Next()
	}
}
