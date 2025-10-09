package middleware

import (
	"esdc-backend/internal/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("❌ JWT Middleware: No Authorization header provided")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			fmt.Println("❌ JWT Middleware: Invalid Authorization header format (missing 'Bearer ' prefix)")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format. Use: Bearer <token>"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		fmt.Printf("🔑 JWT Middleware: Validating token: %s...\n", tokenString[:min(20, len(tokenString))])

		// Parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return service.SecretKey, nil
		})

		if err != nil {
			fmt.Printf("❌ JWT Middleware: Token parsing error: %v\n", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			fmt.Println("❌ JWT Middleware: Token is not valid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract claims if needed
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Store claims in context for handlers to use
			c.Set("jwt_claims", claims)

			// Set username in context for handlers to use
			username := claims["username"]
			role := claims["role"]
			c.Set("user", username)
			c.Set("username", username)
			c.Set("role", role)
			fmt.Printf("✅ JWT Middleware: Token valid - User: %v, Role: %v\n", username, role)
		}

		c.Next()
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
