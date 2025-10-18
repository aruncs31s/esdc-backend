package shop

import (
	"esdc-backend/internal/module/shop/handler"
	"esdc-backend/internal/module/shop/routes"

	"github.com/gin-gonic/gin"
)

func RegisterShopRoutes(r *gin.Engine, cartHandler handler.CartHandler, productHandler handler.ProductHandler, wishlistHandler handler.WishlistHandler) {
	routes.RegisterCartRoutes(r, cartHandler)
	routes.RegisterProductRoutes(r, productHandler)
	routes.RegisterWishlistRoutes(r, wishlistHandler)
}
