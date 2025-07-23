package db

import (
	"database/sql"
	"path/filepath"
)

func CreateDB() (*sql.DB, error) {
	dbLocation := filepath.Join("internal", "db", "data.db")
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return nil, err
	}
	return db, nil
}
