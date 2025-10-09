package handler

import (
	"esdc-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	service *service.WishlistService
}

func NewWishlistHandler(service *service.WishlistService) *WishlistHandler {
	return &WishlistHandler{service: service}
}

// GetWishlist godoc
// @Summary Get user's wishlist
// @Description Get all items in the authenticated user's wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Wishlist retrieved successfully"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /wishlist [get]
func (h *WishlistHandler) GetWishlist(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
			"details": "User not authenticated",
		})
		return
	}

	items, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Internal server error",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    items,
	})
}

// AddToWishlist godoc
// @Summary Add product to wishlist
// @Description Add a product to the authenticated user's wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object{product_id=uint} true "Add to wishlist request"
// @Success 200 {object} map[string]interface{} "Product added to wishlist"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 409 {object} map[string]interface{} "Product already in wishlist"
// @Router /wishlist [post]
func (h *WishlistHandler) AddToWishlist(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
			"details": "User not authenticated",
		})
		return
	}

	var req struct {
		ProductID uint `json:"product_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Product ID is required",
		})
		return
	}

	wishlist, err := h.service.Add(userID.(uint), req.ProductID)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "product already in wishlist" {
			statusCode = http.StatusConflict
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   "Error adding to wishlist",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product added to wishlist",
		"data": gin.H{
			"id":         wishlist.ID,
			"product_id": wishlist.ProductID,
		},
	})
}

// RemoveFromWishlist godoc
// @Summary Remove item from wishlist
// @Description Remove a specific item from the wishlist
// @Tags wishlist
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Wishlist item ID"
// @Success 200 {object} map[string]interface{} "Item removed from wishlist"
// @Failure 400 {object} map[string]interface{} "Invalid wishlist item ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /wishlist/{id} [delete]
func (h *WishlistHandler) RemoveFromWishlist(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Invalid wishlist item ID",
		})
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Internal server error",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Item removed from wishlist",
	})
}
