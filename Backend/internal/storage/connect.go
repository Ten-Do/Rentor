package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: keep-alive, connection pool, etc.

func Connect(storage_path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", storage_path)
	if err != nil {
		return nil, fmt.Errorf("Connect: failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Connect: failed to ping database: %w", err)
	}

	return db, nil
}
