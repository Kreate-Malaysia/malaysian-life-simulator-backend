package models

type Feedback struct {
	modelImpl
	ScenarioId int    `json:"scenario_id"`
	Feedback   string `json:"feedback"`
}

// Get and Set methods for ScenarioId
func (f *Feedback) GetScenarioId() int {
	return f.ScenarioId
}

func (f *Feedback) SetScenarioId(scenarioId int) {
	f.ScenarioId = scenarioId
}

// Get and Set methods for Feedback
func (f *Feedback) GetFeedback() string {
	return f.Feedback
}

func (f *Feedback) SetFeedback(feedback string) {
	f.Feedback = feedback
}
