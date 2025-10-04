package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerAdminRoutes(r *gin.Engine, adminHandler handler.AdminHandler) {
	adminRoutes := r.Group("/api/admin")
	{
		adminRoutes.GET("/users", adminHandler.GetAllUsers)
		adminRoutes.GET("/stats", adminHandler.GetUsersStats)
		adminRoutes.DELETE("/users/:id", adminHandler.DeleteUser)
		adminRoutes.POST("/users", adminHandler.CreateUser)
	}

}
