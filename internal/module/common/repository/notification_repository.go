package repository

import (
	"esdc-backend/internal/module/common/model"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	SendNotification(userID uint, title string, message string) error
	MarkAsRead(notificationID uint) error
	GetUserNotifications(userID uint) (*[]model.Notification, error)
	GetSingleNotification(notificationID uint) (model.Notification, error)
	DeleteNotification(notificationID uint) error
	// Add more methods as needed
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}
func (r *notificationRepository) SendNotification(userID uint, title string, message string) error {
	notification := model.Notification{
		UserID:   userID,
		Title:    title,
		Message:  message,
		Read:     false,
		Achieved: false,
	}
	if err := r.db.Create(&notification).Error; err != nil {
		return err
	}
	return nil
}
func (r *notificationRepository) MarkAsRead(notificationID uint) error {
	err := r.db.
		Model(&model.Notification{}).
		Where("id = ?", notificationID).
		Updates(map[string]interface{}{
			"read":    true,
			"read_at": gorm.Expr("strftime('%s','now')"),
		}).Error
	return err
}
func (r *notificationRepository) GetUserNotifications(userID uint) (*[]model.Notification, error) {
	var notifications []model.Notification
	err := r.db.
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&notifications).Error
	return &notifications, err
}
func (r *notificationRepository) GetSingleNotification(notificationID uint) (model.Notification, error) {
	var notification model.Notification
	err := r.db.
		Where("id = ?", notificationID).
		First(&notification).Error
	return notification, err
}
func (r *notificationRepository) DeleteNotification(notificationID uint) error {
	err := r.db.
		Where("id = ?", notificationID).
		Delete(&model.Notification{}).Error
	return err
}
