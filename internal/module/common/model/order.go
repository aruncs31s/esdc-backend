package model

import "time"

// Order represents a customer order
type Order struct {
	ID        uint        `gorm:"primaryKey" json:"id"`
	UserID    uint        `json:"user_id"`
	Total     float64     `json:"total"`
	Status    string      `json:"status"` // e.g., pending, completed, cancelled
	Items     []OrderItem `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	OrderID   uint    `json:"order_id"`
	ProductID uint    `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
