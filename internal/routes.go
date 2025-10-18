package internal

import (
	"esdc-backend/internal/initializer"

	auth "github.com/aruncs31s/esdcauthmodule"
	project "github.com/aruncs31s/esdcprojectmodule"
	user "github.com/aruncs31s/esdcusermodule"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	// CORS must be applied FIRST, before any routes
	UseCors(r)
	db := &initializer.DB
	auth.InitAuthModule(r, db)
	// Initialize Project Module
	project.InitProjectModule(r, db)
	// Register Public Project Routes
	project.RegisterPublicProjectRoutes()

	// Register User Routes
	user.InitUserModule(db)
	user.RegisterPublicUserRoutes(r)

	auth.AddJWTMiddleware(r)

	project.RegisterPrivateProjectRoutes(r)
	// Middleware for protected routes

	// Serve static files (uploaded files)
	// r.Static("/uploads", "./uploads")
	return r
}
