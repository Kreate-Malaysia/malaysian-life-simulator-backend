package controller

import (
	"encoding/json"
	"gin/services"
	"net/http"
)

type PlayerController struct {
    PlayerService *services.PlayerService
}

// NewPlayerController creates a new instance of PlayerController
func NewPlayerController(playerService *services.PlayerService) *PlayerController {
    return &PlayerController{PlayerService: playerService}
}

// CreatePlayer handles the creation of a new player
func (pc *PlayerController) HandleCreatePlayer(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        UserID int    `json:"user_id"`
        Name   string `json:"name"`
        Gender string `json:"gender"`
        Race   string `json:"race"`
    }

    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || payload.UserID == 0 || payload.Name == "" || payload.Gender == "" || payload.Race == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    player, err := pc.PlayerService.CreatePlayer(payload.UserID, payload.Name, payload.Gender, payload.Race)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

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
