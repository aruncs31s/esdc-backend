package factory

import (
	authHandler "esdc-backend/internal/module/auth/handler"
	projectHandler "esdc-backend/internal/module/common/handler"

	"gorm.io/gorm"
)

type HandlerFactory interface {
	GetAuthHandler() authHandler.AuthHandler
	GetProjectHandler() projectHandler.ProjectHandler
}

type handlerFactory struct {
	serviceFactory ServiceFactory
}

func NewHandlerFactory(db *gorm.DB) HandlerFactory {
	repositoryFactory := NewRepositoryFactory(db)
	serviceFactory := NewServiceFactory(repositoryFactory)
	return &handlerFactory{
		serviceFactory: serviceFactory,
	}
}

func (f *handlerFactory) GetAuthHandler() authHandler.AuthHandler {
	return authHandler.NewAuthHandler(
		f.serviceFactory.GetAuthService(),
	)
}
func (f *handlerFactory) GetProjectHandler() projectHandler.ProjectHandler {
	return projectHandler.NewProjectHandler(
		f.serviceFactory.GetProjectService(),
	)
}
