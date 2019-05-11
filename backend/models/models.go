package models

import "time"

type Recipe struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name"`
	Url         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
	Directions  *string `json:"directions,omitempty"`
}

type Ingredient struct {
	ID 					int `json:"id,omitempty"`
	Name        string `json:"name"`
	KCAL				int  `json:"kcal,omitempty"`
}


type RecipeIngredient struct {
	ID					int
	Amount			int
	Unit				string
	Name        string
	KCAL				int
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
