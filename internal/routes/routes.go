package routes

import (
	"esdc-backend/internal/handler"
	"esdc-backend/internal/initializer"
	"esdc-backend/internal/middleware"
	"esdc-backend/internal/repository"
	"esdc-backend/internal/service"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) *gin.Engine {
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			if strings.HasPrefix(origin, "http://localhost:") {
				return true
			}
			return origin == "http://192.168.29.49:3000" || origin == "https://esdc.vercel.app"
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

	// Chat Routes (before JWT middleware)
	RegisterChatRoutes(r)

	registerUserRoutes(r, userHandler)

	// Projects Routes
	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)
	projectHandler := handler.NewProjectHandler(projectService)

	registerProjectsRoutes(r, projectHandler)

	// Product Routes (public, no JWT required)
	productRepository := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)
	registerProductRoutes(r, productHandler)

	// File Upload Routes (can be protected with middleware if needed)
	fileService := service.NewFileService("./uploads")
	fileHandler := handler.NewFileHandler(fileService)
	registerFileRoutes(r, fileHandler)

	// Serve static files (uploaded files)
	r.Static("/uploads", "./uploads")

	// Now use middleware to protect the routes
	r.Use(middleware.JwtMiddleware())
	postsRepository := repository.NewPostsRepository(db)
	postsService := service.NewPostsService(postsRepository)
	postsHandler := handler.NewPostsHandler(postsService)

	registerPostRoutes(r, postsHandler)

	// Shopping Routes (protected with JWT middleware)
	cartRepository := repository.NewCartRepository(db)
	cartService := service.NewCartService(cartRepository, productRepository)
	cartHandler := handler.NewCartHandler(cartService)
	registerCartRoutes(r, cartHandler)

	wishlistRepository := repository.NewWishlistRepository(db)
	wishlistService := service.NewWishlistService(wishlistRepository)
	wishlistHandler := handler.NewWishlistHandler(wishlistService)
	registerWishlistRoutes(r, wishlistHandler)

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository, productRepository)
	orderHandler := handler.NewOrderHandler(orderService)
	registerOrderRoutes(r, orderHandler)

	// Admin Routes
	adminService := service.NewAdminService(userRepository, projectRepository)
	adminHandler := handler.NewAdminHandler(adminService)
	registerAdminRoutes(r, adminHandler)

	return r
}
