package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerProtectedProjectsRoutes(r *gin.Engine, projectHandler handler.ProjectHandler) {
	projectRoutes := r.Group("/api/projects")
	{
		projectRoutes.POST("", projectHandler.CreateProject)

	}
}
func registerProjectsRoutes(r *gin.Engine, projectHandler handler.ProjectHandler) {
	projectRoutes := r.Group("/api/projects")
	{
		projectRoutes.GET("", projectHandler.GetAllProjects)
		projectRoutes.GET("/:id", projectHandler.GetProject)

	}
}

