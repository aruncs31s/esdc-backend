package model

import (
	"time"
)

type Wishlist struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID"`
}

func (Wishlist) TableName() string {
	return "wishlists"
}

type WishlistItem struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Image     string  `json:"image"`
	Category  string  `json:"category"`
	Stock     int     `json:"stock"`
}

func (WishlistItem) TableName() string {
	return "wishlists"
}
