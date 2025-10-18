package service

import (
	"esdc-backend/internal/module/common/model"
	"esdc-backend/internal/module/common/repository"
)

type PostsService interface {
	GetAllPosts() (*[]model.Post, error)
}
type postsService struct {
	postsRepo repository.PostsRepository
}

func NewPostsService(postsRepo repository.PostsRepository) PostsService {
	return &postsService{postsRepo: postsRepo}
}

func (s *postsService) GetAllPosts() (*[]model.Post, error) {
	posts, err := s.postsRepo.FindAll()
	if err != nil {
		return nil, err
	}
	return posts, nil

}
