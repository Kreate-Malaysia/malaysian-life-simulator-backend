package main

import (
	"gin/config"
	"gin/database"
	"gin/googleoauth"
	"gin/services"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
    // Load environment config
    cfg, err := config.LoadConfig()
    if (err != nil) {
        log.Fatalf("Error loading config: %v", err)
    }

    // Initialize database connection
    database.InitDB(cfg)

    // Initialize Google OAuth with the loaded config
    googleoauth.InitializeGoogleOAuth(cfg)

	// Initialize services
	userService := services.NewUserService(database.DB)

	googleController := googleoauth.NewGoogleController(userService)

    // Register the route for handling OAuth callback
    http.HandleFunc("/api/oauth/callback", googleController.HandleOAuthCallback)

    // Start the HTTP server
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}