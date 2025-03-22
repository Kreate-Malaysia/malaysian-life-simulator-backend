package user

import (
    "fmt"
    "gin/database"
)

// SaveUser saves the user's email and name into the database
func SaveUser(email, name string) error {
    query := `INSERT INTO users (email, name) VALUES ($1, $2)
              ON CONFLICT (email) DO NOTHING` // Avoid duplicate entries
    _, err := database.DB.Exec(query, email, name)
    if err != nil {
        return fmt.Errorf("failed to save user: %v", err)
    }
    return nil
}