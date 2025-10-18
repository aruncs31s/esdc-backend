package service

import (
	"errors"
	"esdc-backend/internal/module/shop/model"
	"esdc-backend/internal/module/shop/repository"

	"gorm.io/gorm"
)

type CartService interface {
	CartServiceReader
	CartServiceWriter
}
type cartService struct {
	reader CartServiceReader
	writer CartServiceWriter
}

func newCartServiceReader(cartRepo repository.CartRepository) CartServiceReader {
	return &cartServiceReader{
		cartRepo: cartRepo,
	}
}
func newCartServiceWriter(cartRepo repository.CartRepository, productRepo repository.ProductRepository) CartServiceWriter {
	return &cartServiceWriter{
		cartRepo:    cartRepo,
		productRepo: productRepo,
	}
}

func NewCartService(cartRepo repository.CartRepository, productRepo repository.ProductRepository) CartService {
	reader := newCartServiceReader(cartRepo)
	writer := newCartServiceWriter(cartRepo, productRepo)
	return &cartService{
		reader: reader,
		writer: writer,
	}
}

type CartServiceReader interface {
	GetByUserID(userID uint) ([]model.CartItem, float64, error)
}
type CartServiceWriter interface {
	Add(userID, productID uint, quantity int) (*model.Cart, error)
	Update(id uint, quantity int) error
	Delete(id uint) error
	Clear(userID uint) error
}
type cartServiceReader struct {
	cartRepo repository.CartRepository
}

type cartServiceWriter struct {
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func (s *cartService) GetByUserID(userID uint) ([]model.CartItem, float64, error) {
	return s.reader.GetByUserID(userID)
}
func (s *cartService) Add(userID, productID uint, quantity int) (*model.Cart, error) {
	return s.writer.Add(userID, productID, quantity)
}
func (s *cartService) Update(id uint, quantity int) error {
	return s.writer.Update(id, quantity)
}
func (s *cartService) Delete(id uint) error {
	return s.writer.Delete(id)
}
func (s *cartService) Clear(userID uint) error {
	return s.writer.Clear(userID)
}

func (s *cartServiceReader) GetByUserID(userID uint) ([]model.CartItem, float64, error) {
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

func (s *cartServiceWriter) Add(userID, productID uint, quantity int) (*model.Cart, error) {
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

func (s *cartServiceWriter) Update(id uint, quantity int) error {
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

func (s *cartServiceWriter) Delete(id uint) error {
	return s.cartRepo.Delete(id)
}

func (s *cartServiceWriter) Clear(userID uint) error {
	return s.cartRepo.Clear(userID)
}
