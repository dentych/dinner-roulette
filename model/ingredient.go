package model

type Ingredient struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Amount int    `json:"amount,omitempty"`
	Unit   string `json:"unit,omitempty"`
}
