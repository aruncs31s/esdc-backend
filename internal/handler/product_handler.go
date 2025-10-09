package handler

import (
	"esdc-backend/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

// GetAll godoc
// @Summary Get all products
// @Description Get all products with optional filtering by category and search
// @Tags products
// @Accept json
// @Produce json
// @Param category query string false "Product category"
// @Param search query string false "Search term for name or description"
// @Param limit query int false "Limit number of results" default(50)
// @Param offset query int false "Offset for pagination" default(0)
// @Success 200 {object} map[string]interface{} "Products retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /products [get]
func (h *ProductHandler) GetAll(c *gin.Context) {
	category := c.Query("category")
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	products, total, err := h.service.GetAll(category, search, limit, offset)
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
		"data":    products,
		"total":   total,
	})
}

// GetByID godoc
// @Summary Get product by ID
// @Description Get a specific product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} map[string]interface{} "Product retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Invalid product ID"
// @Failure 404 {object} map[string]interface{} "Product not found"
// @Router /products/{id} [get]
func (h *ProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request data",
			"details": "Invalid product ID",
		})
		return
	}

	product, err := h.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Not found",
			"details": "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}
