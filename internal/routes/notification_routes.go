package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerNotificationRoutes(r *gin.Engine, notificationHandler handler.NotificationHandler) {
	notificationRoutes := r.Group("/api/notifications")
	{
		notificationRoutes.POST("", notificationHandler.SendNotification)
		notificationRoutes.GET("/:id", notificationHandler.GetSingleNotification)
		notificationRoutes.GET("", notificationHandler.GetUserNotifications)
		notificationRoutes.PUT("/:id/read", notificationHandler.MarkAsRead)
	}
}
