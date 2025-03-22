package models

type Model interface {
	GetId() int
	SetId(id int)
}

type modelImpl struct {
	Id int `json:"id"`
}

func (m *modelImpl) GetId() int {
	return m.Id
}

func (m *modelImpl) SetId(id int) {
	m.Id = id
}