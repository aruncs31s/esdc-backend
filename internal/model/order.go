package model

import (
	"time"
)

type Order struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	UserID    uint        `json:"user_id" gorm:"not null"`
	Total     float64     `json:"total" gorm:"type:decimal(10,2);not null"`
	Status    string      `json:"status" gorm:"default:'pending'"`
	CreatedAt time.Time   `json:"created_at"`
	Items     []OrderItem `json:"items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	OrderID   uint      `json:"order_id" gorm:"not null"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Name      string    `json:"name" gorm:"-"`
	Quantity  int       `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"type:decimal(10,2);not null"`
	CreatedAt time.Time `json:"created_at"`
	Product   Product   `json:"-" gorm:"foreignKey:ProductID"`
}
