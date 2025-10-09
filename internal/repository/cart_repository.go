package repository

import (
	"esdc-backend/internal/model"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{db: db}
}

func (r *CartRepository) GetByUserID(userID uint) ([]model.Cart, error) {
	var carts []model.Cart
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *CartRepository) GetByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Preload("Product").First(&cart, id).Error
	return &cart, err
}

func (r *CartRepository) FindByUserAndProduct(userID, productID uint) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	return &cart, err
}

func (r *CartRepository) Add(cart *model.Cart) error {
	return r.db.Create(cart).Error
}

func (r *CartRepository) Update(cart *model.Cart) error {
	return r.db.Save(cart).Error
}

func (r *CartRepository) Delete(id uint) error {
	return r.db.Delete(&model.Cart{}, id).Error
}

func (r *CartRepository) Clear(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}
