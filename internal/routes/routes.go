package routes

import (
	"esdc-backend/internal/handler"
	"esdc-backend/internal/initializer"
	"esdc-backend/internal/middleware"
	"esdc-backend/internal/repository"
	"esdc-backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:4200" ||
				origin == "http://192.168.29.49:3000" || origin == "http://localhost:5173"
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "sentry-trace", "baggage"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))

	jwtService := service.NewJWTService()
	db := &initializer.DB
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository, jwtService)
	userHandler := handler.NewUserHandler(userService)

	registerUserRoutes(r, userHandler)

	// Now use middleware to protect the routes
	r.Use(middleware.JwtMiddleware())
	postsRepository := repository.NewPostsRepository(db)
	postsService := service.NewPostsService(postsRepository)
	postsHandler := handler.NewPostsHandler(postsService)

	registerPostRoutes(r, postsHandler)

	return r

}
