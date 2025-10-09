package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerUserRoutes(r *gin.Engine, userHandler handler.UserHandler) {
	userRoutes := r.Group("/api/user")
	{
		userRoutes.POST("/login", userHandler.Login)
		userRoutes.POST("/register", userHandler.Register)
	}
	// }
}
