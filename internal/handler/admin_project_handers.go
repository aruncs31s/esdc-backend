package handler

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type AdminProjectHandlers interface {
	GetAllProjects(c *gin.Context)
	GetProjectByID(c *gin.Context)
	CreateProject(c *gin.Context)
	UpdateProject(c *gin.Context)
	DeleteProject(c *gin.Context)
}

type adminProjectHandler struct {
	projectService service.ProjectService
	responseHelper responses.ResponseHelper
}

func newAdminProjectHandler(projectService service.ProjectService) AdminProjectHandlers {
	responseHelper := responses.NewResponseHelper()
	return &adminProjectHandler{
		responseHelper: responseHelper,
		projectService: projectService,
	}
}

func verifyAdminRole(c *gin.Context, responseHelper responses.ResponseHelper) bool {
	role := c.GetString("role")
	if role != "admin" {
		responseHelper.Unauthorized(c, "Admin role required. Your role: "+role)
		return false
	}
	return true
}
func (h *adminProjectHandler) GetAllProjects(c *gin.Context) {
	if !verifyAdminRole(c, h.responseHelper) {
		return
	}
	allProjects, err := h.projectService.GetAllProjects()
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to retrieve projects", err)
		return
	}
	h.responseHelper.Success(c, allProjects)

}
func (h *adminProjectHandler) GetProjectByID(c *gin.Context) {
	// Implementation here
}
func (h *adminProjectHandler) CreateProject(c *gin.Context) {
	// Implementation here
	if !verifyAdminRole(c, h.responseHelper) {
		return
	}
	userName := c.GetString("user")
	if userName == "" {
		h.responseHelper.Unauthorized(c, "User not authenticated")
		return
	}
	var projectDTO dto.ProjectCreation
	if err := c.ShouldBindJSON(&projectDTO); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request payload", err.Error())
		return
	}

	createdProject, err := h.projectService.CreateProject(userName, projectDTO)

	if err != nil || createdProject == nil { // Error while creating project
		h.responseHelper.InternalError(c, "Failed to create project", err)
		return
	}
	h.responseHelper.Success(c, createdProject)

}
func (h *adminProjectHandler) UpdateProject(c *gin.Context) {
	// Implementation here
}
func (h *adminProjectHandler) DeleteProject(c *gin.Context) {
	// Implementation here
}
