package main

import (
    "gin/config"
    "gin/database"
    "gin/googleoauth"
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

    // Register the route for handling OAuth callback
    http.HandleFunc("/api/oauth/callback", googleoauth.HandleOAuthCallback)

    // Start the HTTP server
    log.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}