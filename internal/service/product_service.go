package service

import (
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(category, search string, limit, offset int) ([]model.Product, int64, error) {
	if limit <= 0 {
		limit = 50
	}
	if offset < 0 {
		offset = 0
	}
	return s.repo.GetAll(category, search, limit, offset)
}

func (s *ProductService) GetByID(id uint) (*model.Product, error) {
	return s.repo.GetByID(id)
}
