package repository

import (
	"esdc-backend/internal/module/common/model"

	"gorm.io/gorm"
)

type ChatBotRepository interface {
	SaveMessage(message *model.ChatBotMessage) error
	// GetMessagesByChatBotID(chatBotID int) ([]model.Message, error)
	GetMessageByUserID(userID int) ([]model.ChatBotMessage, error)
	GetMessageByCategory(category string) ([]model.ChatBotMessage, error)
}

type chatBotRepository struct {
	db *gorm.DB
}

func NewChatBotRepository(db *gorm.DB) ChatBotRepository {
	return &chatBotRepository{db: db}
}
func (r *chatBotRepository) SaveMessage(message *model.ChatBotMessage) error {
	if err := r.db.Create(message).Error; err != nil {
		return err
	}
	return nil
}

//	func (r *chatBotRepository) GetMessagesByChatBotID(chatBotID int) ([]model.Message, error) {
//		var messages []model.Message
//		if err := r.db.Where("chat_bot_id = ?", chatBotID).Find(&messages).Error; err != nil {
//			return nil, err
//		}
//		return messages, nil
//	}
func (r *chatBotRepository) GetMessageByUserID(userID int) ([]model.ChatBotMessage, error) {
	var messages []model.ChatBotMessage
	if err := r.db.Where("asked_by = ?", userID).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
func (r *chatBotRepository) GetMessageByCategory(category string) ([]model.ChatBotMessage, error) {
	var messages []model.ChatBotMessage
	if err := r.db.Where("category = ?", category).Find(&messages).Error; err != nil {
		return nil, err
	}
	return messages, nil
}
