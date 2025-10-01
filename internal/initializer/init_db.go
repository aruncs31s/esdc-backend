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
	err = DB.AutoMigrate(
		model.Post{},
		model.User{},
		model.Project{},
	)
	if err != nil {
		panic("failed to migrate database")
	}
}
