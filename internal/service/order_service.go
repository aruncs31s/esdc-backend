package service

import (
	"errors"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"

	"gorm.io/gorm"
)

type OrderService struct {
	orderRepo   *repository.OrderRepository
	productRepo *repository.ProductRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, productRepo *repository.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) GetByUserID(userID uint, limit, offset int) ([]model.Order, int64, error) {
	if limit <= 0 {
		limit = 20
	}
	if offset < 0 {
		offset = 0
	}

	orders, total, err := s.orderRepo.GetByUserID(userID, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// Populate product names in order items
	for i := range orders {
		for j := range orders[i].Items {
			orders[i].Items[j].Name = orders[i].Items[j].Product.Name
		}
	}

	return orders, total, nil
}

func (s *OrderService) GetByID(id, userID uint) (*model.Order, error) {
	order, err := s.orderRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Ensure user owns this order
	if order.UserID != userID {
		return nil, errors.New("unauthorized access to order")
	}

	// Populate product names in order items
	for i := range order.Items {
		order.Items[i].Name = order.Items[i].Product.Name
	}

	return order, nil
}

func (s *OrderService) Create(userID uint, items []struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required"`
}) (*model.Order, error) {
	if len(items) == 0 {
		return nil, errors.New("order must contain at least one item")
	}

	var total float64
	var orderItems []model.OrderItem

	// Validate and calculate
	for _, item := range items {
		product, err := s.productRepo.GetByID(item.ProductID)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("product not found")
			}
			return nil, err
		}

		if product.Stock < item.Quantity {
			return nil, errors.New("insufficient stock for product: " + product.Name)
		}

		itemTotal := product.Price * float64(item.Quantity)
		total += itemTotal

		orderItems = append(orderItems, model.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     product.Price,
		})
	}

	order := &model.Order{
		UserID: userID,
		Total:  total,
		Status: "pending",
		Items:  orderItems,
	}

	err := s.orderRepo.Create(order)
	return order, err
}
