package handler

import (
	"esdc-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminHandler interface {
	AdminUserHandlers
	AdminProjectHandlers
}

type adminHandler struct {
	projectHandler AdminProjectHandlers
	userHandler    AdminUserHandlers
}

func NewAdminHandler(adminService service.AdminService, projectService service.ProjectService) AdminHandler {
	adminProjectHandler := newAdminProjectHandler(projectService)
	adminUserHandler := newAdminUserHandler(adminService)
	return &adminHandler{
		projectHandler: adminProjectHandler,
		userHandler:    adminUserHandler,
	}
}

func (h *adminHandler) GetAllUsers(c *gin.Context) {
	h.userHandler.GetAllUsers(c)
}
func (h *adminHandler) GetUsersStats(c *gin.Context) {
	h.userHandler.GetUsersStats(c)
}
func (h *adminHandler) DeleteUser(c *gin.Context) {
	h.userHandler.DeleteUser(c)
}
func (h *adminHandler) CreateUser(c *gin.Context) {
	h.userHandler.CreateUser(c)
}
func (h *adminHandler) GetAllProjects(c *gin.Context) {
	h.projectHandler.GetAllProjects(c)
}
func (h *adminHandler) GetProjectByID(c *gin.Context) {
	h.projectHandler.GetProjectByID(c)
}
func (h *adminHandler) CreateProject(c *gin.Context) {
	h.projectHandler.CreateProject(c)
}
func (h *adminHandler) UpdateProject(c *gin.Context) {
	h.projectHandler.UpdateProject(c)
}
func (h *adminHandler) DeleteProject(c *gin.Context) {
	h.projectHandler.DeleteProject(c)
}
