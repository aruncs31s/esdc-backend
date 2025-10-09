package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerCartRoutes(r *gin.Engine, handler *handler.CartHandler) {
	cart := r.Group("/api/cart")
	{
		cart.GET("", handler.GetCart)
		cart.POST("", handler.AddToCart)
		cart.PUT("/:id", handler.UpdateCart)
		cart.DELETE("/:id", handler.RemoveFromCart)
		cart.DELETE("", handler.ClearCart)
	}
}
