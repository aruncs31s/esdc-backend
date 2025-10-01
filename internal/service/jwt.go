package service

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

// TODO: Fix this
// var SecretKey = []byte(os.Getenv("JWT_SECRET"))
var SecretKey = []byte("something_big")

type JWTService interface {
	CreateToken(username, email, role string) (string, error)
}

func NewJWTService() JWTService {
	return &jwtService{}
}

type jwtService struct{}

func (s *jwtService) CreateToken(username, email, role string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if SecretKey == nil {
		SecretKey = []byte(secretKey)
	}

	log.Println("Secret Key: ", secretKey)
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  email,
		"user": username,
		"role": role,
		"iss":  "esdc-backend",
		"exp":  time.Now().Add(60 * time.Hour).Unix(),
		"iat":  time.Now().Unix(),
	})
	tokenString, err := claims.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func getRole(email string) string {
	if email == "aruncs31ss@gmail.com" || email == "aruncs31s@gmail.com" {
		return "admin"
	}
	return "user"
}
