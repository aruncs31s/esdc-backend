package factory

import (
	"esdc-backend/internal/module/auth/service"
	authService "esdc-backend/internal/module/auth/service"
	projectService "esdc-backend/internal/module/common/service"
)

type ServiceFactory interface {
	GetAuthService() authService.AuthService
	GetProjectService() projectService.ProjectService
}

type serviceFactory struct {
	repoFactory RepositoryFactory
}

func NewServiceFactory(repositoryFactory RepositoryFactory) ServiceFactory {
	return &serviceFactory{
		repoFactory: repositoryFactory,
	}
}

func (f *serviceFactory) GetAuthService() service.AuthService {
	return service.NewAuthService(
		f.repoFactory.GetAuthRepository(),
		f.repoFactory.GetUserRepository(),
		service.NewJWTService(),
	)
}
func (f *serviceFactory) GetProjectService() projectService.ProjectService {
	return projectService.NewProjectService(
		f.repoFactory.GetProjectRepository(),
		f.repoFactory.GetUserRepository(),
	)
}
