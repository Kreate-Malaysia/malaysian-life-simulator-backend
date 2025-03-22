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
		err := rows.Scan(&scenario.Id, &scenario.Description)
		if err != nil {
			log.Println("Error scanning scenario row:", err)
			return nil, err
		}
		scenarios = append(scenarios, scenario)
	}

	return scenarios, nil
}

// GetScenarioByID retrieves a single scenario by its ID
func (sc *ScenarioService) GetScenarioByID(id int) (*models.Scenario, error) {
    query := `
        SELECT id, description, is_choice, is_story, is_random, leads_to, is_conditional
        FROM scenarios
        WHERE id = $1
    `
    row := sc.DB.QueryRow(query, id)

    var scenario models.Scenario
    err := row.Scan(&scenario.Id, &scenario.Description, &scenario.IsChoice, &scenario.IsStory, &scenario.IsRandom, &scenario.LeadsTo, &scenario.IsConditional)
    if err == sql.ErrNoRows {
        log.Println("No scenario found with the given ID:", id)
        return nil, nil
    } else if err != nil {
        log.Println("Error querying scenario by ID:", err)
        return nil, err
    }

    return &scenario, nil
}