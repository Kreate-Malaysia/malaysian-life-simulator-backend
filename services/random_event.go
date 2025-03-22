package services

import (
	"database/sql"
	"fmt"
	"gin/models"
	"math/rand"
	"time"
)

type RandomEventService struct {
	DB *sql.DB
}

func NewRandomEventService(db *sql.DB) *RandomEventService {
	return &RandomEventService{DB: db}
}

// Get all random events for a given scenario
func (res *RandomEventService) GetRandomEvents(scenarioID int) ([]models.RandomEvent, error) {
	query := `SELECT id, description, probability, leads_to FROM random_events WHERE scenario_id = $1`
	
	rows, err := res.DB.Query(query, scenarioID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch random events: %v", err)
	}
	defer rows.Close()

	var events []models.RandomEvent
	for rows.Next() {
		var event models.RandomEvent
		err := rows.Scan(&event.Id, &event.Description, &event.Probability, &event.LeadsTo)
		if err != nil {
			return nil, fmt.Errorf("failed to scan random event: %v", err)
		}
		events = append(events, event)
	}

	return events, nil
}

// Roll probability to determine the next scenario
func (res *RandomEventService) SelectRandomEvent(events []models.RandomEvent) *models.RandomEvent {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	roll := r.Float64()
	var cumulativeProbability float64 = 0

	for _, event := range events {
		cumulativeProbability += event.Probability
		if roll <= cumulativeProbability {
			return &event
		}
	}

	return &events[len(events)-1]
}