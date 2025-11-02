package sqlite

import (
	"database/sql"
	"fmt"
)

type SQLiteStorage struct {
	db *sql.DB
}

// NewSQLiteStorage creates a new SQLiteStorage instance
func NewSQLiteStorage(storage_path string) (*SQLiteStorage, error) {
	db, err := sql.Open("sqlite3", storage_path)
	if err != nil {
		return nil, fmt.Errorf("NewSQLiteStorage: failed to open database: %w", err)
	}

	return &SQLiteStorage{db: db}, nil
}
