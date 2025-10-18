package routes

import (
	"esdc-backend/internal/module/shop/handler"

	"github.com/gin-gonic/gin"
)

func RegisterWishlistRoutes(r *gin.Engine, h handler.WishlistHandler) {
	wishlist := r.Group("/api/wishlist")
	{
		wishlist.GET("", h.GetWishlist)
		wishlist.POST("", h.AddToWishlist)
		wishlist.DELETE("/:id", h.RemoveFromWishlist)
	}
}
