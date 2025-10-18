package initializer

import (
	"log"

	model "github.com/aruncs31s/esdcmodels"
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

	if err := db.AutoMigrate(model.User{}, model.Github{}, model.Notification{}, model.Technologies{}, model.UserDetails{}, model.Project{}, model.Teams{}, model.Tag{}); err != nil {
		log.Fatal("failed to migrate database")
	}

}
