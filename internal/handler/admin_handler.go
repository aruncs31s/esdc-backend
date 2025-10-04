package handler

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler interface {
	GetAllUsers(c *gin.Context)
	GetUsersStats(c *gin.Context)
	DeleteUser(c *gin.Context)
	CreateUser(c *gin.Context)
}
type adminHandler struct {
	adminService   service.AdminService
	responseHelper responses.ResponseHelper
	// adminService    service.AdminService

}

func NewAdminHandler(adminService service.AdminService) AdminHandler {
	responseHelper := responses.NewResponseHelper()
	return &adminHandler{
		responseHelper: responseHelper,
		adminService:   adminService,
	}
}

func (h *adminHandler) GetAllUsers(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		h.responseHelper.Unauthorized(c, "Normal User Can not access this page.")
		return
	}
	users, err := h.adminService.GetAllUsers()
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to retrieve users", err)
		return
	}
	h.responseHelper.Success(c, users)
}

func (h *adminHandler) GetUsersStats(c *gin.Context) {
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
func (h *adminHandler) DeleteUser(c *gin.Context) {
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
func (h *adminHandler) CreateUser(c *gin.Context) {
	role := c.GetString("role")

	if role != "admin" {
		h.responseHelper.Unauthorized(c, "Normal User Can not access this page.")
		return
	}
	var user dto.RegisterRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request data", err.Error())
		return
	}
	err := h.adminService.CreateUser(user)
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to create user", err)
		return
	}
	h.responseHelper.Success(c, user)
}
