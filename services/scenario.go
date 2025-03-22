package services

import (
	"database/sql"
	"gin/models"
	"log"
)

// ScenarioService struct
type ScenarioService struct {
	DB *sql.DB
}

func NewScenarioService(db *sql.DB) *ScenarioService {
	return &ScenarioService{DB: db}
}

func (sc *ScenarioService) GetAllScenario() ([]models.Scenario, error) {
	query := `
		SELECT *
		FROM scenario
	`
	rows, err := sc.DB.Query(query)
	if err != nil {
		log.Println("Error querying scenario:", err)
		return nil, err
	}

	// Defer closing the connection
	defer rows.Close()
	// Loop and append rows into an array
	var scenarios []models.Scenario
	for rows.Next() {
		var scenario models.Scenario
		err := rows.Scan(&scenario.Id, &scenario.Name)
		if err != nil {
			log.Println("Error scanning scenario row:", err)
			return nil, err
		}
		scenarios = append(scenarios, scenario)
	}

	return scenarios, nil
}