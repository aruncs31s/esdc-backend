package service

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// Global secret key - initialized once
var SecretKey []byte

func init() {
	// Initialize SecretKey from environment or use default
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "something_big"
		log.Println("⚠️  WARNING: Using default JWT secret. Set JWT_SECRET environment variable for production!")
	}
	SecretKey = []byte(secretKey)
	log.Printf("✅ JWT Secret Key initialized (length: %d bytes)\n", len(SecretKey))
}

type JWTService interface {
	CreateToken(username, email, role, name string) (string, error)
}

func NewJWTService() JWTService {
	return &jwtService{}
}

type jwtService struct{}

func (s *jwtService) CreateToken(username, email, role, name string) (string, error) {
	log.Printf("🔐 Creating JWT token for user: %s, role: %s\n", username, role)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      email,
		"username": username,
		"role":     role,
		"iss":      "esdc-backend",
		"exp":      time.Now().Add(60 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})

	// Use the global SecretKey (same one used in middleware)
	tokenString, err := claims.SignedString(SecretKey)
	if err != nil {
		log.Printf("❌ Error creating token: %v\n", err)
		return "", err
	}

	log.Printf("✅ Token created successfully for %s\n", username)
	return tokenString, nil
}

func GetRole(email string) string {
	if email == "aruncs31ss@gmail.com" || email == "aruncs31s@gmail.com" {
		return "admin"
	}
	return "user"
}
