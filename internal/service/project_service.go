package service

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
	"fmt"
)

type ProjectService interface {
	GetAllProjects() ([]model.Project, error)
	CreateProject(user string, project dto.ProjectCreation) (*model.Project, error)
	// GetProject(id int) (model.Project, error)
	// UpdateProject(id int, project model.Project) (model.Project, error)
	// DeleteProject(id int) error
}

type projectService struct {
	projectRepo repository.ProjectRepository
	userRepo    repository.UserRepository
}

func NewProjectService(
	projectRepo repository.ProjectRepository,
	userRepo repository.UserRepository,
) ProjectService {
	return &projectService{
		projectRepo: projectRepo,
		userRepo:    userRepo,
	}
}

func (s *projectService) GetAllProjects() ([]model.Project, error) {
	return s.projectRepo.GetAll()
}
func (s *projectService) CreateProject(user string, project dto.ProjectCreation) (*model.Project, error) {
	userID, err := s.userRepo.FindUserIDByUsername(user)
	if err != nil {
		return nil, err
	}
	// Build []model.User slice for contributors
	contributors := make([]model.User, 0)
	if project.Contributers != nil {
		for _, username := range *project.Contributers {
			u, err := s.userRepo.FindByUsername(username)
			if err != nil {
				return nil, fmt.Errorf("contributor '%s' not found: %w", username, err)
			}
			contributors = append(contributors, u)
		}
	}
	tags := make([]model.Tag, 0)
	if project.Tags != nil {
		for _, tag := range *project.Tags {
			tags = append(tags, model.Tag{
				Name: tag,
			})
		}
	}
	technologies := make([]model.Technologies, 0)
	if project.Technologies != nil {
		for _, tech := range *project.Technologies {
			technologies = append(technologies, model.Technologies{
				Name: tech,
			})
		}
	}
	// Create the new project
	newProject := model.Project{
		Title:        project.Title,
		Image:        project.Image,
		Description:  project.Description,
		GithubLink:   project.GithubLink,
		Tags:         &tags,
		CreatedBy:    userID,
		LiveUrl:      project.LiveUrl,
		Technologies: &technologies,
		ModifiedBy:   &userID,
		Contributors: contributors,
	}

	// Save the project and its relationships
	if err := s.projectRepo.Create(&newProject); err != nil {
		return nil, err
	}

	return &newProject, nil
}
