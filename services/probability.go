package services

import (
	"database/sql"
	"gin/models"
	"log"
)

type ProbabilityService struct {
	DB *sql.DB
}

func NewProbabilityService(db *sql.DB) *ProbabilityService {
	return &ProbabilityService{DB: db}
}

func (p *ProbabilityService) GetAllRelatedProbabilities(choiceId int) ([]models.Probability, error) {

	rows, err := p.DB.Query(`SELECT next_scenario, probability FROM probabilities WHERE choice_id = $1`, choiceId)
	if err != nil {
		log.Println("Error querying probabilities:", err)
		return nil, err
	}

	// Defer closing the connection
	defer rows.Close()

	var probabilities []models.Probability
	for rows.Next() {
		var probability models.Probability
		err := rows.Scan(&probability.NextScenario, &probability.Probability)
		if err != nil {
			log.Println("Error scanning probability row:", err)
			return nil, err
		}
		probabilities = append(probabilities, probability)
	}
	return probabilities, nil
}