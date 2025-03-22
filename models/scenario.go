package models

type Scenario struct {
	modelImpl
	Name     string    `json:"name"`
}

func (sc *Scenario) GetName() string {
	return sc.Name
}