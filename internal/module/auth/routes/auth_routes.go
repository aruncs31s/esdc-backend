package routes

import (
	"esdc-backend/internal/module/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, authHandler handler.AuthHandler) {
	userRoutes := r.Group("/api/user")
	{
		userRoutes.POST("/login", authHandler.Login)
		userRoutes.POST("/register", authHandler.Register)
	}
}
