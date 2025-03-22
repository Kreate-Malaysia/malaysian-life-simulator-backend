package services

import (
	"database/sql"
	"gin/models"
	"log"
)

// VocabularyService struct
type VocabularyService struct {
	DB *sql.DB
}

func NewVocabularyService(db *sql.DB) *VocabularyService {
	return &VocabularyService{DB: db}
}

func (v*VocabularyService) GetAllVocabularies() ([]models.Vocabulary, error) {
	query := `
	SELECT v.id, v.word, sc.name AS scenario
	FROM scenario_vocabularies v
	JOIN scenarios sc ON v.scenario_id = sc.id
`
	rows, err := v.DB.Query(query)
	if err != nil {
		log.Println("Error querying vocab:", err)
		return nil, err
	}

		// Defer closing the connection
		defer rows.Close()
		// Loop and append rows into an array
		var vocabularies []models.Vocabulary
		for rows.Next() {
			var vocab models.Vocabulary
			err := rows.Scan(&vocab.Id, &vocab.Word, &vocab.Scenario)
			if err != nil {
				log.Println("Error scanning vocabulary row:", err)
				return nil, err
			}
			vocabularies = append(vocabularies, vocab)
		}
	
		return vocabularies, nil
}