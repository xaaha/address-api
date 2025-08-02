// Package db contains migration logic
package db

import (
	"fmt"
	"path/filepath"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
	"github.com/xaaha/address-api/internal/data"
)

// Should handle creating tables and inserting data

// CreateDBAndTables creates sqlite tables and db
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

	addresses, err := data.ReadJSON("data")
	if err != nil {
		return err
	}

	for _, addr := range addresses {
		if err = InsertAddress(db, addr); err != nil {
			return err
		}
	}

	fmt.Printf("Inserted %d addresses\n", len(addresses))

	return nil
}
