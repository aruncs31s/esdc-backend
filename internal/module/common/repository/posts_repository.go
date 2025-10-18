package repository

import (
	"esdc-backend/internal/module/common/model"

	"gorm.io/gorm"
)

type PostsRepository interface {
	FindAll() (*[]model.Post, error)
}

type postsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) PostsRepository {
	return &postsRepository{db: db}
}

func (r *postsRepository) FindAll() (*[]model.Post, error) {
	var posts []model.Post
	err := r.db.Find(&posts).Error
	if err != nil {
		return nil, err
	}
	return &posts, nil
}
