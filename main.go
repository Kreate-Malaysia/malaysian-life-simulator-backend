package main

import (
	"gin/config"
	"gin/database"
	"log"

	_ "github.com/lib/pq"
)
func main() {
	
	// Load environment config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Initialize databse connection
	database.InitDB(cfg)
}