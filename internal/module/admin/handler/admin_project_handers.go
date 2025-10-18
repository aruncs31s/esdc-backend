package handler

import (
	"esdc-backend/internal/module/common/dto"

	"esdc-backend/internal/module/admin/service"
	common_Service "esdc-backend/internal/module/common/service"
	"strconv"

	"github.com/aruncs31s/responsehelper"
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
	adminService   service.AdminService
	projectService common_Service.ProjectService
	responseHelper responsehelper.ResponseHelper
}

func newAdminProjectHandler(projectService common_Service.ProjectService, adminService service.AdminService) AdminProjectHandlers {
	responseHelper := responsehelper.NewResponseHelper()
	return &adminProjectHandler{
		responseHelper: responseHelper,
		projectService: projectService,
		adminService:   adminService,
	}
}

func verifyAdminRole(c *gin.Context, responseHelper responsehelper.ResponseHelper) bool {
	role := c.GetString("role")
	if role != "admin" {
		responseHelper.Unauthorized(c, "Admin role required. Your role: "+role)
		return false
	}
	return true
}
func getPaginationParams(c *gin.Context) (string, string) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "50")
	return pageStr, limitStr
}
func getPaginationParamsInt(c *gin.Context) (int, int) {
	page, limit := getPaginationParams(c)
	pageInt := 1
	limitInt := 50
	if p, err := strconv.Atoi(page); err == nil && p > 0 {
		pageInt = p
	}
	if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 1000 {
		limitInt = l
	}
	return limitInt, (pageInt - 1) * limitInt
}

func (h *adminProjectHandler) GetAllProjects(c *gin.Context) {
	if !verifyAdminRole(c, h.responseHelper) {
		return
	}
	limit, offset := getPaginationParamsInt(c)

	allProjects, err := h.adminService.GetProjectsEssentialInfo(limit, offset)
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
