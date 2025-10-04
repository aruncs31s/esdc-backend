package model

type UserDetails struct {
	ID    uint   `gorm:"primaryKey"`
	Email string `gorm:"unique;not null"`
	Phone string `gorm:"not null"`
}

func (UserDetails) TableName() string {
	return "user_details"
}
