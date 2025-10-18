package handler

import (
	"esdc-backend/internal/module/common/model"
	"esdc-backend/internal/module/common/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(service *service.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// GetOrders godoc
// @Summary Get user's orders
// @Description Get all orders for the authenticated user with pagination
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param limit query int false "Limit number of results" default(20)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} map[string]interface{} "Orders retrieved successfully"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /orders [get]
func (h *OrderHandler) GetOrders(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
			"details": "User not authenticated",
		})
		return
	}

	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	orders, err := h.service.GetByUserID(userID.(uint), limit, offset)
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
		"data":    orders,
	})
}

// GetOrderByID godoc
// @Summary Get order by ID
// @Description Get a specific order by its ID for the authenticated user
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{} "Order retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid order ID"
// @Failure 401 {object} map[string]interface{} "User not authenticated or unauthorized access"
// @Failure 404 {object} map[string]interface{} "Order not found"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
			"details": "User not authenticated",
		})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Invalid order ID",
		})
		return
	}

	order, err := h.service.GetByID(uint(id), userID.(uint))
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "unauthorized access to order" {
			statusCode = http.StatusUnauthorized
		} else if err.Error() == "record not found" {
			statusCode = http.StatusNotFound
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   "Error retrieving order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    order,
	})
}

// CreateOrder godoc
// @Summary Create a new order
// @Description Create a new order with the specified items
// @Tags orders
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body object{items=[]object{product_id=uint,quantity=int}} true "Create order request"
// @Success 200 {object} map[string]interface{} "Order created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request data or insufficient stock"
// @Failure 401 {object} map[string]interface{} "User not authenticated"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(c *gin.Context) {
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
		Items []struct {
			ProductID uint `json:"product_id" binding:"required"`
			Quantity  int  `json:"quantity" binding:"required"`
		} `json:"items" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Items are required",
		})
		return
	}

	// Build the order with items
	var orderItems []model.OrderItem
	var total float64

	for _, item := range req.Items {
		orderItems = append(orderItems, model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}

	order := &model.Order{
		UserID: userID.(uint),
		Items:  orderItems,
		Total:  total,
		Status: "pending",
	}

	err := h.service.Create(order)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if err.Error() == "product not found" {
			statusCode = http.StatusNotFound
		} else if err.Error() == "order must contain at least one item" ||
			(len(err.Error()) > 19 && err.Error()[:19] == "insufficient stock") {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"success": false,
			"error":   "Error creating order",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Order created successfully",
		"data":    order,
	})
}
