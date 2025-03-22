package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gin/models"
	"math/rand"
)

type PlayerService struct {
    DB *sql.DB
}

// NewPlayerService creates a new instance of PlayerService
func NewPlayerService(db *sql.DB) *PlayerService {
    return &PlayerService{DB: db}
}

// CreatePlayer creates a new player in the database
func (ps *PlayerService) CreatePlayer(userID int, name string, gender string, race string) (*models.Player, error) {
    randomLuck := rand.Intn(101) // Generate a random number between 0 and 100

    player := &models.Player{
        UserID:          userID,
        Name:            name,
		Race:			 race,
		Gender: 		 gender,
        Intelligence:    50,
        Charisma:        50,
        Popularity:      50,
        Strength:        50,
        Wealth:          50,
        Luck:            randomLuck,
        CurrentScenario: 0 ,
        EventHistory:    []int{},
    }

    query := `
        INSERT INTO players (user_id, name, race, gender, intelligence, charisma, popularity, strength, wealth, luck, current_scenario, event_history)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
        RETURNING id
    `
    err := ps.DB.QueryRow(query, player.UserID, player.Name, player.Race,player.Gender, player.Intelligence, player.Charisma, player.Popularity, player.Strength, player.Wealth, player.Luck, player.CurrentScenario, "{}").Scan(&player.ID)
    if err != nil {
        return nil, fmt.Errorf("failed to create player: %v", err)
    }

    return player, nil
}

// GetPlayer retrieves a player by ID
func (ps *PlayerService) GetPlayer(playerID int) (*models.Player, error) {
    player := &models.Player{}
    query := `
        SELECT id, user_id, name, race, gender, intelligence, charisma, popularity, strength, wealth, luck, current_scenario, event_history
        FROM players
        WHERE id = $1
    `
    row := ps.DB.QueryRow(query, playerID)
    var eventHistory []byte
    err := row.Scan(&player.ID, &player.UserID, &player.Name, &player.Race, &player.Gender, &player.Intelligence, &player.Charisma, &player.Popularity, &player.Strength, &player.Wealth, &player.Luck, &player.CurrentScenario, &eventHistory)
    if err == sql.ErrNoRows {
        return nil, fmt.Errorf("player not found")
    } else if err != nil {
        return nil, fmt.Errorf("failed to retrieve player: %v", err)
    }

    // Convert eventHistory from JSON to []int
    if err := json.Unmarshal(eventHistory, &player.EventHistory); err != nil {
        return nil, fmt.Errorf("failed to parse event history: %v", err)
    }

    return player, nil
}

// UpdatePlayer updates a player's attributes in the database
func (ps *PlayerService) UpdatePlayerStats(playerID, choiceID int) error {
    var intelligenceChange, strengthChange, charismaChange, popularityChange, scenarioId int

    // Fetch stat modifications from the choice
    err := ps.DB.QueryRow(`
        SELECT intelligence_change, strength_change, charisma_change, popularity_change, scenario_id
        FROM choices WHERE id = $1
    `, choiceID).Scan(&intelligenceChange, &strengthChange, &charismaChange, &popularityChange, &scenarioId)
    if err != nil {
        return err
    }

    // Update player stats
    _, err = ps.DB.Exec(`
        UPDATE players SET 
            intelligence = intelligence + $1,
            strength = strength + $2,
            charisma = charisma + $3
            popularity = popularity + $4
            current_scenario = $5
            event_history = array_append(event_history, $6)
        WHERE id = $7
    `, intelligenceChange, strengthChange, charismaChange, popularityChange, scenarioId, choiceID, playerID)

    return err
}