package routes

import (
	"esdc-backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func registerChatbotRoutes(r *gin.Engine, chatbotHandler handler.ChatBotHandler) {
	chatbotRoutes := r.Group("/api/chatbot")
	{
		chatbotRoutes.POST("/ask", chatbotHandler.AskAI)

	}
}
