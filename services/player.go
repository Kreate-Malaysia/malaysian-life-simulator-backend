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
func (ps *PlayerService) UpdatePlayer(player *models.Player) error {
    query := `
        UPDATE players
        SET intelligence = $1, charisma = $2, popularity = $3, strength = $4, wealth = $5, luck = $6, current_scenario = $7, event_history = $8
        WHERE id = $9
    `
    eventHistory, err := json.Marshal(player.EventHistory)
    if err != nil {
        return fmt.Errorf("failed to serialize event history: %v", err)
    }

    _, err = ps.DB.Exec(query, player.Intelligence, player.Charisma, player.Popularity, player.Strength, player.Wealth, player.Luck, player.CurrentScenario, eventHistory, player.ID)
    if err != nil {
        return fmt.Errorf("failed to update player: %v", err)
    }

    return nil
}