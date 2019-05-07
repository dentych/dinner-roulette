package models

import "time"

type Recipe struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Url         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	Directions  *string `json:"directions,omitempty"`
}

type User struct {
	ID               int
	Email            string
	PasswordHash     string
	FirstName        string
	LastName         string
	CreatedTimestamp time.Time
}

type Session struct {
	ID        int
	UserId    int
	SessionId string
	ValidTo   time.Time
}

type MealPlan struct {
	Recipes []Recipe
}