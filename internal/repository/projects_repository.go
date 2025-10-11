package repository

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"

	"gorm.io/gorm"
)

type ProjectRepository interface {
	GetAll() ([]model.Project, error)
	GetEssentialInfo(limit, offset int) ([]dto.ProjectsEssentialInfo, error)
	Create(project *model.Project) error
	GetByID(id int) (model.Project, error)
	// Update(id int, project model.Project) (model.Project, error)
	// Delete(id int) error
	GetProjectsCount() (int, error)
	FindOrCreateTag(name string) (*model.Tag, error)
	FindOrCreateTechnology(name string) (*model.Technologies, error)
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// For Admin Pannel
func (r *projectRepository) GetEssentialInfo(limit, offset int) ([]dto.ProjectsEssentialInfo, error) {
	var projects []dto.ProjectsEssentialInfo
	query := r.db.Model(&model.Project{}).Select("projects.id, projects.title, users.name as created_by, projects.status").
		Joins("LEFT JOIN users ON projects.created_by = users.id").
		Limit(limit).Offset(offset)
	if err := query.Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *projectRepository) GetAll() ([]model.Project, error) {
	var projects []model.Project
	if err := r.db.Preload("Contributors").Preload("Creator").Preload("Tags").Preload("Technologies").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
func (r *projectRepository) GetByID(id int) (model.Project, error) {
	var project model.Project
	if err := r.db.Preload("Contributors").Preload("Creator").Preload("Tags").Preload("Technologies").First(&project, id).Error; err != nil {
		return model.Project{}, err
	}
	return project, nil
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

func (r *projectRepository) FindOrCreateTag(name string) (*model.Tag, error) {
	var tag model.Tag
	if err := r.db.Where("name = ?", name).FirstOrCreate(&tag, model.Tag{Name: name}).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *projectRepository) FindOrCreateTechnology(name string) (*model.Technologies, error) {
	var tech model.Technologies
	if err := r.db.Where("name = ?", name).FirstOrCreate(&tech, model.Technologies{Name: name}).Error; err != nil {
		return nil, err
	}
	return &tech, nil
}
