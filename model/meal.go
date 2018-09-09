package model

type Meal struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Url string `json:"url,omitempty"`
}
