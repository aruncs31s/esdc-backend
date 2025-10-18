package routes

import (
	"esdc-backend/internal/module/common/handler"

	"github.com/gin-gonic/gin"
)

func RegisterPublicProjectRoutes(r *gin.Engine, projectHandler handler.ProjectHandler) {
	publicProjectRoutes := r.Group("/api/projects")
	{
		publicProjectRoutes.GET("", projectHandler.GetPublicProjects)
		publicProjectRoutes.GET("/:id", projectHandler.GetProject)

	}
}
