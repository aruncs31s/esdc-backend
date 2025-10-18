package handler

import (
	"esdc-backend/internal/module/shop/dto"
	"esdc-backend/internal/module/shop/service"
	"esdc-backend/utils"
	"net/http"
	"strconv"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type WishlistHandler interface {
	GetWishlist(c *gin.Context)
	AddToWishlist(c *gin.Context)
	RemoveFromWishlist(c *gin.Context)
}
type wishlistHandler struct {
	service        service.WishlistService
	responseHelper responsehelper.ResponseHelper
}

func NewWishlistHandler(service service.WishlistService) WishlistHandler {
	return &wishlistHandler{
		service: service,
	}
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
func (h *wishlistHandler) GetWishlist(c *gin.Context) {
	userID, failed := verifyUserID(c, h)
	if failed {
		return
	}

	items, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		h.responseHelper.InternalError(c, "Something Wrong", err)
		return
	}

	h.responseHelper.Success(c, items)

}

func verifyUserID(c *gin.Context, h *wishlistHandler) (any, bool) {
	userID, exists := c.Get("user_id")
	if !exists {
		h.responseHelper.Unauthorized(c, "user not authenticated")
		return nil, true
	}
	return userID, false
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
func (h *wishlistHandler) AddToWishlist(c *gin.Context) {
	userID, failed := verifyUserID(c, h)
	if failed {
		return
	}

	product, failed := utils.GetJSONData[dto.AddToWishlistRequest](c, h.responseHelper, "Invalid request data", "Err Payload")
	if failed {
		return
	}

	wishlist, err := h.service.Add(userID.(uint), product.ProductID)
	if err != nil {
		if err.Error() == "product already in wishlist" {
			h.responseHelper.BadRequest(c, "Product already in wishlist", err.Error())
			return
		}
		h.responseHelper.InternalError(c, "Something Wrong", err)
		return
	}
	data := map[string]interface{}{
		"id":         wishlist.ID,
		"product_id": wishlist.ProductID,
	}

	h.responseHelper.Success(c, data)

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
func (h *wishlistHandler) RemoveFromWishlist(c *gin.Context) {
	// OPTIMIZE: this.
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
