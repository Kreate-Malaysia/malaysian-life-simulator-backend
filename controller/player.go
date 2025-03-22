package controller

import (
	"encoding/json"
	"gin/services"
	"net/http"
)

type PlayerController struct {
    PlayerService *services.PlayerService
    UserService   *services.UserService
}

// NewPlayerController creates a new instance of PlayerController
func NewPlayerController(playerService *services.PlayerService, UserService *services.UserService) *PlayerController {
    return &PlayerController{PlayerService: playerService}
}

// CreatePlayer handles the creation of a new player
func (pc *PlayerController) HandleCreatePlayer(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        Name        string `json:"name"`
        Gender      string `json:"gender"`
        Race        string `json:"race"`
        SchoolType  string `json:"school_type"`
    }

    // Decode the JSON request body
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Extract the JWT token from the Authorization header
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
        return
    }

    // Remove the "Bearer " prefix from the token
    const bearerPrefix = "Bearer "
    if len(authHeader) <= len(bearerPrefix) || authHeader[:len(bearerPrefix)] != bearerPrefix {
        http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
        return
    }
    jwtToken := authHeader[len(bearerPrefix):]

    // Decode the JWT token to extract user_id
    decodedData, err := pc.UserService.DecodeJWT(jwtToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    // Extract user_id from the claims
    userID, ok := decodedData["user_id"].(int) // JWT claims are parsed as float64
    if !ok {
        http.Error(w, "Invalid user_id in token", http.StatusUnauthorized)
        return
    }

    // Call the PlayerService to create a new player
    player, err := pc.PlayerService.CreatePlayer(int(userID), requestBody.Name, requestBody.Gender, requestBody.Race)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with the created player
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(player)
}

// GetPlayer handles retrieving a player by ID
func (pc *PlayerController) HandleGetPlayer(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        PlayerID int `json:"player_id"`
    }

    // Decode the request body into requestBody struct
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Fetch player from the service
    player, err := pc.PlayerService.GetPlayer(requestBody.PlayerID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    // Send response as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(player)
}

// UpdatePlayerStatsHandler handles the request to update a player's stats
func (pc *PlayerController) HandleUpdatePlayerStats(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        PlayerID   int            `json:"player_id"`
        NewStats   map[string]int `json:"new_stats"`
        NewScenario int           `json:"new_scenario"`
    }

    // Decode JSON request body
    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Ensure all required fields are provided
    if requestBody.PlayerID == 0 {
        http.Error(w, "Missing player_id", http.StatusBadRequest)
        return
    }

    if requestBody.NewScenario == 0 {
        http.Error(w, "Missing new_scenario", http.StatusBadRequest)
        return
    }

    // Call service function to update stats
    err := pc.PlayerService.UpdatePlayerStats(requestBody.PlayerID, requestBody.NewStats, requestBody.NewScenario)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Send success response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "Player stats updated successfully"})
}
