package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

// GoogleOAuthService struct
type GoogleOAuthService struct {
	DB *sql.DB
}

func NewGoogleOAuthService(db *sql.DB) *GoogleOAuthService {
	return &GoogleOAuthService{DB: db}
}

// GoogleTokenPayload represents the payload received from the frontend
type GoogleTokenPayload struct {
    AccessToken string `json:"access_token"`
    Name        string `json:"name"`
}

// ValidateGoogleAccessToken validates the Google access token and retrieves the user's email and name
func (goa *GoogleOAuthService) ValidateGoogleAccessToken(accessToken string) (string, string, error) {
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