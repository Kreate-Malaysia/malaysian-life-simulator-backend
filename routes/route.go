package routes

import (
	"gin/config"
	"gin/controller"
	"gin/database"
	"gin/services"
	"log"

	"github.com/gin-gonic/gin"
)

// SetupRoutes sets up all the routes for the application
func SetupRoutes(router *gin.Engine) {
	// Load environment config
	cfg, err := config.LoadConfig()
	if (err != nil) {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize database connection
	database.InitDB(cfg)

	// Initialize services
	googleService := services.NewGoogleOAuthService(database.DB)
	userService := services.NewUserService(database.DB, googleService)

	// Initialize controllers
	userController := controller.NewUserController(userService, googleService)

    // Register the route for login
    router.POST("/api/login", func(ctx *gin.Context) {
        userController.HandleLogin(ctx.Writer, ctx.Request)
    })

    // Register the route for signup
    router.POST("/api/signup", func(ctx *gin.Context) {
        userController.HandleSignup(ctx.Writer, ctx.Request)
    })

    // Register the route for login using Google OAuth
    router.POST("/api/google/login", func(ctx *gin.Context) {
        userController.HandleLoginWithGoogleAccessToken(ctx.Writer, ctx.Request)
    })

    // Register the route for signup using Google OAuth
    router.POST("/api/google/signup", func(ctx *gin.Context) {
        userController.HandleSignupWithGoogleAccessToken(ctx.Writer, ctx.Request)
    })
}