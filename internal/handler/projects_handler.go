package handler

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ProjectHandler interface {
	GetAllProjects(c *gin.Context)
	CreateProject(c *gin.Context)
	// GetProject(c *gin.Context)
	// UpdateProject(c *gin.Context)
	// DeleteProject(c *gin.Context)
}
type projectHandler struct {
	responseHelper responses.ResponseHelper
	projectService service.ProjectService
}

func NewProjectHandler(projectService service.ProjectService) ProjectHandler {
	responseHelper := responses.NewResponseHelper()
	return &projectHandler{
		responseHelper: responseHelper,
		projectService: projectService,
	}
}

func (h *projectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.projectService.GetAllProjects()
	if err != nil {
		h.responseHelper.NotFound(c, "No projects found")
		return
	}
	h.responseHelper.Success(c, projects)
}

func (h *projectHandler) CreateProject(c *gin.Context) {
	var project dto.ProjectCreation
	if err := c.ShouldBindJSON(&project); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	createdProject, err := h.projectService.CreateProject(project)
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to create project", err)
		return
	}
	h.responseHelper.Created(c, createdProject)
}
