package dto

type NotificationRequest struct {
	Username string `json:"username" binding:"required"` // target username
	Title    string `json:"title" binding:"required"`
	Message  string `json:"message" binding:"required"`
}

type NotificationResponse struct {
	ID        uint   `json:"id"`
	UserID    uint   `json:"user_id"`
	Message   string `json:"message"`
	Title     string `json:"title"`
	Achieved  *bool  `json:"achieved"`
	Read      bool   `json:"read"`
	ReadAt    *int64 `json:"read_at"`
	CreatedAt int64  `json:"created_at"`
}
