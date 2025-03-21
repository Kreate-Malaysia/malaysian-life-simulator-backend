package database

import (
	"database/sql"
	"fmt"
	"gin/config"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB(cfg *config.Config) {
	connStr := cfg.DB.Url

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	DB = db

	fmt.Println("Database connected successfully")
}

