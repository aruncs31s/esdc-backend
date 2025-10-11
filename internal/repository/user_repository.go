package repository

import (
	"esdc-backend/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (model.User, error)
	FindUsersByUsernames(usernames []string) ([]model.User, error)
	FindByID(id int) (model.User, error)
	FindUserIDByUsername(username string) (int, error)
	FindByEmail(email string) (model.User, error)
	CreateUser(user *model.User) error
	UpdateUser(user model.User) error
	GetAllUsers() ([]model.User, error)
	GetUsersCount() (int, error)
	DeleteUserByID(userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
func (r *userRepository) FindByID(id int) (model.User, error) {
	var user model.User
	result := r.db.First(&user, id)
	return user, result.Error
}

func (r *userRepository) FindByUsername(username string) (model.User, error) {
	var user model.User
	result := r.db.Where("username = ?", username).First(&user)
	return user, result.Error
}
func (r *userRepository) FindUsersByUsernames(usernames []string) ([]model.User, error) {
	var users []model.User
	result := r.db.Where("username IN ?", usernames).Find(&users)
	return users, result.Error
}

func (r *userRepository) FindUserIDByUsername(username string) (int, error) {
	var user model.User
	result := r.db.Select("id").Where("username = ?", username).First(&user)
	return int(user.ID), result.Error
}
func (r *userRepository) FindByEmail(email string) (model.User, error) {
	var user model.User
	result := r.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (r *userRepository) CreateUser(user *model.User) error {
	result := r.db.Create(&user)
	return result.Error
}

func (r *userRepository) UpdateUser(user model.User) error {
	result := r.db.Save(&user)
	return result.Error
}
func (r *userRepository) GetAllUsers() ([]model.User, error) {
	var users []model.User
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *userRepository) GetUsersCount() (int, error) {
	var count int64
	result := r.db.Model(&model.User{}).Count(&count)
	return int(count), result.Error
}
func (r *userRepository) DeleteUserByID(userID uint) error {
	result := r.db.Delete(&model.User{}, userID)
	return result.Error
}
