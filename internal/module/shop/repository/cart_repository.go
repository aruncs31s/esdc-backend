package repository

import (
	"esdc-backend/internal/module/shop/model"

	"gorm.io/gorm"
)

type CartRepository interface {
	CartRepositoryReader
	CartRepositoryWriter
}
type cartRepository struct {
	reader CartRepositoryReader
	writer CartRepositoryWriter
}

type CartRepositoryReader interface {
	GetByUserID(userID uint) ([]model.Cart, error)
	GetByID(id uint) (*model.Cart, error)
	FindByUserAndProduct(userID, productID uint) (*model.Cart, error)
}
type CartRepositoryWriter interface {
	Add(cart *model.Cart) error
	Update(cart *model.Cart) error
	Delete(id uint) error
	Clear(userID uint) error
}
type cartRepositoryReader struct {
	db *gorm.DB
}
type cartRepositoryWriter struct {
	db *gorm.DB
}

func newCartRepositoryReader(db *gorm.DB) CartRepositoryReader {
	return &cartRepositoryReader{
		db: db,
	}
}
func newCartRepositoryWriter(db *gorm.DB) CartRepositoryWriter {
	return &cartRepositoryWriter{
		db: db,
	}
}
func NewCartRepository(db *gorm.DB) *cartRepository {
	reader := newCartRepositoryReader(db)
	writer := newCartRepositoryWriter(db)
	return &cartRepository{
		reader: reader,
		writer: writer,
	}
}
func (r *cartRepository) GetByUserID(userID uint) ([]model.Cart, error) {
	return r.reader.GetByUserID(userID)
}
func (r *cartRepository) GetByID(id uint) (*model.Cart, error) {
	return r.reader.GetByID(id)
}
func (r *cartRepository) FindByUserAndProduct(userID, productID uint) (*model.Cart, error) {
	return r.reader.FindByUserAndProduct(userID, productID)
}
func (r *cartRepository) Add(cart *model.Cart) error {
	return r.writer.Add(cart)
}
func (r *cartRepository) Update(cart *model.Cart) error {
	return r.writer.Update(cart)
}
func (r *cartRepository) Delete(id uint) error {
	return r.writer.Delete(id)
}
func (r *cartRepository) Clear(userID uint) error {
	return r.writer.Clear(userID)
}
func (r *cartRepositoryReader) GetByUserID(userID uint) ([]model.Cart, error) {
	var carts []model.Cart
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *cartRepositoryReader) GetByID(id uint) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Preload("Product").First(&cart, id).Error
	return &cart, err
}

func (r *cartRepositoryReader) FindByUserAndProduct(userID, productID uint) (*model.Cart, error) {
	var cart model.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error
	return &cart, err
}

func (r *cartRepositoryWriter) Add(cart *model.Cart) error {
	return r.db.Create(cart).Error
}

func (r *cartRepositoryWriter) Update(cart *model.Cart) error {
	return r.db.Save(cart).Error
}

func (r *cartRepositoryWriter) Delete(id uint) error {
	return r.db.Delete(&model.Cart{}, id).Error
}

func (r *cartRepositoryWriter) Clear(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&model.Cart{}).Error
}
