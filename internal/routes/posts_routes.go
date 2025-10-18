package routes

import (
	"esdc-backend/internal/module/common/handler"

	"github.com/gin-gonic/gin"
)

func registerPostRoutes(r *gin.Engine, postHandler handler.PostHandler) {
	postRoutes := r.Group("/api/posts")
	{
		postRoutes.GET("/", postHandler.GetAllPosts)
	}
}
