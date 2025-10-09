package repository

import (
	"esdc-backend/internal/model"

	"gorm.io/gorm"
)

type WishlistRepository struct {
	db *gorm.DB
}

func NewWishlistRepository(db *gorm.DB) *WishlistRepository {
	return &WishlistRepository{db: db}
}

func (r *WishlistRepository) GetByUserID(userID uint) ([]model.Wishlist, error) {
	var wishlists []model.Wishlist
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&wishlists).Error
	return wishlists, err
}

func (r *WishlistRepository) FindByUserAndProduct(userID, productID uint) (*model.Wishlist, error) {
	var wishlist model.Wishlist
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&wishlist).Error
	return &wishlist, err
}

func (r *WishlistRepository) Add(wishlist *model.Wishlist) error {
	return r.db.Create(wishlist).Error
}

func (r *WishlistRepository) Delete(id uint) error {
	return r.db.Delete(&model.Wishlist{}, id).Error
}
