package service

import (
	"errors"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"

	"gorm.io/gorm"
)

type CartService struct {
	cartRepo    *repository.CartRepository
	productRepo *repository.ProductRepository
}

func NewCartService(cartRepo *repository.CartRepository, productRepo *repository.ProductRepository) *CartService {
	return &CartService{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func (s *CartService) GetByUserID(userID uint) ([]model.CartItem, float64, error) {
	carts, err := s.cartRepo.GetByUserID(userID)
	if err != nil {
		return nil, 0, err
	}

	var items []model.CartItem
	var total float64

	for _, cart := range carts {
		subtotal := cart.Product.Price * float64(cart.Quantity)
		item := model.CartItem{
			ID:        cart.ID,
			ProductID: cart.ProductID,
			Name:      cart.Product.Name,
			Price:     cart.Product.Price,
			Image:     cart.Product.Image,
			Quantity:  cart.Quantity,
			Subtotal:  subtotal,
		}
		items = append(items, item)
		total += subtotal
	}

	return items, total, nil
}

func (s *CartService) Add(userID, productID uint, quantity int) (*model.Cart, error) {
	// Validate product exists and has stock
	product, err := s.productRepo.GetByID(productID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	if product.Stock < quantity {
		return nil, errors.New("insufficient stock")
	}

	// Check if item already exists in cart
	existingCart, err := s.cartRepo.FindByUserAndProduct(userID, productID)
	if err == nil {
		// Update quantity
		existingCart.Quantity += quantity
		if existingCart.Quantity > product.Stock {
			return nil, errors.New("insufficient stock")
		}
		err = s.cartRepo.Update(existingCart)
		return existingCart, err
	}

	// Create new cart item
	cart := &model.Cart{
		UserID:    userID,
		ProductID: productID,
		Quantity:  quantity,
	}

	err = s.cartRepo.Add(cart)
	return cart, err
}

func (s *CartService) Update(id uint, quantity int) error {
	cart, err := s.cartRepo.GetByID(id)
	if err != nil {
		return err
	}

	// Check stock
	if cart.Product.Stock < quantity {
		return errors.New("insufficient stock")
	}

	cart.Quantity = quantity
	return s.cartRepo.Update(cart)
}

func (s *CartService) Delete(id uint) error {
	return s.cartRepo.Delete(id)
}

func (s *CartService) Clear(userID uint) error {
	return s.cartRepo.Clear(userID)
}
