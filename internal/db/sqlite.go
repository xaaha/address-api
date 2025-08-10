package db

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/xaaha/address-api/internal/data"
)

// CreateDB creates db named data
func CreateDB(dbName string) (*sql.DB, error) {
	dbLocation := filepath.Join("internal", "db", dbName)
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

// InsertAddressesInBulk inserts a slice of addresses into the database within a single transaction
func InsertAddressesInBulk(db *sql.DB, addresses []data.Address, path string) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// Defer a rollback in case anything fails.
	defer tx.Rollback()

	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	sqlStmt := string(sqlBytes)
	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, addr := range addresses {
		_, err = stmt.Exec(
			addr.Name,
			addr.Address,
			addr.Phone,
			addr.CountryCode,
			addr.Country,
		)
		if err != nil {
			return err
		}
	}
	return tx.Commit()
}
