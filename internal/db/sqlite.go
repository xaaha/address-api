package db

import (
	"database/sql"
	"path/filepath"
)

// Should expose Connect() and return *sql.DB

func CreateDB() (*sql.DB, error) {
	dbLocation := filepath.Join("internal", "db", "data.db")
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return nil, err
	}
	return db, nil
}
