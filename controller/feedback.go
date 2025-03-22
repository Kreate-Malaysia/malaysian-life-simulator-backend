package controller

import (
	"encoding/json"
	"gin/services"
	"net/http"
)

type FeedbackController struct {
    FeedbackService *services.FeedbackService
}

// NewFeedbackController creates a new instance of FeedbackController
func NewFeedbackController(feedbackService *services.FeedbackService) *FeedbackController {
    return &FeedbackController{FeedbackService: feedbackService}
}

// CreateFeedback handles the POST request to create feedback
func (fc *FeedbackController) HandleCreateFeedback(w http.ResponseWriter, r *http.Request) {
    var payload struct {
        ScenarioId int    `json:"scenario_id"`
        Feedback   string `json:"feedback"`
    }

    // Decode the JSON request body
    err := json.NewDecoder(r.Body).Decode(&payload)
    if err != nil || payload.ScenarioId == 0 || payload.Feedback == "" {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Call the FeedbackService to create feedback
    feedback, err := fc.FeedbackService.CreateFeedback(payload.ScenarioId, payload.Feedback)
    if err != nil {
        http.Error(w, "Failed to create feedback", http.StatusInternalServerError)
        return
    }

    // Respond with the created feedback
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(feedback)
}