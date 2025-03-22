package main

import (
	"gin/config"
	"gin/controller"
	"gin/database"
	"gin/services"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
    // Load environment config
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize database connection
    database.InitDB(cfg)
    
	// Initialize services
    googleService := services.NewGoogleOAuthService(database.DB)
	userService := services.NewUserService(database.DB,googleService)
    
	userController := controller.NewUserController(userService,googleService)
    
    
    // Register the route for login
    http.HandleFunc("/api/login", userController.HandleLogin)
    // Register the route for signup
    http.HandleFunc("/api/signup", userController.HandleSignup)
    // Register the route for login using Google OAuth
    http.HandleFunc("/api/google/login", userController.HandleLoginWithGoogleAccessToken)
    // Register the route for signup using Google OAuth
    http.HandleFunc("/api/google/signup", userController.HandleSignupWithGoogleAccessToken)

    // Start the HTTP server
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}