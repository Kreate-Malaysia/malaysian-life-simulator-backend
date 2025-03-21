package models

type Player struct {
	modelImpl
	UserID          int    `json:"user_id"`
	Name            string `json:"name"`
	Intelligence    int    `json:"intelligence" default:"50"`
	Charisma        int    `json:"charisma" default:"50"`
	Popularity      int    `json:"popularity" default:"50"`
	Strength        int    `json:"strength" default:"50"`
	Luck            int    `json:"luck" default:"50"`
	CurrentScenario int    `json:"current_scenario"`
	EventHistory    []int  `json:"event_history" default:"[]"`
	Race            string `json:"race"`
	Gender          string `json:"gender"`
	School          string `json:"school"`
	StudentType     string `json:"student_type"`
}

// Get and Set methods for Popularity
func (p *Player) GetPopularity() int {
	return p.Popularity
}

func (p *Player) SetPopularity(value int) {
	p.Popularity = value
}

// Get and Set methods for Strength
func (p *Player) GetStrength() int {
	return p.Strength
}

func (p *Player) SetStrength(value int) {
	p.Strength = value
}

// Get Race
func (p *Player) GetRace() string {
	return p.Race
}

func (p *Player) SetRace(value string) {
	p.Race = value
}

// Get Gender
func (p *Player) GetGender() string {
	return p.Gender
}

func (p *Player) SetGender(value string) {
	p.Gender = value
}

func (p *Player) GetSchool() string {
	return p.School
}

func (p *Player) SetSchool(value string) {
	p.School = value
}

func (p *Player) GetStudentType() string {
	return p.StudentType
}

func (p *Player) SetStudentType(value string) {
	p.StudentType = value
}

// Get and Set methods for Charisma
func (p *Player) GetCharisma() int {
	return p.Charisma
}

func (p *Player) SetCharisma(value int) {
	p.Charisma = value
}

// Get and Set methods for Luck
func (p *Player) GetLuck() int {
	return p.Luck
}

func (p *Player) SetLuck(value int) {
	p.Luck = value
}

// Get and Set methods for CurrentScenario
func (p *Player) GetCurrentScenario() int {
	return p.CurrentScenario
}

func (p *Player) SetCurrentScenario(value int) {
	p.CurrentScenario = value
}

// Get and Set methods for EventHistory
func (p *Player) GetEventHistory() []int {
	return p.EventHistory
}

func (p *Player) SetEventHistory(value []int) {
	p.EventHistory = value
}