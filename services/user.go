package services

import (
	"database/sql"
	"fmt"
	"gin/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserService struct {
	DB *sql.DB
	GoogleService *GoogleOAuthService
}

func NewUserService(db *sql.DB, googleService *GoogleOAuthService) *UserService {
	return &UserService{DB: db, GoogleService: googleService}
}

// getJWTSecret dynamically loads the JWT secret from the configuration
func (u *UserService) getJWTSecret() ([]byte, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %v", err)
	}
	return []byte(cfg.JWT.JWTSecret), nil
}

// GenerateJWT generates a JWT token for the given email and name
func (u *UserService) GenerateJWT(email, name string) (string, error) {
    jwtSecret, err := u.getJWTSecret()
    if err != nil {
        return "", fmt.Errorf("failed to get JWT secret: %v", err)
    }

    claims := jwt.MapClaims{
        "email": email,
        "name":  name,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// SaveUser saves the user's email and name into the database
func (u *UserService) SaveUser(email, name string) error {
	query := `INSERT INTO users (email, name) VALUES ($1, $2)
              ON CONFLICT (email) DO NOTHING` // Avoid duplicate entries
	_, err := u.DB.Exec(query, email, name)
	if err != nil {
		return fmt.Errorf("failed to save user: %v", err)
	}
	return nil
}

// Login checks if the user exists in the database and returns a JWT token if valid
func (u *UserService) Login(email string) (string, error) {
	var name string
	query := `SELECT name FROM users WHERE email = $1`
	err := u.DB.QueryRow(query, email).Scan(&name)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		return "", fmt.Errorf("failed to query user: %v", err)
	}

	// Generate JWT token
	token, err := u.GenerateJWT(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

// Signup handles both Google OAuth and email-based signups
func (u *UserService) Signup(email, name string) (string, error) {
	// Save the user to the database
	err := u.SaveUser(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to save user: %v", err)
	}

	// Generate JWT token
	token, err := u.GenerateJWT(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

// LoginWithGoogleAccessToken logs in a user using their Google access token
func (u *UserService) LoginWithGoogleAccessToken(accessToken string) (string, error) {
	// Validate the Google access token and extract the user's email and name
	email, name, err := u.GoogleService.ValidateGoogleAccessToken(accessToken)
	if err != nil {
		return "", fmt.Errorf("failed to validate Google access token: %v", err)
	}

	// Check if the user exists in the database
	var existingName string
	query := `SELECT name FROM users WHERE email = $1`
	err = u.DB.QueryRow(query, email).Scan(&existingName)
	if err == sql.ErrNoRows {
		return "", fmt.Errorf("user not found")
	} else if err != nil {
		return "", fmt.Errorf("failed to query user: %v", err)
	}

	// Generate JWT token
	token, err := u.GenerateJWT(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

// SignupWithGoogleAccessToken signs up a user using their Google access token
func (u *UserService) SignupWithGoogleAccessToken(accessToken string) (string, error) {
	// Validate the Google access token and extract the user's email and name
	email, name, err := u.GoogleService.ValidateGoogleAccessToken(accessToken)
	if err != nil {
		return "", fmt.Errorf("failed to validate Google access token: %v", err)
	}

	// Save the user to the database
	err = u.SaveUser(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to save user: %v", err)
	}

	// Generate JWT token
	token, err := u.GenerateJWT(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}