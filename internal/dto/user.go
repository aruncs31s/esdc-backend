package dto

type LoginRequest struct {
	// Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}
type RegisterRequest struct {
	Username       string `json:"username"`
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	GithubUsername string `json:"github_username"`
	Role           string `json:"role" binding:"required,oneof=admin user"`
	Password       string `json:"password" binding:"required,min=6"`
	Link           string `json:"link"`
}
