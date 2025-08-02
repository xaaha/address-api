// Package db contains migration logic
package db

import (
	"fmt"
	"path/filepath"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Should handle creating tables and inserting data

// CreateDBAndTables creates sqlite tables and db
// TODO: Update this to read the entire dir
func CreateDBAndTables() error {
	db, err := CreateDB()
	if err != nil {
		return err
	}
	defer func() {
		if err = db.Close(); err != nil {
			fmt.Println("Error closing db: ", err)
		}
	}()

	sqlPath := filepath.Join("db", "migrations", "001_create_addresses_table.sql")
	if err = ExecSQLFile(db, sqlPath); err != nil {
		return err
	}

	return nil
}
