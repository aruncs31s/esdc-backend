package initializer

import (
	"log"

	"github.com/joho/godotenv"
)

func InitDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
