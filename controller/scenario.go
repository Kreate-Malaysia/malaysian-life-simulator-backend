package controller

import (
	"encoding/json"
	"gin/services"
	"net/http"
)

type ScenarioController struct {
    ScenarioService *services.ScenarioService
}

// NewScenarioController creates a new instance of ScenarioController
func NewScenarioController(scenarioService *services.ScenarioService) *ScenarioController {
    return &ScenarioController{ScenarioService: scenarioService}
}

// GetScenarioByID handles the HTTP POST request to retrieve a scenario by its ID from the request body
func (sc *ScenarioController) GetScenarioByID(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        ID int `json:"id"`
    }

    // Decode the JSON request body
    err := json.NewDecoder(r.Body).Decode(&payload)
    if (err != nil || payload.ID <= 0) {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call the service to retrieve the scenario
    scenario, err := sc.ScenarioService.GetScenarioByID(payload.ID)
    if err != nil {
        http.Error(w, "Failed to retrieve scenario", http.StatusInternalServerError)
        return
    }

    // If no scenario is found, return a 404 response
    if scenario == nil {
        http.Error(w, "Scenario not found", http.StatusNotFound)
        return
    }

    // Respond with the scenario data in JSON format
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(scenario)
}