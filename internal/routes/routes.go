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
	// CORS must be applied FIRST, before any routes
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "http://localhost:3000", "http://192.168.29.49:3000", "https://esdc.vercel.app", "http://localhost:9090", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept", "sentry-trace", "baggage"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * 3600, // 12 hours
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
	projectService := service.NewProjectService(projectRepository, userRepository)
	projectHandler := handler.NewProjectHandler(projectService)

	registerProjectsRoutes(r, projectHandler)

	//	ChatBot Routes
	chatBotRepository := repository.NewChatBotRepository(db)
	ollamaRepository := repository.NewOllamaRepository(db)
	chatBotService := service.NewChatBotService(chatBotRepository, ollamaRepository, userRepository)
	chatBotHandler := handler.NewChatBotHandler(chatBotService)

	registerChatbotRoutes(r, chatBotHandler)

	r.Use(middleware.JwtMiddleware())
	registerProtectedProjectsRoutes(r, projectHandler)

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
	adminHandler := handler.NewAdminHandler(adminService, projectService)
	registerAdminRoutes(r, adminHandler)

	return r
}
