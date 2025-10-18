package dto

type AddToWishlistRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
}
