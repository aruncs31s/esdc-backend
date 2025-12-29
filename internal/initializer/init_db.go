package initializer

import (
	db "esdc-backend/external"
	"log"

	model "github.com/aruncs31s/esdcmodels"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db := db.GetDB()

	if err := db.AutoMigrate(model.User{}, model.Github{}, model.Notification{}, model.Technologies{}, model.UserDetails{}, model.Project{}, model.Teams{}, model.Tag{}); err != nil {
		log.Fatal("failed to migrate database")
	}
	DB = db
	return db
}
