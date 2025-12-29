package internal

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func UseCors(r *gin.Engine) {
	config := cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:3000",
			"http://192.168.29.49:3000",
			"https://esdc.vercel.app",
			"http://localhost:9090",
			"http://localhost:8080",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "sentry-trace", "baggage", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
	}

	r.Use(cors.New(config))
}
