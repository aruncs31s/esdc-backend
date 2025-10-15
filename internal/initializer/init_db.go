package initializer

import (
	"esdc-backend/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("database/db.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = *db

	if err := db.AutoMigrate(&model.User{}, &model.Project{}, &model.Tag{}, &model.Github{}, &model.Submission{}, &model.UserDetails{}, &model.Technologies{}); err != nil {
		panic("failed to migrate database , USER")
	}

	if err := DB.AutoMigrate(model.Teams{}); err != nil {
		panic("failed to migrate database , TEAMS")
	}
	// Migrate the schema
	if err := DB.AutoMigrate(model.Post{}); err != nil {
		panic("failed to migrate database , POST")
	}

	// Shopping tables
	if err := DB.AutoMigrate(model.Product{}); err != nil {
		panic("failed to migrate database , PRODUCT")
	}
	if err := DB.AutoMigrate(model.Cart{}); err != nil {
		panic("failed to migrate database , CART")
	}
	if err := DB.AutoMigrate(model.Wishlist{}); err != nil {
		panic("failed to migrate database , WISHLIST")
	}
	if err := DB.AutoMigrate(model.Order{}); err != nil {
		panic("failed to migrate database , ORDER")
	}
	if err := DB.AutoMigrate(model.OrderItem{}); err != nil {
		panic("failed to migrate database , ORDER_ITEM")
	}
	if err := DB.AutoMigrate(model.Ollama{}); err != nil {
		panic("failed to migrate database , PAYMENT")
	}
	if err := DB.AutoMigrate(model.ChatBotMessage{}); err != nil {
		panic("failed to migrate database , CHAT_BOT_MESSAGE")
	}
	if err := DB.AutoMigrate(model.Notification{}); err != nil {
		panic("failed to migrate database , NOTIFICATION")
	}
	admins := []model.User{
		{
			Name:     "Arun CS",
			Email:    "arun31s@gamil.com",
			Role:     "admin",
			Password: "12345678",
			Status:   "active",
			Username: "aruncs",
			Github: &model.Github{
				Username: "arun31s",
			},
		},
		{
			Name:     "ESDC Admin",
			Email:    "esdc@gcek.ac.in",
			Role:     "admin",
			Password: "12345678",
			Status:   "active",
			Username: "esdc",
			Github: &model.Github{
				Username: "esdc",
			},
		},
	}
	for _, admin := range admins {
		var existingUser model.User
		result := db.Where("email = ?", admin.Email).First(&existingUser)
		if result.Error == gorm.ErrRecordNotFound {
			db.Create(&admin)
		}
	}

}
