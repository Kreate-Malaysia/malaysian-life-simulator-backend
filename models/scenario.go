package models

type Scenario struct {
    modelImpl
    ID            int    `json:"id"`
    Description   string `json:"description"`
    IsChoice      bool   `json:"is_choice"`
    IsStory       bool   `json:"is_story"`
    IsRandom      bool   `json:"is_random"`
    LeadsTo       bool   `json:"leads_to"`
    IsConditional bool   `json:"is_conditional"`
}

// Get and Set methods for ID
func (sc *Scenario) GetID() int {
    return sc.ID
}

func (sc *Scenario) SetID(id int) {
    sc.ID = id
}

// Get and Set methods for Description
func (sc *Scenario) GetDescription() string {
    return sc.Description
}

func (sc *Scenario) SetDescription(description string) {
    sc.Description = description
}

// Get and Set methods for IsChoice
func (sc *Scenario) GetIsChoice() bool {
    return sc.IsChoice
}

func (sc *Scenario) SetIsChoice(isChoice bool) {
    sc.IsChoice = isChoice
}

// Get and Set methods for IsStory
func (sc *Scenario) GetIsStory() bool {
    return sc.IsStory
}

func (sc *Scenario) SetIsStory(isStory bool) {
    sc.IsStory = isStory
}

// Get and Set methods for IsRandom
func (sc *Scenario) GetIsRandom() bool {
    return sc.IsRandom
}

func (sc *Scenario) SetIsRandom(isRandom bool) {
    sc.IsRandom = isRandom
}

// Get and Set methods for LeadsTo
func (sc *Scenario) GetLeadsTo() bool {
    return sc.LeadsTo
}

func (sc *Scenario) SetLeadsTo(leadsTo bool) {
    sc.LeadsTo = leadsTo
}

// Get and Set methods for IsConditional
func (sc *Scenario) GetIsConditional() bool {
    return sc.IsConditional
}

func (sc *Scenario) SetIsConditional(isConditional bool) {
    sc.IsConditional = isConditional
}