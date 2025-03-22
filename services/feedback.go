package services

import (
	"database/sql"
	"fmt"
	"gin/models"
)

type FeedbackService struct {
    DB *sql.DB
}

// NewFeedbackService creates a new instance of FeedbackService
func NewFeedbackService(db *sql.DB) *FeedbackService {
    return &FeedbackService{DB: db}
}

// CreateFeedback creates a new feedback entry in the database
func (fs *FeedbackService) CreateFeedback(scenarioId int, feedbackText string) (*models.Feedback, error) {
    feedback := &models.Feedback{
        ScenarioId: scenarioId,
        Feedback:   feedbackText,
    }

    query := `
        INSERT INTO feedback (scenario_id, feedback)
        VALUES ($1, $2)
        RETURNING id
    `
    err := fs.DB.QueryRow(query, feedback.ScenarioId, feedback.Feedback).Scan(&feedback.Id)
    if (err != nil) {
        return nil, fmt.Errorf("failed to create feedback: %v", err)
    }

    return feedback, nil
}
