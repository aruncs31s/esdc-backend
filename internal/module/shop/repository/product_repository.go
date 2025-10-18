package repository

import (
	"esdc-backend/internal/module/shop/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAll(category, search string, limit, offset int) ([]model.Product, int64, error)
	GetByID(id uint) (*model.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAll(category, search string, limit, offset int) ([]model.Product, int64, error) {
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

func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}
