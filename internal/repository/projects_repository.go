package repository

import (
	"esdc-backend/internal/model"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAll() ([]model.Project, error)
	Create(project *model.Project) error
	// GetByID(id int) (model.Project, error)
	// Update(id int, project model.Project) (model.Project, error)
	// Delete(id int) error
	GetProjectsCount() (int, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

func (r *projectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project
	if err := r.db.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *projectRepository) Create(project *model.Project) error {
	if err := r.db.Create(project).Error; err != nil {
		return err
	}
	return nil
}
func (r *projectRepository) GetProjectsCount() (int, error) {
	var count int64
	result := r.db.Model(&model.Project{}).Count(&count)
	return int(count), result.Error
}
