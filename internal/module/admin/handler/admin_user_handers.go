package handler

import (
	"esdc-backend/internal/module/admin/dto"
	"esdc-backend/internal/module/admin/service"
	"strconv"

	"github.com/aruncs31s/responsehelper"
	"github.com/gin-gonic/gin"
)

type AdminUserHandlers interface {
	GetAllUsers(c *gin.Context)
	GetUsersStats(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
}
type adminUserHandler struct {
	adminService   service.AdminService
	responseHelper responsehelper.ResponseHelper
}

func newAdminUserHandler(adminService service.AdminService) AdminUserHandlers {
	responseHelper := responsehelper.NewResponseHelper()
	return &adminUserHandler{
		responseHelper: responseHelper,
		adminService:   adminService,
	}
}

// GetAllUsers godoc
// @Summary Get all users (Admin only)
// @Description Retrieve all users - requires admin role
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Users retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized - admin role required"
// @Failure 500 {object} map[string]interface{} "Failed to retrieve users"
// @Router /admin/users [get]
func (h *adminUserHandler) GetAllUsers(c *gin.Context) {
	role := c.GetString("role")
	username := c.GetString("username")

	println("🔍 Admin Handler - GetAllUsers:")
	println("  User:", username)
	println("  Role:", role)

	if role != "admin" {
		h.responseHelper.Unauthorized(c, "Admin role required. Your role: "+role)
		return
	}
	users, err := h.adminService.GetAllUsers()
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to retrieve users", err)
		return
	}
	h.responseHelper.Success(c, users)
}

// GetUsersStats godoc
// @Summary Get user statistics (Admin only)
// @Description Get statistics about users - requires admin role
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "User stats retrieved successfully"
// @Failure 401 {object} map[string]interface{} "Unauthorized - admin role required"
// @Failure 500 {object} map[string]interface{} "Failed to retrieve user stats"
// @Router /admin/users/stats [get]
func (h *adminUserHandler) GetUsersStats(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		h.responseHelper.Unauthorized(c, "Normal User Can not access this page.")
		return
	}
	stats, err := h.adminService.GetUsersStats()
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to retrieve users stats", err)
		return
	}
	h.responseHelper.Success(c, stats)
}

// DeleteUser godoc
// @Summary Delete user (Admin only)
// @Description Delete a user by ID - requires admin role
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path int true "User ID"
// @Success 200 {object} map[string]interface{} "User deleted successfully"
// @Failure 400 {object} map[string]interface{} "Invalid user ID"
// @Failure 401 {object} map[string]interface{} "Unauthorized - admin role required"
// @Failure 500 {object} map[string]interface{} "Failed to delete user"
// @Router /admin/users/{id} [delete]
func (h *adminUserHandler) DeleteUser(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		h.responseHelper.Unauthorized(c, "Normal User Can not access this page.")
		return
	}
	id := c.Param("id")
	// log.Fatal("id", id)
	userID, err := strconv.Atoi(id)
	if err != nil {
		h.responseHelper.BadRequest(c, "Invalid user ID", err.Error())
		return
	}
	if err := h.adminService.DeleteUser(userID); err != nil {
		h.responseHelper.InternalError(c, "Failed to delete user", err)
		return
	}
	h.responseHelper.Success(c, gin.H{"message": "User deleted successfully"})
}

// CreateUser godoc
// @Summary Create user (Admin only)
// @Description Create a new user - requires admin role
// @Tags admin
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param user body dto.AdminRegisterRequest true "User creation data"
// @Success 200 {object} map[string]interface{} "User created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request data"
// @Failure 401 {object} map[string]interface{} "Unauthorized - admin role required"
// @Failure 500 {object} map[string]interface{} "Failed to create user"
// @Router /admin/users [post]
func (h *adminUserHandler) CreateUser(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		h.responseHelper.Unauthorized(c, "Normal User Can not access this page.")
		return
	}
	var user dto.AdminRegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request data", err.Error())
		return
	}
	err := h.adminService.CreateUser(user)
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to create user", err)
		return
	}
	h.responseHelper.Success(c, "User created successfully")
}
