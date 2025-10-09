package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerProductRoutes(r *gin.Engine, handler *handler.ProductHandler) {
	products := r.Group("/api/products")
	{
		products.GET("", handler.GetAll)
		products.GET("/:id", handler.GetByID)
	}
}
