package services

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gin/models"
	"math/rand"
	"strconv"
)

type PlayerService struct {
    DB *sql.DB
}

// NewPlayerService creates a new instance of PlayerService
func NewPlayerService(db *sql.DB) *PlayerService {
    return &PlayerService{DB: db}
}

// CreatePlayer creates a new player in the database
func (ps *PlayerService) CreatePlayer(userID int, name string, gender string, race string, school string) (*models.Player, error) {
    randomLuck := rand.Intn(101) // Generate a random number between 0 and 100

    player := &models.Player{
        UserID:          userID,
        Name:            name,
		Race:			 race,
		Gender: 		 gender,
        School:          school,
        StudentType:     "",
        Intelligence:    50,
        Charisma:        50,
        Popularity:      50,
        Strength:        50,
        Luck:            randomLuck,
        CurrentScenario: 0 ,
        EventHistory:    []int{},
    }

    query := `
        INSERT INTO players (user_id, name, race, gender, school, student_type, intelligence, charisma, popularity, strength, luck, current_scenario, event_history)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
        RETURNING id
    `
    err := ps.DB.QueryRow(query, player.UserID, player.Name, player.Race, player.Gender, player.School, player.StudentType, player.Intelligence, player.Charisma, player.Popularity, player.Strength, player.Luck, player.CurrentScenario, "{}").Scan(&player.Id)
    if err != nil {
        return nil, fmt.Errorf("failed to create player: %v", err)
    }

    return player, nil
}

// GetPlayer retrieves a player by ID
func (ps *PlayerService) GetPlayer(playerID int) (*models.Player, error) {
    player := &models.Player{}
    query := `
        SELECT id, user_id, name, race, gender, school, student_type, intelligence, charisma, popularity, strength, luck, current_scenario, event_history
        FROM players
        WHERE id = $1
    `
    row := ps.DB.QueryRow(query, playerID)
    var eventHistory []byte
    err := row.Scan(&player.Id, &player.UserID, &player.Name, &player.Race, &player.Gender, &player.Intelligence, &player.Charisma, &player.Popularity, &player.Strength, &player.Luck, &player.CurrentScenario, &eventHistory)
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

// UpdatePlayer replaces a player's attributes with new values 
func (ps *PlayerService) UpdatePlayerStats(playerID int, newStats map[string]int, newScenario int) error {
    // Ensure all expected stats are provided
    requiredStats := []string{"intelligence", "charisma", "popularity", "strength"}
    for _, stat := range requiredStats {
        if _, exists := newStats[stat]; !exists {
            return fmt.Errorf("missing required stat: %s", stat)
        }
    }

    // Construct the SQL query dynamically
    query := "UPDATE players SET "
    args := []interface{}{}
    argIndex := 1

    for stat, value := range newStats {
        if argIndex > 1 {
            query += ", "
        }
        query += stat + " = $" + strconv.Itoa(argIndex)
        args = append(args, value)
        argIndex++
    }

    // Append the current scenario to event_history and update current_scenario
    query += ", event_history = array_append(event_history, current_scenario), current_scenario = $" + strconv.Itoa(argIndex)
    args = append(args, newScenario)
    argIndex++

    query += " WHERE id = $" + strconv.Itoa(argIndex)
    args = append(args, playerID)

    // Execute the query
    _, err := ps.DB.Exec(query, args...)
    return err
}


