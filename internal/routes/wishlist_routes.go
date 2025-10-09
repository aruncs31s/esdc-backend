package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerWishlistRoutes(r *gin.Engine, handler *handler.WishlistHandler) {
	wishlist := r.Group("/api/wishlist")
	{
		wishlist.GET("", handler.GetWishlist)
		wishlist.POST("", handler.AddToWishlist)
		wishlist.DELETE("/:id", handler.RemoveFromWishlist)
	}
}
