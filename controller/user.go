package controller

import (
	"encoding/json"
	"fmt"
	"gin/services"
	"net/http"
)

type UserController struct {
	UserService *services.UserService
    GoogleService *services.GoogleOAuthService
}

func NewUserController(userService *services.UserService, googleService *services.GoogleOAuthService) *UserController {
	return &UserController{UserService: userService, GoogleService: googleService}
}

// HandleLogin handles the login endpoint
func (u *UserController) HandleLogin(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var payload struct {
        Email string `json:"email"`
    }
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || payload.Email == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    token, err := u.UserService.Login(payload.Email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// HandleSignup handles the signup endpoint
func(u *UserController) HandleSignup(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var payload struct {
        Email string `json:"email"`
        Name  string `json:"name"`
    }
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || payload.Email == "" || payload.Name == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    token, err := u.UserService.Signup(payload.Email, payload.Name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// LoginWithGoogleAccessToken logs in a user using their Google access token
func(u *UserController) HandleLoginWithGoogleAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var payload struct {
        AccessToken string `json:"access_token"`
    }

	err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || payload.AccessToken == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	// Validate the Google access token and extract the user's email and name
	email, _, err := u.GoogleService.ValidateGoogleAccessToken(payload.AccessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to validate Google access token: %v", err), http.StatusUnauthorized)
		return 
	}

    token, err := u.UserService.Login(email)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

// SignupWithGoogleAccessToken signs up a user using their Google access token
func (u *UserController) HandleSignupWithGoogleAccessToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

	var payload struct {
        AccessToken string `json:"access_token"`
    }

	err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || payload.AccessToken == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

	// Validate the Google access token and extract the user's email and name
	email, name, err := u.GoogleService.ValidateGoogleAccessToken(payload.AccessToken)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to validate Google access token: %v", err), http.StatusUnauthorized)
		return 
	}

    token, err := u.UserService.Signup(email, name)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}