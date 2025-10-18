package model

type Notification struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"column:user_id;not null" json:"user_id"`
	Title     string `gorm:"column:title;not null" json:"title"`
	Message   string `gorm:"column:message;not null" json:"message"`
	Achieved  bool   `gorm:"column:achieved;default:false" json:"achieved"`
	Read      bool   `gorm:"column:read;default:false" json:"read"`
	ReadAt    *int64 `gorm:"column:read_at" json:"read_at"`
	CreatedAt int64  `gorm:"column:created_at;autoCreateTime" json:"created_at"`
}

func (Notification) TableName() string {
	return "notifications"
}
