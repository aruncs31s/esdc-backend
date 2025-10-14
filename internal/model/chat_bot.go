package model

import "time"

type ChatBotMessage struct {
	ID        int       `gorm:"primaryKey"`
	AskedBy   *uint     `gorm:"column:asked_by"`
	Role      string    `gorm:"column:role"`
	Category  *string   `gorm:"column:category"` // project product etc
	Content   string    `gorm:"column:content"`  // the message.
	Response  *string   `gorm:"column:response"`
	CreatedAt time.Time `gorm:"column:created_at"`
}

func (ChatBotMessage) TableName() string {
	return "chat_bot_messages"
}
