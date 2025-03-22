package models

type Vocabulary struct {
	modelImpl
	Word       string   `json:"name"`
	ScenarioId int   	`json:"scenario_id"` // Foreign key reference
	Scenario   string 	`json:"scenarios,omitempty"` // Scenario name (optional)
}

func (v *Vocabulary) GetWord() string {
	return v.Word
}