package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerOrderRoutes(r *gin.Engine, handler *handler.OrderHandler) {
	orders := r.Group("/api/orders")
	{
		orders.GET("", handler.GetOrders)
		orders.GET("/:id", handler.GetOrderByID)
		orders.POST("", handler.CreateOrder)
	}
}
