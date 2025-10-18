package repository

import (
	"esdc-backend/internal/module/common/model"

	"gorm.io/gorm"
)

type OllamaRepository interface {
	GetAllChats() (*[]model.Ollama, error)
	GetChatByUserID(user_id int) (*[]model.Ollama, error)
	SaveChat(chat *model.Ollama) error
}
type ollamaRepository struct {
	db *gorm.DB
}

func NewOllamaRepository(db *gorm.DB) OllamaRepository {
	return &ollamaRepository{
		db: db,
	}
}
func (r *ollamaRepository) GetAllChats() (*[]model.Ollama, error) {
	var chats []model.Ollama
	if err := r.db.Find(&chats).Error; err != nil {
		return nil, err
	}
	return &chats, nil
}
func (r *ollamaRepository) GetChatByUserID(user_id int) (*[]model.Ollama, error) {
	var chats []model.Ollama
	if err := r.db.Where("user_id = ?", user_id).Find(&chats).Error; err != nil {
		return nil, err
	}
	return &chats, nil
}
func (r *ollamaRepository) SaveChat(chat *model.Ollama) error {
	if err := r.db.Create(chat).Error; err != nil {
		return err
	}
	return nil
}
