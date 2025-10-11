package handler

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/handler/responses"
	"esdc-backend/internal/service"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProjectHandler interface {
	GetAllProjects(c *gin.Context)
	CreateProject(c *gin.Context)
	GetProject(c *gin.Context)
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

// GetAllProjects godoc
// @Summary Get all projects
// @Description Retrieve all projects
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{} "Projects retrieved successfully"
// @Failure 404 {object} map[string]interface{} "No projects found"
// @Router /projects [get]
func (h *projectHandler) GetAllProjects(c *gin.Context) {
	projects, err := h.projectService.GetAllProjects()
	if err != nil {
		h.responseHelper.NotFound(c, "No projects found")
		return
	}
	h.responseHelper.Success(c, projects)
}

// CreateProject godoc
// @Summary Create a new project
// @Description Create a new project with the provided data
// @Tags projects
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param project body dto.ProjectCreation true "Project creation data"
// @Success 201 {object} map[string]interface{} "Project created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request body"
// @Failure 500 {object} map[string]interface{} "Failed to create project"
// @Router /projects [post]
func (h *projectHandler) CreateProject(c *gin.Context) {
	user := c.GetString("user")
	if user == "" {
		h.responseHelper.Unauthorized(c, "User not authenticated")
		return
	}
	var project dto.ProjectCreation
	if err := c.ShouldBindJSON(&project); err != nil {
		h.responseHelper.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	createdProject, err := h.projectService.CreateProject(user, project)
	if err != nil && strings.Contains(err.Error(), "UNIQUE constraint failed") {
		h.responseHelper.Conflict(c, "Project with the same name already exists", err.Error())
		return
	}
	if err != nil {
		h.responseHelper.InternalError(c, "Failed to create project", err)
		return
	}
	h.responseHelper.Created(c, createdProject)
}
func (h *projectHandler) GetProject(c *gin.Context) {
	// user := c.GetString("user")
	idStr := c.Param("id")
	if idStr == "" {
		h.responseHelper.BadRequest(c, "Project ID is required", "invalid id")
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.responseHelper.BadRequest(c, "Invalid project ID", err.Error())
		return
	}

	project, err := h.projectService.GetProject(id)
	if err != nil {
		h.responseHelper.NotFound(c, "Project not found")
		return
	}
	h.responseHelper.Success(c, project)
}
