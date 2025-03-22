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

// // GetVocabWithCategory fetches vocab with category names
// func (s *VocabService) GetVocabWithCategory() ([]models.Vocab, error) {
// 	query := `
// 		SELECT v.id, v.word, v.meaning, c.name AS category 
// 		FROM vocab v
// 		JOIN categories c ON v.category_id = c.id
// 	`

// 	rows, err := s.DB.Query(query)
// 	if err != nil {
// 		log.Println("Error querying vocab:", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

	// var vocabs []models.Vocab
	// for rows.Next() {
	// 	var vocab models.Vocab
	// 	err := rows.Scan(&vocab.ID, &vocab.Word, &vocab.Meaning, &vocab.Category)
	// 	if err != nil {
	// 		log.Println("Error scanning vocab row:", err)
	// 		return nil, err
	// 	}
	// 	vocabs = append(vocabs, vocab)
	// }

	// return vocabs, nil
// }