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
	playerService := services.NewPlayerService(database.DB)
	feedbackService := services.NewFeedbackService(database.DB)

	// Initialize controllers
	userController := controller.NewUserController(userService, googleService)
	playerController := controller.NewPlayerController(playerService)
	feedbackController := controller.NewFeedbackController(feedbackService)

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

	//Create Player
	router.POST("/api/player/create", func(ctx *gin.Context) {
		playerController.HandleCreatePlayer(ctx.Writer, ctx.Request)
	})

	//Get Player
	router.GET("/api/player/get", func(ctx *gin.Context) {
		playerController.HandleGetPlayer(ctx.Writer, ctx.Request)
	})

	//Update Player
	router.POST("/api/player/update", func(ctx *gin.Context) {
		playerController.HandleUpdatePlayerStats(ctx.Writer, ctx.Request)
	})

	//Create Feedback
	router.POST("/api/feedback", func(ctx *gin.Context) {
		feedbackController.HandleCreateFeedback(ctx.Writer, ctx.Request)
	})

	//Get Scenario
	router.GET("/api/scenario", func(ctx *gin.Context) {
		scenarioController := controller.NewScenarioController(services.NewScenarioService(database.DB))
		scenarioController.GetScenarioByID(ctx.Writer, ctx.Request)
	})
}