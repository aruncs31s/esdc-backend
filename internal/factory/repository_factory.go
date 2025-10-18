package factory

import (
	"esdc-backend/internal/module/auth/repository"
	authRepo "esdc-backend/internal/module/auth/repository"
	projectRepo "esdc-backend/internal/module/common/repository"
	userRepo "esdc-backend/internal/module/user/repository"

	"gorm.io/gorm"
)

type RepositoryFactory interface {
	GetAuthRepository() authRepo.AuthRepository
	GetUserRepository() userRepo.UserRepository
	GetProjectRepository() projectRepo.ProjectRepository
}

type repositoryFactory struct {
	db *gorm.DB
}

func NewRepositoryFactory(db *gorm.DB) RepositoryFactory {
	return &repositoryFactory{db: db}
}

func (f *repositoryFactory) GetAuthRepository() authRepo.AuthRepository {
	return repository.NewAuthRepository(
		f.db,
	)
}
func (f *repositoryFactory) GetUserRepository() userRepo.UserRepository {
	return userRepo.NewUserRepository(
		f.db,
	)
}
func (f *repositoryFactory) GetProjectRepository() projectRepo.ProjectRepository {
	return projectRepo.NewProjectRepository(
		f.db,
	)
}
