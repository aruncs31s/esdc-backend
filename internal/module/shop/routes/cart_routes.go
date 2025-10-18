package routes

import (
	handler "esdc-backend/internal/module/shop/handler"

	"github.com/gin-gonic/gin"
)

func RegisterCartRoutes(r *gin.Engine, handler handler.CartHandler) {
	cart := r.Group("/api/cart")
	{
		cart.GET("", handler.GetCart)
		cart.POST("", handler.AddToCart)
		cart.PUT("/:id", handler.UpdateCart)
		cart.DELETE("/:id", handler.RemoveFromCart)
		cart.DELETE("", handler.ClearCart)
	}
}
