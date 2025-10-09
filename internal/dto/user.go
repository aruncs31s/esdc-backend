package dto

// LoginRequest represents the login request payload
// @Description Login request payload
type LoginRequest struct {
	// Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email" example:"user@example.com"` // User email address
	Password string `json:"password" binding:"required" example:"password123"`         // User password
}

// AdminRegisterRequest represents admin user registration request
// @Description Admin user registration request payload
type AdminRegisterRequest struct {
	RegisterRequest
	Role string `json:"role" example:"admin"` // User role (e.g., "admin", "user")
}
// RegisterRequest represents user registration request
// @Description User registration request payload
type RegisterRequest struct {
	Username       string `json:"username" example:"aruncs31s"`                                    // Username (optional)
	Name           string `json:"name" binding:"required" example:"Arun CS"`                        // Full name
	Email          string `json:"email" binding:"required,email" example:"aruncs31ss@gmail.com"`    // Email address
	GithubUsername string `json:"github_username" example:"aruncs31s"`                             // GitHub username (optional)
	Password       string `json:"password" binding:"required,min=6" example:"password123"`          // Password (minimum 6 characters)
}
