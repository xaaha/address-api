package db

import (
	"database/sql"
	"os"
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

func ExecSQLFile(db *sql.DB, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlBytes))

	return err
}
