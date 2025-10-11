package service

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
	"time"
)

type AdminService interface {
	GetAllProjects(limit, offset int) ([]dto.ProjectsEssentialInfo, error)
	GetAllUsers() ([]model.User, error)
	GetUsersStats() (*dto.UsersStats, error)
	DeleteUser(userID int) error
	CreateUser(user dto.AdminRegisterRequest) error
}

type adminService struct {
	userRepo    repository.UserRepository
	projectRepo repository.ProjectRepository
}

func NewAdminService(userRepo repository.UserRepository, projectRepo repository.ProjectRepository) AdminService {
	return &adminService{
		userRepo:    userRepo,
		projectRepo: projectRepo,
	}
}
func (s *adminService) GetAllProjects(limit, offset int) ([]dto.ProjectsEssentialInfo, error) {
	projects, err := s.projectRepo.GetEssentialInfo(limit, offset)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (s *adminService) GetAllUsers() ([]model.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (s *adminService) GetUsersStats() (*dto.UsersStats, error) {

	usersCount, err := s.userRepo.GetUsersCount()
	if err != nil {
		return nil, err
	}
	projectsCount, err := s.projectRepo.GetProjectsCount()
	if err != nil {
		return nil, err
	}
	var activeUser = 2
	var totalChallenges = 5

	data := dto.UsersStats{
		TotalUsers:      usersCount,
		TotalProjects:   projectsCount,
		TotalChallenges: totalChallenges,
		ActiveUsers:     activeUser,
	}

	return &data, nil
}
func (s *adminService) DeleteUser(userID int) error {
	err := s.userRepo.DeleteUserByID(uint(userID))
	if err != nil {
		return err
	}
	return nil
}
func (s *adminService) CreateUser(user dto.AdminRegisterRequest) error {
	newUser := model.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
		Github: &model.Github{
			Username: user.GithubUsername,
		},
		CreatedAt: time.Time{}.Unix(),
		UpdatedAt: time.Time{}.Unix(),
	}
	err := s.userRepo.CreateUser(&newUser)
	if err != nil {
		return err
	}
	return nil
}
