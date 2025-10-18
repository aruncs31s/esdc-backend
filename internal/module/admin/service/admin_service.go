package service

import (
	admin_dto "esdc-backend/internal/module/admin/dto"
	"esdc-backend/internal/module/admin/service/mapper"
	common_dto "esdc-backend/internal/module/common/dto"
	common_model "esdc-backend/internal/module/common/model"
	common_repo "esdc-backend/internal/module/common/repository"
	user_repo "esdc-backend/internal/module/user/repository"
	"time"
)

type AdminService interface {
	// GetProjectsEssentialInfo retrieves essential information about projects for admin panel
	// Example response structure:
	// {
	//   "id": 1,
	//   "title": "Project Title",
	//   "created_by": "Arun C S",
	//   "status": "Archived",
	//   "created_at": "2023-01-01 12:00:00",
	//   "updated_at": "2023-01-02 12:00:00"
	// }
	GetProjectsEssentialInfo(limit, offset int) ([]common_dto.ProjectsEssentialInfo, error)
	GetAllUsers() (*[]admin_dto.UserDataForAdmin, error)
	GetUsersStats() (*common_dto.UsersStats, error)
	DeleteUser(userID int) error
	CreateUser(user admin_dto.AdminRegisterRequest) error
}

type adminService struct {
	userRepo    user_repo.UserRepository
	projectRepo common_repo.ProjectRepository
}

func NewAdminService(userRepo user_repo.UserRepository, projectRepo common_repo.ProjectRepository) AdminService {
	return &adminService{
		userRepo:    userRepo,
		projectRepo: projectRepo,
	}
}

func (s *adminService) GetProjectsEssentialInfo(limit, offset int) ([]common_dto.ProjectsEssentialInfo, error) {
	projects, err := s.projectRepo.GetEssentialInfo(limit, offset)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (s *adminService) GetAllUsers() (*[]admin_dto.UserDataForAdmin, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	filteredUsers := mapper.MapToUserDataForAdmin(users)
	return filteredUsers, nil
}

func (s *adminService) GetUsersStats() (*common_dto.UsersStats, error) {

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

	data := getUserStats(usersCount, projectsCount, totalChallenges, activeUser)

	return &data, nil
}

func getUserStats(usersCount int, projectsCount int, totalChallenges int, activeUser int) common_dto.UsersStats {
	data := common_dto.UsersStats{
		TotalUsers:      usersCount,
		TotalProjects:   projectsCount,
		TotalChallenges: totalChallenges,
		ActiveUsers:     activeUser,
	}
	return data
}
func (s *adminService) DeleteUser(userID int) error {
	err := s.userRepo.DeleteUserByID(uint(userID))
	if err != nil {
		return err
	}
	return nil
}
func (s *adminService) CreateUser(user admin_dto.AdminRegisterRequest) error {
	newUser := getUserData(user)
	err := s.userRepo.CreateUser(&newUser)
	if err != nil {
		return err
	}
	return nil
}

func getUserData(user admin_dto.AdminRegisterRequest) common_model.User {
	newUser := common_model.User{
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
		Password: user.Password,
		Github: &common_model.Github{
			Username: user.GithubUsername,
		},
		CreatedAt: time.Time{}.Unix(),
		UpdatedAt: time.Time{}.Unix(),
	}
	return newUser
}
