package handler

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"
	"strings"

	"github.com/gin-gonic/gin"
)

type ChatBotHandler interface {
	AskAI(c *gin.Context)
}
type chatBotHandler struct {
	responseHelper responses.ResponseHelper
	chatBotService service.ChatBotService
}

func NewChatBotHandler(chatBotService service.ChatBotService) ChatBotHandler {
	responseHelper := responses.NewResponseHelper()
	return &chatBotHandler{
		responseHelper: responseHelper,
		chatBotService: chatBotService,
	}
}

func (h *chatBotHandler) AskAI(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		h.responseHelper.Unauthorized(c, "user not logged in")
		return
	}
	var messageQuery dto.ChatBotRequest
	if err := c.ShouldBindJSON(&messageQuery); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	response, err := h.chatBotService.Ask(&user, strings.TrimSpace(messageQuery.QueryMessage))
	if err != nil {
		h.responseHelper.InternalError(c, "something bad happend", err)
		return
	}
	h.responseHelper.Success(c, response)
}
