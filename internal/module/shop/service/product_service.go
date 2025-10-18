package service

import (
	"esdc-backend/internal/module/shop/model"
	repository "esdc-backend/internal/module/shop/repository"
)

type ProductService interface {
	GetAll(category, search string, limit, offset int) ([]model.Product, int64, error)
	GetByID(id uint) (*model.Product, error)
}
type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAll(category, search string, limit, offset int) ([]model.Product, int64, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetAll(category, search, limit, offset)
}

func (s *productService) GetByID(id uint) (*model.Product, error) {
	return s.repo.GetByID(id)
}
