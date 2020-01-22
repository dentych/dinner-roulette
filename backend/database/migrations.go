package database

import (
	"fmt"
	"github.com/dentych/dinner-dash/config"
	"github.com/dentych/dinner-dash/logging"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(config config.DatabaseConfig) {
	var m *migrate.Migrate
	var err error

	connectionString := createConnectionString(config)

	if m, err = migrate.New("file://./migrations", connectionString); err != nil {
		logging.Error.Fatalf("Failed to migrate: %v", err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange{
		logging.Error.Fatalf("Failed to run migrate UP: %v", err)
	}
}

func createConnectionString(config config.DatabaseConfig) string {
	format := "postgres://%s:%s@%s:5432/%s?sslmode=disable"
	return fmt.Sprintf(format, config.Username, config.Password, config.Hostname, config.Database)
}
