package services

import (
	"database/sql"
	"gin/models"
	"log"
)

// ChoiceService struct
type ChoiceService struct {
	DB *sql.DB
}

func NewChoiceService(db *sql.DB) *ChoiceService {
	return &ChoiceService{DB: db}
}

func (v*ChoiceService) GetAllChoices() ([]models.Choice, error) {
	query := `
	SELECT c.id, c.choice_text, c.intelligence_change, c.charisma_change, c.popularity_change, c.strength_change, sc.name AS scenario
	FROM choice c
	JOIN scenarios sc ON c.scenario_id = sc.id
`
	rows, err := v.DB.Query(query)
	if err != nil {
		log.Println("Error querying vocab:", err)
		return nil, err
	}

		// Defer closing the connection
		defer rows.Close()
		// Loop and append rows into an array
		var choices []models.Choice
		for rows.Next() {
			var choice models.Choice
			err := rows.Scan(&choice.Id, &choice.ChoiceText, &choice.IntelligenceChange, &choice.CharismaChange, &choice.PopularityChange, &choice.StrengthChange, &choice.Scenario)
			if err != nil {
				log.Println("Error scanning choice row:", err)
				return nil, err
			}
			choices = append(choices, choice)
		}
	
		return choices, nil
}

func (v*ChoiceService) GetChoices(scenarioID int) ([]models.Choice, error) {
	query := `
		SELECT c.id, c.choice_text, c.intelligence_change, c.charisma_change, c.popularity_change, c.strength_change, sc.name AS scenario
		FROM choice c
		JOIN scenarios sc ON c.scenario_id = sc.id
		WHERE sc.id = $1
	`

	rows, err := v.DB.Query(query, scenarioID)
	if err != nil {
		log.Println("Error querying choices:", err)
		return nil, err
	}

	// Defer closing the connection
	defer rows.Close()
	// Loop and append rows into an array
	var choices []models.Choice
	for rows.Next() {
		var choice models.Choice
		err := rows.Scan(&choice.Id, &choice.ChoiceText, &choice.IntelligenceChange, &choice.CharismaChange, &choice.PopularityChange, &choice.StrengthChange, &choice.Scenario)
		if err != nil {
			log.Println("Error scanning choice row:", err)
			return nil, err
		}
		choices = append(choices, choice)
	}

	return choices, nil
}