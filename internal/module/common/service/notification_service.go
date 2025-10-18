package service

import (
	"errors"
	"esdc-backend/internal/module/common/dto"
	"esdc-backend/internal/module/common/model"
	"esdc-backend/internal/module/common/repository"
	user_repo "esdc-backend/internal/module/user/repository"
)

type NotificationService interface {
	SendNotification(data dto.NotificationRequest) error
	MarkAsRead(username string, notificationID uint) error
	GetUserNotifications(username string) ([]model.Notification, error)
	GetSingleNotification(notificationID uint) (*model.Notification, error)
	// Add more methods as needed
	// Define methods for notification service here
}

type notificationService struct {
	notificationRepo repository.NotificationRepository
	userRepo         user_repo.UserRepository
}

func NewNotificationService(
	notificationRepo repository.NotificationRepository,
	userRepo user_repo.UserRepository,
) NotificationService {
	return &notificationService{
		notificationRepo: notificationRepo,
		userRepo:         userRepo,
	}
}

func (s *notificationService) SendNotification(data dto.NotificationRequest) error {

	// Get User ID
	userID, err := s.userRepo.FindUserIDByUsername(data.Username)
	if err != nil {
		return err
	}

	s.notificationRepo.SendNotification(userID, data.Title, data.Message)

	return nil
}
func (s *notificationService) MarkAsRead(username string, notificationID uint) error {
	// check if the notification belongs to the user
	userID, err := s.userRepo.FindUserIDByUsername(username)
	if err != nil {
		return err
	}
	notification, err := s.notificationRepo.GetSingleNotification(notificationID)
	if err != nil {
		return err
	}
	if notification.UserID != userID {
		return errors.New("notification does not belong to user")
	}
	return s.notificationRepo.MarkAsRead(notificationID)
}
func (s *notificationService) GetUserNotifications(username string) ([]model.Notification, error) {
	userID, err := s.userRepo.FindUserIDByUsername(username)
	if err != nil {
		return nil, err
	}
	notifications, err := s.notificationRepo.GetUserNotifications(userID)
	if err != nil {
		return nil, err
	}
	return *notifications, nil
}
func (s *notificationService) GetSingleNotification(notificationID uint) (*model.Notification, error) {
	notification, err := s.notificationRepo.GetSingleNotification(notificationID)
	if err != nil {
		return nil, err
	}
	return &notification, nil
}
