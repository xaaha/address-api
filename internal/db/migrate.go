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
// It does that by creating db and reading, executing sql statement to create table
// and then inserting addresses to the created tables from the dirPath (jsonFiles)
func CreateDBAndTables(dirPath string) error {
	db, err := CreateDB()
	if err != nil {
		return err
	}
	defer func() {
		if err = db.Close(); err != nil {
			fmt.Println("Error closing db: ", err)
		}
	}()

	createAddrSQLFile := filepath.Join("db", "migrations", "001_create_addresses_table.sql")
	if err = ExecSQLFile(db, createAddrSQLFile); err != nil {
		return err
	}

	addresses, err := data.ReadJSON(dirPath)
	if err != nil {
		return err
	}

	insertAddrSQLFile := filepath.Join("db", "migrations", "002_insert_address.sql")

	for _, addr := range addresses {
		if err = InsertAddress(db, addr, insertAddrSQLFile); err != nil {
			return err
		}
	}

	fmt.Printf("Inserted %d addresses\n", len(addresses))

	return nil
}
