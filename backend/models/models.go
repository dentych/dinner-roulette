package models

import "time"

type Recipe struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name"`
	Url  string `json:"url,omitempty"`
}

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

type Session struct {
	ID        int
	UserId    int
	SessionId string
	ValidTo   time.Time
}
