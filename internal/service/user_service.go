package service

import (
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
)

type UserService interface {
	Login(email, password string) (string, error)
	Register(username, email, password string) error
	VerifyEmail(token string) error
	ForgotPassword(email string) error
	ResetPassword(token, newPassword string) error
}
type userService struct {
	userRepo   repository.UserRepository
	jwtService JWTService
}

func NewUserService(userRepo repository.UserRepository, jwtService JWTService) UserService {
	return &userService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (s *userService) Login(email, password string) (string, error) {

	// Check if the user exists
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}
	if user.Password != password {
		return "", err
	}
	// Generate JWT token
	token, err := s.jwtService.CreateToken(user.Username, user.Email, user.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *userService) Register(username, email, password string) error {

	user := model.User{
		Username: username,
		Email:    email,
		Password: password,
		Role:     getRole(username),
	}
	err := s.userRepo.CreateUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (s *userService) VerifyEmail(token string) error {
	// Implementation of email verification logic
	return nil
}

func (s *userService) ForgotPassword(email string) error {
	// Implementation of forgot password logic
	return nil
}

func (s *userService) ResetPassword(token, newPassword string) error {
	// Implementation of reset password logic
	return nil
}
