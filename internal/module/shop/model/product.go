package model

import (
	"time"
)

// Product represents a product in the system
// @Description Product model
type Product struct {
	ID          uint      `json:"id" gorm:"primaryKey" example:"1"`                           // Product ID
	Name        string    `json:"name" gorm:"not null" example:"Sample Product"`              // Product name
	Description string    `json:"description" example:"This is a sample product description"` // Product description
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null" example:"29.99"`   // Product price
	Image       string    `json:"image" example:"https://example.com/product.jpg"`            // Product image URL
	Category    string    `json:"category" example:"electronics"`                             // Product category
	Stock       int       `json:"stock" gorm:"default:0" example:"100"`                       // Available stock
	CreatedAt   time.Time `json:"created_at" example:"2023-01-01T00:00:00Z"`                  // Creation timestamp
	UpdatedAt   time.Time `json:"updated_at" example:"2023-01-01T00:00:00Z"`                  // Last update timestamp
}

func (Product) TableName() string {
	return "products"
}
