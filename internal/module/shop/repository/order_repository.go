package repository

import (
	"esdc-backend/internal/module/shop/model"

	"gorm.io/gorm"
)

type OrderRepository interface {
	GetByUserID(userID uint, limit, offset int) ([]model.Order, int64, error)
	GetByID(id uint) (*model.Order, error)
	Create(order *model.Order) error
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetByUserID(userID uint, limit, offset int) ([]model.Order, int64, error) {
	var orders []model.Order
	var total int64

	query := r.db.Preload("Items.Product").Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := query.Order("created_at DESC").Limit(limit).Offset(offset).Find(&orders).Error
	return orders, total, err
}

func (r *orderRepository) GetByID(id uint) (*model.Order, error) {
	var order model.Order
	err := r.db.Preload("Items.Product").First(&order, id).Error
	return &order, err
}

func (r *orderRepository) Create(order *model.Order) error {
	return r.db.Create(order).Error
}
