package googleoauth

import (
	"encoding/json"
	"fmt"
	"gin/config"
	"gin/services"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte // JWT secret extracted from cfg

type GoogleController struct {
	UserService *services.UserService
}

func NewGoogleController(userService *services.UserService) *GoogleController {
	return &GoogleController{UserService: userService}
}


// InitializeGoogleOAuth initializes the package-level configuration
func InitializeGoogleOAuth(cfg *config.Config) {
	jwtSecret = []byte(cfg.JWT.JWTSecret)
}

// GoogleTokenPayload represents the payload received from the frontend
type GoogleTokenPayload struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

// JWTResponse represents the response containing the JWT token
type JWTResponse struct {
	Token string `json:"token"`
}

// GenerateJWT generates a JWT token for the given email
func GenerateJWT(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// ValidateGoogleAccessToken validates the Google access token and retrieves the user's email and name
func ValidateGoogleAccessToken(accessToken string) (string, string, error) {
	resp, err := http.Get(fmt.Sprintf("https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=%s", accessToken))
	if err != nil {
		return "", "", fmt.Errorf("failed to validate access token: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("invalid access token, status code: %d", resp.StatusCode)
	}

	var tokenInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tokenInfo); err != nil {
		return "", "", fmt.Errorf("failed to parse token info response: %v", err)
	}

	if tokenInfo.Email == "" {
		return "", "", fmt.Errorf("email not found in token info")
	}

	return tokenInfo.Email, tokenInfo.Name, nil
}

// HandleOAuthCallback handles the request from the frontend
func (c *GoogleController)  HandleOAuthCallback(w http.ResponseWriter, r *http.Request) {
	
	// Parse the JSON payload from the frontend
	var payload GoogleTokenPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || payload.AccessToken == "" {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	email, name, err := ValidateGoogleAccessToken(payload.AccessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to validate access token: %v", err), http.StatusUnauthorized)
		return
	}

	// Call the user service to save the user
	err = c.UserService.SaveUser(email, name)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to save user: %v", err), http.StatusInternalServerError)
		return
	}

	// Generate a JWT token for the user
	token, err := GenerateJWT(email)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	response := JWTResponse{Token: token}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}