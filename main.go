package main

import (
	"esdc-backend/internal/initializer"
	"esdc-backend/internal/routes"

	_ "esdc-backend/docs" // This will be generated

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title ESDC Backend API
// @version 1.0
// @description This is the ESDC Backend API server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9090
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func init() {
	initializer.InitDB()
	initializer.InitDotenv()
}

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r = routes.RegisterRoutes(r)

	// Swagger endpoint

	r.Run(":9090")
}
