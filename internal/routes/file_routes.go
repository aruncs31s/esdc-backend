package routes

import (
	"esdc-backend/internal/module/common/handler"

	"github.com/gin-gonic/gin"
)

func registerFileRoutes(r *gin.Engine, fileHandler handler.FileHandler) {
	fileRoutes := r.Group("/api/files")
	{
		// Single image upload
		fileRoutes.POST("/upload/image", fileHandler.UploadImage)

		// Single file upload (any type)
		fileRoutes.POST("/upload", fileHandler.UploadFile)

		// Multiple files upload
		fileRoutes.POST("/upload/multiple", fileHandler.UploadMultipleFiles)
	}
}
