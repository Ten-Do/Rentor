package storage

import (
	"database/sql"
	"fmt"
	"rentor/internal/logger"

	"github.com/pressly/goose/v3"
)

const (
	directionUp   = "up"
	directionDown = "down"
)

func RunMigrations(db *sql.DB, migrationsDir string, direction string) error {

	// Set the dialect for goose
	if err := goose.SetDialect("sqlite3"); err != nil {
		return fmt.Errorf("runMigrations: failed to set goose dialect: %w", err)
	}

	switch direction {
	case directionUp:
		if err := goose.Up(db, migrationsDir); err != nil {
			return fmt.Errorf("runMigrations: failed to apply up migrations: %w", err)
		}
		logger.Info("Database migrations applied successfully", logger.Field("direction", directionUp))
	case directionDown:
		if err := goose.Down(db, migrationsDir); err != nil {
			return fmt.Errorf("runMigrations: failed to apply down migrations: %w", err)
		}
		logger.Info("Database migrations rolled back successfully", logger.Field("direction", directionDown))
	default:
		return fmt.Errorf("invalid migration direction: %s", direction)
	}
	return nil
}
