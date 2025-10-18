package middleware

import (
	"esdc-backend/internal/module/auth/service"
	"fmt"
	"net/http"
	"strings"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		responseHelper := responsehelper.NewResponseHelper()
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			responseHelper.Unauthorized(c, "Authorization header required")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			responseHelper.Unauthorized(c, "Invalid authorization format. Use: Bearer <token>")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		// Parse the token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return service.SecretKey, nil
		})

		if err != nil {
			responseHelper.Unauthorized(c, "Invalid or expired token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			responseHelper.Unauthorized(c, "Token is not valid")
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
