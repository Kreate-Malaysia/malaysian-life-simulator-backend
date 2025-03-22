package models

type Probability struct {
	modelImpl
	ChoiceId     	int    `json:"choice_id"`
	NextScenario    int    `json:"next_scenario"`
	Probability		int    `json:"probability"`
}

func (p *Probability) GetChoiceId() int {
	return p.ChoiceId;
}

func (p *Probability) GetNextScenario() int {
	return p.NextScenario;
}

func (p *Probability) GetProbability() int {
	return p.Probability;
}