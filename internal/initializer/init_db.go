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

	// Migrate the schema

	if err := DB.AutoMigrate(model.Post{}); err != nil {
		panic("failed to migrate database , POST")
	}
	if err := DB.AutoMigrate(model.User{}); err != nil {
		panic("failed to migrate database , USER")
	}
	if err := DB.AutoMigrate(model.Project{}); err != nil {
		panic("failed to migrate database , PROJECT")
	}
	if err := DB.AutoMigrate(model.Github{}); err != nil {
		panic("failed to migrate database , GITHUB")
	}
	if err := DB.AutoMigrate(model.UserDetails{}); err != nil {
		panic("failed to migrate database , CHALLENGE")
	}

}
