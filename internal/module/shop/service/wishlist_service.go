package service

import (
	"errors"
	"esdc-backend/internal/module/shop/model"
	"esdc-backend/internal/module/shop/repository"

	"gorm.io/gorm"
)

type WishlistService interface {
	GetByUserID(userID uint) ([]model.WishlistItem, error)
	Add(userID, productID uint) (*model.Wishlist, error)
	Delete(id uint) error
}
type wishlistService struct {
	wishlistRepo repository.WishlistRepository
}

func NewWishlistService(wishlistRepo repository.WishlistRepository) WishlistService {
	return &wishlistService{
		wishlistRepo: wishlistRepo,
	}
}

func (s *wishlistService) GetByUserID(userID uint) ([]model.WishlistItem, error) {
	wishlists, err := s.wishlistRepo.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var items []model.WishlistItem
	for _, wishlist := range wishlists {
		item := model.WishlistItem{
			ID:        wishlist.ID,
			ProductID: wishlist.ProductID,
			Name:      wishlist.Product.Name,
			Price:     wishlist.Product.Price,
			Image:     wishlist.Product.Image,
			Category:  wishlist.Product.Category,
			Stock:     wishlist.Product.Stock,
		}
		items = append(items, item)
	}

	return items, nil
}

func (s *wishlistService) Add(userID, productID uint) (*model.Wishlist, error) {
	// Check if already exists
	existing, err := s.wishlistRepo.FindByUserAndProduct(userID, productID)
	if err == nil && existing.ID > 0 {
		return nil, errors.New("product already in wishlist")
	}

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	wishlist := &model.Wishlist{
		UserID:    userID,
		ProductID: productID,
	}

	err = s.wishlistRepo.Add(wishlist)
	return wishlist, err
}

func (s *wishlistService) Delete(id uint) error {
	return s.wishlistRepo.Delete(id)
}
