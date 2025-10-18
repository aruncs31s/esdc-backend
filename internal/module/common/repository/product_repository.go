package repository

import (
	"esdc-backend/internal/module/shop/model"

	"gorm.io/gorm"
)

// ProductRepository defines the interface for product data operations
type ProductRepository interface {
	GetByID(id uint) (*model.Product, error)
	GetAll(limit, offset int) ([]model.Product, error)
	Create(product *model.Product) error
	Update(id uint, product *model.Product) error
	Delete(id uint) error
}

// productRepository is the implementation of ProductRepository
type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new product repository
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetByID(id uint) (*model.Product, error) {
	var product model.Product
	err := r.db.First(&product, id).Error
	return &product, err
}

func (r *productRepository) GetAll(limit, offset int) ([]model.Product, error) {
	var products []model.Product
	err := r.db.Limit(limit).Offset(offset).Find(&products).Error
	return products, err
}

func (r *productRepository) Create(product *model.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(id uint, product *model.Product) error {
	return r.db.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&model.Product{}, id).Error
}
