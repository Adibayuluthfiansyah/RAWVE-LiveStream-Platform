package middleware

import (
	"github.com/gin-gonic/gin"
	// Uncomment these imports when enabling production auth:
	// "net/http"
	// "os"
	// "strings"
	// "github.com/clerk/clerk-sdk-go/v2"
	// clerkjwt "github.com/clerk/clerk-sdk-go/v2/jwt"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// DEVELOPMENT MODE: Hardcoded user_id
		// TODO: Uncomment production auth before deploying
		c.Set("user_id", "Adibayu")
		c.Next()

		// PRODUCTION AUTH (commented for development)
		// authHeader := c.GetHeader("Authorization")
		// if authHeader == "" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error:": " Access denied token not found"})
		// 	c.Abort()
		// 	return
		// }
		// parts := strings.Split(authHeader, " ")
		// if len(parts) != 2 || parts[0] != "Bearer" {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Access denied token not found"})
		// 	c.Abort()
		// 	return
		// }
		// tokenString := parts[1]
		// clerk.SetKey(os.Getenv("CLERK_SECRET_KEY"))
		// claims, err := clerkjwt.Verify(c.Request.Context(), &clerkjwt.VerifyParams{
		// 	Token: tokenString,
		// })
		// if err != nil {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Access not valid or expired"})
		// 	c.Abort()
		// 	return
		// }
		//
		// c.Set("user_id", claims.Subject)
		// c.Next()
	}
}
