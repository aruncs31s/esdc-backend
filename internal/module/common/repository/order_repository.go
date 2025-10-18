package repository

import (
	"esdc-backend/internal/module/common/model"

	"gorm.io/gorm"
)

// OrderRepository defines the interface for order data operations
type OrderRepository interface {
	GetByUserID(userID uint, limit, offset int) ([]model.Order, error)
	GetByID(id, userID uint) (*model.Order, error)
	Create(order *model.Order) error
	Update(id uint, order *model.Order) error
	Delete(id, userID uint) error
}

// orderRepository is the implementation of OrderRepository
type orderRepository struct {
	db *gorm.DB
}

// NewOrderRepository creates a new order repository
func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetByUserID(userID uint, limit, offset int) ([]model.Order, error) {
	var orders []model.Order
	err := r.db.Where("user_id = ?", userID).Limit(limit).Offset(offset).Find(&orders).Error
	return orders, err
}

func (r *orderRepository) GetByID(id, userID uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&order).Error
	return &order, err
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) Update(id uint, order *model.Order) error {
	return r.db.Model(&model.Order{}).Where("id = ?", id).Updates(order).Error
}

func (r *orderRepository) Delete(id, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&model.Order{}).Error
}
