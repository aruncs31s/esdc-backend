package handler

import (
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	GetAllPosts(c *gin.Context)
}

type postsHandler struct {
	postsService   service.PostsService
	responseHelper responses.ResponseHelper
}

func NewPostsHandler(postsService service.PostsService) PostHandler {
	responseHelper := responses.NewResponseHelper()
	return &postsHandler{postsService: postsService, responseHelper: responseHelper}
}

func (h *postsHandler) GetAllPosts(c *gin.Context) {
	// Check if the user is an admin
	role := c.GetString("role")

	if role != "admin" {
		h.responseHelper.Unauthorized(c, "You are not authorized to access this resource")
		return
	}

	posts, err := h.postsService.GetAllPosts()
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to retrieve posts", err)
		return
	}
	h.responseHelper.Success(c, posts)
}
