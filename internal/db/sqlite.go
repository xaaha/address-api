package db

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/xaaha/address-api/internal/data"
)

// Should expose Connect() and return *sql.DB

// CreateDB creates db named data
func CreateDB() (*sql.DB, error) {
	dbLocation := filepath.Join("internal", "db", "data.db")
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// ExecSQLFile takes in db and .sql file and executes the statement
func ExecSQLFile(db *sql.DB, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlBytes))

	return err
}

// InsertAddress inserts address to the db created in CreateDB
func InsertAddress(db *sql.DB, addr data.Address, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	sqlStmt := string(sqlBytes)
	_, err = db.Exec(
		sqlStmt,
		addr.Name,
		addr.Address,
		addr.Phone,
		addr.CountryCode,
		addr.Country,
	)
	return err
}
