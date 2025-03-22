package models

type Choice struct {
	modelImpl
	ChoiceText      	string  `json:"choice_text"`
	IntelligenceChange 	int   	`json:"intelligence_change"`
	CharismaChange	 	int   	`json:"charisma_change"`
	PopularityChange 	int   	`json:"popularity_change"`
	StrengthChange	 	int   	`json:"strength_change"`
	ScenarioId		 	int   	`json:"scenario_id"`
	Scenario   			string 	`json:"scenarios,omitempty"` // Scenario name (optional)
}

func (v *Choice) GetChoiceText() string {
	return v.ChoiceText
}

func (v *Choice) GetIntelligenceChange() int {
	return v.IntelligenceChange
}

func (v *Choice) GetCharismaChange() int {
	return v.CharismaChange
}

func (v *Choice) GetPopularityChange() int {
	return v.PopularityChange
}

func (v *Choice) GetStrengthChange() int {
	return v.StrengthChange
}

func (v *Choice) GetScenarioId() int {
	return v.ScenarioId
}