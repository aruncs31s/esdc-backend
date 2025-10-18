package repository

import (
	"esdc-backend/internal/module/shop/model"

	"gorm.io/gorm"
)

type WishlistReader interface {
	GetByUserID(userID uint) ([]model.Wishlist, error)
	FindByUserAndProduct(userID, productID uint) (*model.Wishlist, error)
}

type WishlistWriter interface {
	Add(wishlist *model.Wishlist) error
	Delete(id uint) error
}

type WishlistRepository interface {
	WishlistReader
	WishlistWriter
}
type wishlistRepository struct {
	reader WishlistReader
	writer WishlistWriter
}

type wishlistReader struct {
	db *gorm.DB
}

type wishlistWriter struct {
	db *gorm.DB
}

func newWishlistReader(db *gorm.DB) WishlistReader {
	return &wishlistReader{db: db}
}

func newWishlistWriter(db *gorm.DB) WishlistWriter {
	return &wishlistWriter{db: db}
}

func NewWishlistRepository(db *gorm.DB) WishlistRepository {
	reader := newWishlistReader(db)
	writer := newWishlistWriter(db)
	return &wishlistRepository{
		reader: reader,
		writer: writer,
	}
}

func (r *wishlistRepository) GetByUserID(userID uint) ([]model.Wishlist, error) {
	return r.reader.GetByUserID(userID)
}
func (r *wishlistRepository) FindByUserAndProduct(userID, productID uint) (*model.Wishlist, error) {
	return r.reader.FindByUserAndProduct(userID, productID)
}
func (r *wishlistRepository) Add(wishlist *model.Wishlist) error {
	return r.writer.Add(wishlist)
}
func (r *wishlistRepository) Delete(id uint) error {
	return r.writer.Delete(id)
}

func (r *wishlistReader) GetByUserID(userID uint) ([]model.Wishlist, error) {
	var wishlists []model.Wishlist
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&wishlists).Error
	return wishlists, err
}

func (r *wishlistReader) FindByUserAndProduct(userID, productID uint) (*model.Wishlist, error) {
	var wishlist model.Wishlist
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&wishlist).Error
	return &wishlist, err
}

func (r *wishlistWriter) Add(wishlist *model.Wishlist) error {
	return r.db.Create(wishlist).Error
}

func (r *wishlistWriter) Delete(id uint) error {
	return r.db.Delete(&model.Wishlist{}, id).Error
}
