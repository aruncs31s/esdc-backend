package service

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
	"time"
)

type AdminService interface {
	GetAllProjects(limit, offset int) ([]dto.ProjectsEssentialInfo, error)
	GetAllUsers() ([]dto.UserDataForAdmin, error)
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
func (s *adminService) GetAllUsers() ([]dto.UserDataForAdmin, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	var filteredUsers []dto.UserDataForAdmin
	for _, user := range users {
		filteredUsers = append(filteredUsers, dto.UserDataForAdmin{
			ID:             user.ID,
			Name:           user.Name,
			Email:          user.Email,
			Username:       user.Username,
			GithubUsername: user.Github.Username,
			Role:           user.Role,
			Status:         user.Status,
			CreatedAt:      getCreatedDateFromNumber(user.CreatedAt),
			UpdatedAt:      getCreatedDateFromNumber(user.UpdatedAt),
		})
	}
	return filteredUsers, nil
}

func getCreatedDateFromNumber(timestamp int64) string {
	t := time.Unix(timestamp, 0)
	return t.Format("2006-01-02 15:04:05")
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
