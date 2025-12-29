package internal

import (
	"esdc-backend/internal/initializer"

	"github.com/aruncs31s/azf/shared/logger"
	auth "github.com/aruncs31s/esdcauthmodule"
	project "github.com/aruncs31s/esdcprojectmodule"
	user "github.com/aruncs31s/esdcusermodule"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	// CORS must be applied FIRST, before any routes
	UseCors(r)
	db := initializer.DB
	// auth.InitAuthModule(r, db)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})
	// Initialize Project Module
	project.InitProjectModule(r, db)
	// Register Public Project Routes
	project.RegisterPublicProjectRoutes()
	logger := logger.GetLogger()
	// Register User Routes
	user.InitUserModule(db, logger)

	user.RegisterPublicUserRoutes(r)

	auth.AddJWTMiddleware(r)

	project.RegisterPrivateProjectRoutes(r)
	// Middleware for protected routes

	// Serve static files (uploaded files)
	// r.Static("/uploads", "./uploads")
	return r
}
