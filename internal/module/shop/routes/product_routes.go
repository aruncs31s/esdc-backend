package routes

import (
	handler "esdc-backend/internal/module/shop/handler"

	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, handler handler.ProductHandler) {
	products := r.Group("/api/products")
	{
		products.GET("", handler.GetAll)
		products.GET("/:id", handler.GetByID)
	}
}
