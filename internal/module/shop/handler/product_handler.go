package handler

import (
	service "esdc-backend/internal/module/shop/service"
	"esdc-backend/utils"
	"strconv"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type ProductHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type productHandler struct {
	service        service.ProductService
	responseHelper responsehelper.ResponseHelper
}

func NewProductHandler(service service.ProductService) ProductHandler {
	responseHelper := responsehelper.NewResponseHelper()
	return &productHandler{
		service:        service,
		responseHelper: responseHelper,
	}
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
func (h *productHandler) GetAll(c *gin.Context) {
	category := c.Query("category")
	search := c.Query("search")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	products, total, err := h.service.GetAll(category, search, limit, offset)
	if err != nil {
		h.responseHelper.InternalError(c, "Internal server error", err)
		return
	}
	response := map[string]interface{}{
		"products": products,
		"total":    total,
	}
	h.responseHelper.Success(c, response)
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
func (h *productHandler) GetByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		h.responseHelper.BadRequest(c, utils.ErrBadRequest.Error(), "invalid id")
		return
	}

	product, err := h.service.GetByID(uint(id))
	if err != nil {
		h.responseHelper.NotFound(c, utils.ErrNotFound.Error())
		return
	}

	response := map[string]interface{}{
		"product": product,
	}
	h.responseHelper.Success(c, response)
}
