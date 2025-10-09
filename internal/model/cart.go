package model

import (
	"time"
)

type Cart struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Quantity  int       `json:"quantity" gorm:"default:1"`
	CreatedAt time.Time `json:"created_at"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID"`
}

type CartItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Image     string  `json:"image"`
	Quantity  int     `json:"quantity"`
	Subtotal  float64 `json:"subtotal"`
}
