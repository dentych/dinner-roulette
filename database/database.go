package database

import (
	"github.com/dentych/dinner-dash/logging"
	"github.com/jmoiron/sqlx"
	"os"
)

var ConnectionString = " host=localhost user=postgres user=postgres dbname=dinner-dash password=" + os.Getenv("PQ_PASSWORD") +
	"  sslmode=disable"
var db *sqlx.DB

// Init will setup a new database connection. The method will panic
// if a database connection can not be established.
func Init() {
	db = sqlx.MustConnect("postgres", ConnectionString)
}

// getConnection will return a database object, which can be used to perform queries.
func getConnection() (*sqlx.DB, error) {
	if db == nil {
		logging.Error.Fatal("Database is nil. You must initialize before getting the connection.")
	}

	return db, nil
}
