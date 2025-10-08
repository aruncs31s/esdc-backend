package main

import (
	"esdc-backend/internal/initializer"
	"esdc-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.InitDB()
	initializer.InitDotenv()
}

func main() {
	r := gin.Default()
	r = routes.RegisterRoutes(r)
	r.Run(":9090")
}
