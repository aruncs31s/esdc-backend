package service

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
)

type ProjectService interface {
	GetAllProjects() ([]model.Project, error)
	CreateProject(project dto.ProjectCreation) (model.Project, error)
	// GetProject(id int) (model.Project, error)
	// UpdateProject(id int, project model.Project) (model.Project, error)
	// DeleteProject(id int) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
}

func NewProjectService(projectRepo repository.ProjectRepository) ProjectService {
	return &projectService{projectRepo: projectRepo}
}

func (s *projectService) GetAllProjects() ([]model.Project, error) {
	return s.projectRepo.GetAll()
}

func (s *projectService) CreateProject(project dto.ProjectCreation) (model.Project, error) {
	newProject := model.Project{
		Name:         project.Name,
		Image:        project.Image,
		Description:  project.Description,
		Link:         project.Link,
		Tags:         model.JSONSlice(project.Tags),
		Contributers: model.JSONInts(project.Contributers),
	}
	return s.projectRepo.Create(&newProject)
}
