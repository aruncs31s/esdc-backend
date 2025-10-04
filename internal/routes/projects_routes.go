package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerProjectsRoutes(r *gin.Engine, projectHandler handler.ProjectHandler) {
	projectRoutes := r.Group("/api/projects")
	{
		projectRoutes.GET("/", projectHandler.GetAllProjects)
		projectRoutes.POST("/", projectHandler.CreateProject)

		// projectRoutes.GET("/", projectHandler.List)
		// projectRoutes.GET("/:id", projectHandler.Get)
		// projectRoutes.PUT("/:id", projectHandler.Update)
		// projectRoutes.DELETE("/:id", projectHandler.Delete)
	}
}
