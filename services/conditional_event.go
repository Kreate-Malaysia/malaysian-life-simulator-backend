package services

import (
	"database/sql"
	"fmt"
	"gin/models"
)

type ConditionalEventService struct {
	DB *sql.DB
}

func NewConditionalEventService(db *sql.DB) *ConditionalEventService {
	return &ConditionalEventService{DB: db}
}

// Get the conditional event for a given scenario
func (ces *ConditionalEventService) GetConditionalEvent(scenarioID int) (*models.ConditionalEvent, error) {
	query := `SELECT id, condition_one, condition_two, condition_three, leads_to_if_one, leads_to_if_two, leads_to_if_three 
			  FROM conditional_events WHERE scenario_id = $1`
	
	var event models.ConditionalEvent
	err := ces.DB.QueryRow(query, scenarioID).Scan(
		&event.Id, &event.ConditionOne, &event.ConditionTwo, &event.ConditionThree,
		&event.LeadsToIfOne, &event.LeadsToIfTwo, &event.LeadsToIfThree)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch conditional event: %v", err)
	}

	return &event, nil
}

// Determine the next scenario based on player's attributes
func (ces *ConditionalEventService) EvaluateConditionalEvent(player *models.Player, event *models.ConditionalEvent) int {
	if player.StudentType == event.ConditionOne {
		return event.LeadsToIfOne
	} else if player.StudentType == event.ConditionTwo {
		return event.LeadsToIfTwo
	} else {
		return event.LeadsToIfThree
	}
}