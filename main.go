package main

import (
	"esdc-backend/internal"
	"esdc-backend/internal/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.InitDB()
	initializer.InitDotenv()
}

func main() {
	r := gin.Default()

	r = internal.RegisterRoutes(r)

	r.Run() // listen and serve on 8080
}
