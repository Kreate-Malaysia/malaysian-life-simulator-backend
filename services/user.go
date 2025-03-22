package services

import (
	"database/sql"
	"fmt"
	"gin/config"
	"time"

	"github.com/golang-jwt/jwt"
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
func (u *UserService) GenerateJWT(email, name string, userId int) (string, error) {
    jwtSecret, err := u.getJWTSecret()
    if err != nil {
        return "", fmt.Errorf("failed to get JWT secret: %v", err)
    }

    claims := jwt.MapClaims{
		"user_id": userId,
        "email": email,
        "name":  name,
        "exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtSecret)
}

// DecodeJWT decodes a JWT token and extracts its claims
func (u *UserService) DecodeJWT(tokenString string) (map[string]interface{}, error) {
    // Check if the token is empty
    if tokenString == "" {
        return nil, fmt.Errorf("token is empty")
    }

    // Get the JWT secret
    jwtSecret, err := u.getJWTSecret()
    if err != nil {
        return nil, fmt.Errorf("failed to get JWT secret: %v", err)
    }

    // Parse the token
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Ensure the signing method is HMAC
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return jwtSecret, nil
    })

    if err != nil {
        return nil, fmt.Errorf("failed to parse token: %v", err)
    }

    // Extract claims
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        // Extract specific fields from claims
        userID, ok1 := claims["user_id"].(float64) // JWT claims are parsed as float64
        email, ok2 := claims["email"].(string)
        name, ok3 := claims["name"].(string)

        if !ok1 || !ok2 || !ok3 {
            return nil, fmt.Errorf("invalid token claims")
        }

        // Return the extracted fields as a map
        return map[string]interface{}{
            "user_id": int(userID), // Convert user_id to int
            "email":   email,
            "name":    name,
        }, nil
    }

    return nil, fmt.Errorf("invalid token")
}

// SaveUser saves the user's email and name into the database and returns the user ID
func (u *UserService) SaveUser(email, name string) (int, error) {
    var userId int
    query := `
        INSERT INTO users (email, name)
        VALUES ($1, $2)
        ON CONFLICT (email) DO NOTHING
        RETURNING id
    `
    err := u.DB.QueryRow(query, email, name).Scan(&userId)
    if err == sql.ErrNoRows {
        // If the user already exists, retrieve the existing user ID
        query = `SELECT id FROM users WHERE email = $1`
        err = u.DB.QueryRow(query, email).Scan(&userId)
        if err != nil {
            return 0, fmt.Errorf("failed to retrieve existing user ID: %v", err)
        }
    } else if err != nil {
        return 0, fmt.Errorf("failed to save user: %v", err)
    }

    return userId, nil
}

// Login checks if the user exists in the database and returns a JWT token if valid
func (u *UserService) Login(email string) (string, error) {
	var userId int
    var name string
    query := `SELECT id, name FROM users WHERE email = $1`
    err := u.DB.QueryRow(query, email).Scan(&userId, &name)
    if err == sql.ErrNoRows {
        return "", fmt.Errorf("user not found")
    } else if err != nil {
        return "", fmt.Errorf("failed to query user: %v", err)
    }

	// Generate JWT token
	token, err := u.GenerateJWT(email, name, userId)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

// Signup handles both Google OAuth and email-based signups
func (u *UserService) Signup(email, name string) (string, error) {
	// Save the user to the database
	userId, err := u.SaveUser(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to save user: %v", err)
	}

	// Generate JWT token
	token, err := u.GenerateJWT(email, name, userId)
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

    // Check if the user exists in the database and retrieve their ID
    var userId int
    query := `SELECT id, name FROM users WHERE email = $1`
    err = u.DB.QueryRow(query, email).Scan(&userId, &name)
    if err == sql.ErrNoRows {
        return "", fmt.Errorf("user not found")
    } else if err != nil {
        return "", fmt.Errorf("failed to query user: %v", err)
    }

    // Generate JWT token
    token, err := u.GenerateJWT(email, name, userId)
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
	userId, err := u.SaveUser(email, name)
	if err != nil {
		return "", fmt.Errorf("failed to save user: %v", err)
	}

	// Generate JWT token
	token, err := u.GenerateJWT(email, name, userId)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}