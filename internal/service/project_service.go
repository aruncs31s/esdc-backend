package service

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
	"fmt"
	"strings"
)

type ProjectService interface {
	GetAllProjects(limit, offset int, user string) ([]dto.ProjectResponse, error)
	CreateProject(user string, project dto.ProjectCreation) (*model.Project, error)
	GetProject(id int, user string) (dto.ProjectResponse, error)
	ToggleLikeProject(username string, projectID int) (bool, error) // Returns true if liked, false if unliked
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

func (s *projectService) GetAllProjects(limit, offset int, user string) ([]dto.ProjectResponse, error) {
	projects, err := s.projectRepo.GetAll(limit, offset)

	if err != nil {
		return nil, err
	}
	var userID uint
	if user != "" {
		userID, err = s.userRepo.FindUserIDByUsername(user)
		if err != nil {
			return nil, fmt.Errorf("error fetching user details: %w", err)
		}
		if userID == 0 {
			user = ""
		}
	}
	projectsPresentation := make([]dto.ProjectResponse, 0)
	for _, project := range projects {
		isLiked := false
		if user != "" {
			isLiked, _ = s.projectRepo.IsLiked(userID, project.ID)
		}
		p := dto.ProjectResponse{
			ID:                  project.ID,
			Title:               project.Title,
			Description:         project.Description,
			GithubLink:          project.GithubLink,
			Image:               project.Image,
			LiveUrl:             project.LiveUrl,
			CreatedAt:           project.CreatedAt,
			UpdatedAt:           project.UpdatedAt,
			Status:              project.Status,
			Likes:               project.Likes,
			Cost:                project.Cost,
			Category:            project.Category,
			IsLiked:             isLiked,
			CreatorDetails:      getCreatorDetails(project.Creator),
			ContributorsDetails: getContributorsUsernames(project.Contributors),
			TagsDetails:         getTagsNames(project.Tags),
			TechnologyDetails:   getTechnologiesNames(project.Technologies),
		}
		projectsPresentation = append(projectsPresentation, p)
	}
	return projectsPresentation, nil
}
func getCreatorDetails(creator model.User) dto.Contributor {
	return dto.Contributor{
		ID:    int(creator.ID),
		Name:  creator.Username,
		Email: creator.Email,
	}
}
func getContributorsUsernames(contributors *[]model.User) *[]dto.Contributor {
	if contributors == nil {
		return nil
	}
	usernames := make([]dto.Contributor, 0)
	for _, user := range *contributors {
		usernames = append(usernames, dto.Contributor{
			ID:    int(user.ID),
			Name:  user.Username,
			Email: user.Email,
		})
	}
	// add creator as a contributor also if not already present
	return &usernames
}
func getTagsNames(tags *[]model.Tag) *[]dto.Tag {
	if tags == nil {
		return nil
	}
	names := make([]dto.Tag, 0)
	for _, tag := range *tags {
		names = append(names,
			dto.Tag{
				ID:   int(tag.ID),
				Name: tag.Name,
			})
	}
	return &names
}
func getTechnologiesNames(technologies *[]model.Technologies) *[]dto.Technology {
	if technologies == nil {
		return nil
	}
	names := make([]dto.Technology, 0)

	for _, tech := range *technologies {
		names = append(names, dto.Technology{
			ID:   int(tech.ID),
			Name: tech.Name,
		})
	}
	return &names
}
func (s *projectService) CreateProject(user string, project dto.ProjectCreation) (*model.Project, error) {
	userID, err := s.userRepo.FindUserIDByUsername(user)
	if err != nil {
		return nil, err
	}
	// Build []model.User slice for contributors
	contributors := make([]model.User, 0)
	// First add the creator as a contributor
	creator, err := s.userRepo.FindByID(userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching creator details: %w", err)
	}
	contributors = append(contributors, creator)
	// now add contributors from the request
	if project.Contributers != nil && len(*project.Contributers) > 0 {
		users, err := s.userRepo.FindUsersByUsernames(*project.Contributers)
		if err != nil {
			return nil, fmt.Errorf("error fetching contributors: %w", err)
		}
		if len(users) != len(*project.Contributers) {
			return nil, fmt.Errorf("one or more contributors not found")
		}
		contributors = append(contributors, users...)
	}

	// This one , should if the tag exists in the db , if exists assign its value to the project or else create a new tag and assign it to the project
	tags := make([]model.Tag, 0)
	if project.Tags != nil {
		for _, tagName := range *project.Tags {
			tag, err := s.projectRepo.FindOrCreateTag(tagName)
			if err != nil {
				return nil, fmt.Errorf("error creating/finding tag: %w", err)
			}
			tags = append(tags, *tag)
		}
	}
	technologies := make([]model.Technologies, 0)
	if project.Technologies != nil {
		splittedTechnologies := strings.Split((*project.Technologies)[0], ",")
		for _, techName := range splittedTechnologies {
			filteredTechName := strings.TrimSpace(techName)
			tech, err := s.projectRepo.FindOrCreateTechnology(filteredTechName)
			if err != nil {
				return nil, fmt.Errorf("error creating/finding technology: %w", err)
			}
			technologies = append(technologies, *tech)
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
		ModifiedBy:   &userID,
		Status:       "active",         // Default status
		Likes:        0,                // Default value
		Category:     project.Category, // Set category from request
		LiveUrl:      project.LiveUrl,
		Technologies: &technologies,
		Contributors: &contributors,
	}

	// Save the project and its relationships
	if err := s.projectRepo.Create(&newProject); err != nil {
		return nil, err
	}

	return &newProject, nil
}

func (s *projectService) GetProject(id int, user string) (dto.ProjectResponse, error) {

	project, err := s.projectRepo.GetByID(id)

	if err != nil {
		return dto.ProjectResponse{}, err
	}

	isLiked := false
	if user != "" {
		userID, err := s.userRepo.FindUserIDByUsername(user)
		if err == nil {
			isLiked, _ = s.projectRepo.IsLiked(uint(userID), id)
		}
	}

	p := dto.ProjectResponse{
		ID:                  project.ID,
		Title:               project.Title,
		Description:         project.Description,
		GithubLink:          project.GithubLink,
		Image:               project.Image,
		LiveUrl:             project.LiveUrl,
		CreatedAt:           project.CreatedAt,
		UpdatedAt:           project.UpdatedAt,
		Status:              project.Status,
		Likes:               project.Likes,
		Cost:                project.Cost,
		Category:            project.Category,
		IsLiked:             isLiked,
		CreatorDetails:      getCreatorDetails(project.Creator),
		ContributorsDetails: getContributorsUsernames(project.Contributors),
		TagsDetails:         getTagsNames(project.Tags),
		TechnologyDetails:   getTechnologiesNames(project.Technologies),
	}
	return p, nil
}

func (s *projectService) ToggleLikeProject(username string, projectID int) (bool, error) {
	// Get the user ID from username
	userID, err := s.userRepo.FindUserIDByUsername(username)
	if err != nil {
		return false, err
	}
	// check if already liked.
	isLiked, err := s.projectRepo.IsLiked(uint(userID), projectID)
	if err != nil {
		return false, err
	}
	if isLiked {
		// Unlike
		err = s.projectRepo.UnlikeProject(uint(userID), projectID)
		return false, err
	} else {
		// Like
		err = s.projectRepo.LikeProject(uint(userID), projectID)
		return true, err
	}
}
