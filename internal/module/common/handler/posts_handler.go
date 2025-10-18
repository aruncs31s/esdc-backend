package handler

import (
	"esdc-backend/internal/module/common/service"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type PostHandler interface {
	GetAllPosts(c *gin.Context)
}

type postsHandler struct {
	postsService   service.PostsService
	responseHelper responsehelper.ResponseHelper
}

func NewPostsHandler(postsService service.PostsService) PostHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &postsHandler{postsService: postsService, responseHelper: responseHelper}
}

// GetAllPosts godoc
// @Summary Get all posts (Admin only)
// @Description Retrieve all posts - requires admin role
// @Tags posts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Posts retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized - admin role required"
// @Failure 500 {object} map[string]interface{} "Failed to retrieve posts"
// @Router /posts [get]
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
