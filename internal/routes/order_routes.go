package routes

import (
	"esdc-backend/internal/module/common/handler"

	"github.com/gin-gonic/gin"
)

func registerOrderRoutes(r *gin.Engine, orderHandler *handler.OrderHandler) {
	orders := r.Group("/api/orders")
	{
		orders.GET("", orderHandler.GetOrders)
		orders.GET("/:id", orderHandler.GetOrderByID)
		orders.POST("", orderHandler.CreateOrder)
	}
}
