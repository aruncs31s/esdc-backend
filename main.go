package main

import (
	"esdc-backend/internal"
	"esdc-backend/internal/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.InitDotenv()
	initializer.InitDB()
}

func main() {
	r := gin.Default()

	// Disable redirects for trailing slashes to prevent CORS issues
	r.RedirectTrailingSlash = false
	r.RedirectFixedPath = false

	r = internal.RegisterRoutes(r)

	r.Run() // listen and serve on 8080
}
