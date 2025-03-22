package models

type ConditionalEvent struct {
	modelImpl
	ConditionOne  	string    	`json:"condition_one"`
	ConditionTwo   	string 		`json:"condition_two"`
	ConditionThree  string    	`json:"condition_three,omitempty"`
	LeadsToIfOne    int    		`json:"leads_to_if_one"`
	LeadsToIfTwo   	int			`json:"leads_to_if_two"`
	LeadsToIfThree  int 		`json:"leads_to_if_three,omitempty"`
	Scenario		string 		`json:"scenarios,omitempty"` // Scenario name (optional)
}

func (ce *ConditionalEvent) GetConditionOne() string {
	return ce.ConditionOne
}

func (ce *ConditionalEvent) GetConditionTwo() string {
	return ce.ConditionTwo
}

func (ce *ConditionalEvent) GetConditionThree() string {
	return ce.ConditionThree
}

func (ce *ConditionalEvent) GetLeadsToIfOne() int {
	return ce.LeadsToIfOne
}

func (ce *ConditionalEvent) GetLeadsToIfTwo() int {
	return ce.LeadsToIfTwo
}

func (ce *ConditionalEvent) GetLeadsToIfThree() int {
	return ce.LeadsToIfThree
}

func (ce *ConditionalEvent) GetScenario() string {
	return ce.Scenario
}