package model

type Project struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Image       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Link        string `gorm:"not null"`
}

func (Project) TableName() string {
	return "projects"
}
