package admin

import (
	"esdc-backend/internal/module/admin/handler"

	"github.com/gin-gonic/gin"
)

func RegisterAdminRoutes(r *gin.Engine, adminHandler handler.AdminHandler) {
	adminRoutes := r.Group("/api/admin")
	{
		adminRoutes.GET("/users", adminHandler.GetAllUsers)
		adminRoutes.GET("/stats", adminHandler.GetUsersStats)
		adminRoutes.DELETE("/users/:id", adminHandler.DeleteUser)
		adminRoutes.POST("/users", adminHandler.CreateUser)
	}
	{
		adminRoutes.GET("/projects", adminHandler.GetAllProjects)
		adminRoutes.GET("/projects/:id", adminHandler.GetProjectByID)
		adminRoutes.POST("/projects", adminHandler.CreateProject)
		adminRoutes.PUT("/projects/:id", adminHandler.UpdateProject)
		adminRoutes.DELETE("/projects/:id", adminHandler.DeleteProject)
	}

}
