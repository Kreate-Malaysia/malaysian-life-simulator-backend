package models

type RandomEvent struct {
	modelImpl
	ScenarioId    int    	`json:"scenario_id"`
	Description   string 	`json:"description"`
	Probability   float64   `json:"probability"`
	LeadsTo       int    	`json:"leads_to"`
	Scenario   	  string 	`json:"scenarios,omitempty"` // Scenario name (optional)
}

func (re *RandomEvent) GetScenarioId() int {
	return re.ScenarioId
}

func (re *RandomEvent) GetDescription() string {
	return re.Description
}

func (re *RandomEvent) GetProbability() float64 {
	return re.Probability
}

func (re *RandomEvent) GetLeadsTo() int {
	return re.LeadsTo
}

func (re *RandomEvent) GetScenario() string {
	return re.Scenario
}

// 