package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterChatRoutes(r *gin.Engine) {
	chat := r.Group("/ws")
	{
		chat.GET("/chat", handler.HandleWebSocket)
	}

	api := r.Group("/api/chat")
	{
		api.GET("/messages", handler.GetMessages)
	}
}
