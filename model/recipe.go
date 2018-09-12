package model

type Recipe struct {
	ID int `json:"id,omitempty"`
	Name string `json:"name"`
	Url string `json:"url,omitempty"`
}
