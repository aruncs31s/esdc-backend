package handler

import (
	"esdc-backend/internal/module/shop/dto"
	service "esdc-backend/internal/module/shop/service"
	"net/http"
	"strconv"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type CartHandler interface {
	GetCart(c *gin.Context)
	AddToCart(c *gin.Context)
	UpdateCart(c *gin.Context)
	ClearCart(c *gin.Context)
	RemoveFromCart(c *gin.Context)
}

type cartHandler struct {
	service        service.CartService
	responseHelper responsehelper.ResponseHelper
}

func NewCartHandler(service service.CartService) CartHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &cartHandler{
		service:        service,
		responseHelper: responseHelper,
	}
}

// GetCart godoc
// @Summary Get user's cart
// @Description Get all items in the authenticated user's cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Cart retrieved successfully"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /cart [get]
func (h *cartHandler) GetCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		h.responseHelper.BadRequest(c, "Unauthorized", "User not authenticated")
		return
	}

	items, total, err := h.service.GetByUserID(userID.(uint))
	if err != nil {
		h.responseHelper.InternalError(c, "Internal server error", err)
		return
	}
	response := map[string]interface{}{
		"items": items,
		"total": total,
	}
	h.responseHelper.Success(c, response)
}

// AddToCart godoc
// @Summary Add product to cart
// @Description Add a product to the authenticated user's cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object{product_id=uint,quantity=int} true "Add to cart request"
// @Success 200 {object} map[string]interface{} "Product added to cart"
// @Failure 400 {object} map[string]interface{} "Invalid request data or insufficient stock"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /cart [post]
func (h *cartHandler) AddToCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		h.responseHelper.BadRequest(c, "Unauthorized", "User not authenticated")
		return
	}

	var req dto.AddToCartRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request data", "Product ID is required")
		return
	}

	if req.Quantity <= 0 {
		req.Quantity = 1
	}

	cart, err := h.service.Add(userID.(uint), req.ProductID, req.Quantity)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "product not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "insufficient stock" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   "Error adding to cart",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product added to cart",
		"data": gin.H{
			"id":         cart.ID,
			"product_id": cart.ProductID,
			"quantity":   cart.Quantity,
		},
	})
}

// UpdateCart godoc
// @Summary Update cart item quantity
// @Description Update the quantity of a specific cart item
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Cart item ID"
// @Param request body object{quantity=int} true "Update cart request"
// @Success 200 {object} map[string]interface{} "Cart updated"
// @Failure 400 {object} map[string]interface{} "Invalid request data or insufficient stock"
// @Failure 500 {object} map[string]interface{} "Error updating cart"
// @Router /cart/{id} [put]
func (h *cartHandler) UpdateCart(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Invalid cart item ID",
		})
		return
	}

	var req struct {
		Quantity int `json:"quantity" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Quantity is required",
		})
		return
	}

	if req.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Quantity must be greater than 0",
		})
		return
	}

	err = h.service.Update(uint(id), req.Quantity)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "insufficient stock" {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   "Error updating cart",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Cart updated",
		"data": gin.H{
			"id":       id,
			"quantity": req.Quantity,
		},
	})
}

// RemoveFromCart godoc
// @Summary Remove item from cart
// @Description Remove a specific item from the cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Cart item ID"
// @Success 200 {object} map[string]interface{} "Item removed from cart"
// @Failure 400 {object} map[string]interface{} "Invalid cart item ID"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /cart/{id} [delete]
func (h *cartHandler) RemoveFromCart(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Invalid cart item ID",
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
		"message": "Item removed from cart",
	})
}

// ClearCart godoc
// @Summary Clear user's cart
// @Description Remove all items from the authenticated user's cart
// @Tags cart
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Cart cleared"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /cart/clear [delete]
func (h *cartHandler) ClearCart(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
			"details": "User not authenticated",
		})
		return
	}

	err := h.service.Clear(userID.(uint))
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
		"message": "Cart cleared",
	})
}
