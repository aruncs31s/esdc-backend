package repository

import (
	"esdc-backend/internal/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAll(category, search string, limit, offset int) ([]model.Product, int64, error) {
	var products []model.Product
	var total int64

	query := r.db.Model(&model.Product{})

	if category != "" {
		query = query.Where("category = ?", category)
	}

	if search != "" {
		query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Limit(limit).Offset(offset).Find(&products).Error
	return products, total, err
}

func (r *ProductRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}
