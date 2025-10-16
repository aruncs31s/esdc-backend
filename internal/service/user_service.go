package service

import (
	"esdc-backend/internal/dto"
	"esdc-backend/internal/model"
	"esdc-backend/internal/repository"
	"esdc-backend/utils"
)

type UserService interface {
	Login(email, password string) (string, error)
	Register(data dto.RegisterRequest) error
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
		return "", utils.ErrUserNotExists
	}

	if user.Password != password {
		return "", utils.ErrPasswordDoesNotMatch
	}
	// Generate JWT token
	token, err := s.jwtService.CreateToken(user.Username, user.Email, user.Role, user.Name)
	if err != nil {
		return "", utils.ErrGeneratingJWT
	}
	return token, nil
}

func (s *userService) Register(data dto.RegisterRequest) error {

	user := model.User{
		Name:     data.Name,
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Role:     getRole(data.Username),
	}
	github := model.Github{
		Username: data.GithubUsername,
	}
	user.Github = &github

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
