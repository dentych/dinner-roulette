package database

import (
	"github.com/jmoiron/sqlx"
)

var ConnectionString = "user=postgres user=postgres dbname=dinner-dash password=Ixtj2AokxSdGfbwSRhJorpFkFMxE3Ihy host=35.228.143.86"

func GetConnection() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", ConnectionString)
	if err != nil {
		return nil, err
	}

	return db, nil
}