package services

import (
	"database/sql"
	"fmt"
	"gin/database"
)

// UserService struct
type UserService struct {
	DB *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{DB: db}
}

// SaveUser saves the user's email and name into the database
func (u *UserService) SaveUser(email, name string) error {
    query := `INSERT INTO users (email, name) VALUES ($1, $2)
              ON CONFLICT (email) DO NOTHING` // Avoid duplicate entries
    _, err := database.DB.Exec(query, email, name)
    if err != nil {
        return fmt.Errorf("failed to save user: %v", err)
    }
    return nil
}