package internal

import (
	"esdc-backend/internal/factory"
	"esdc-backend/internal/initializer"
	"esdc-backend/internal/middleware"
	"esdc-backend/internal/module/auth/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	// CORS must be applied FIRST, before any routes
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://192.168.29.49:3000", "https://esdc.vercel.app", "http://localhost:9090", "http://localhost:8080", "http://localhost:5174"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "sentry-trace", "baggage"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
	}))

	db := &initializer.DB

	handlerFactory := factory.NewHandlerFactory(db)
	authHandler := handlerFactory.GetAuthHandler()

	routes.RegisterAuthRoutes(r, authHandler)

	// Middleware for protected routes
	r.Use(middleware.JwtMiddleware())

	// Serve static files (uploaded files)
	r.Static("/uploads", "./uploads")

	return r
}
