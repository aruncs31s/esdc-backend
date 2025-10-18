package service

import (
	"esdc-backend/internal/module/common/model"
	"esdc-backend/internal/module/common/repository"
)

type OrderService struct {
	orderRepo   repository.OrderRepository
	productRepo repository.ProductRepository
}

func NewOrderService(orderRepo repository.OrderRepository, productRepo repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) GetByUserID(userID uint, limit, offset int) ([]model.Order, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	orders, err := s.orderRepo.GetByUserID(userID, limit, offset)
	if err != nil {
		return nil, err
	}

	return orders, nil
}

func (s *OrderService) GetByID(id, userID uint) (*model.Order, error) {
	order, err := s.orderRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) Create(order *model.Order) error {
	return s.orderRepo.Create(order)
}

func (s *OrderService) Update(id uint, order *model.Order) error {
	return s.orderRepo.Update(id, order)
}

func (s *OrderService) Delete(id, userID uint) error {
	return s.orderRepo.Delete(id, userID)
}
