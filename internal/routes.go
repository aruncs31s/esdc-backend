package internal

import (
	"esdc-backend/internal/initializer"

	auth "github.com/aruncs31s/esdcauthmodule"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	// CORS must be applied FIRST, before any routes
	UseCors(r)
	db := &initializer.DB
	auth.InitAuthModule(r, db)
	auth.AddJWTMiddleware(r)
	// Middleware for protected routes

	// Serve static files (uploaded files)
	// r.Static("/uploads", "./uploads")
	return r
}
