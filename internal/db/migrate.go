// Package db contains migration logic
package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// Should handle creating tables and inserting data

type address struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country-code"`
	Country     string `json:"country"`
}

// ReadJSON reads json for now
func ReadJSON() error {
	file := "internal/db/Afghanistan.json"
	jsonByte, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// jsonString := string(jsonByte)
	// fmt.Println(jsonString)

	var testAddress []address
	err = json.Unmarshal(jsonByte, &testAddress)
	if err != nil {
		return err
	}

	for _, value := range testAddress {
		fmt.Println(value.Name)
		fmt.Printf("Inserted %d records from %s\n", len(testAddress), file)
	}

	return nil
}

// CreateDBAndTables creates sqlite tables and db
func CreateDBAndTables() error {
	dbLocation := filepath.Join("internal", "db", "data.db")
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return err
	}

	defer func() {
		if err = db.Close(); err != nil {
			fmt.Println("Error closing db: ", err)
		}
	}()

	sqlPath := filepath.Join("db", "migrations", "001_create_addresses_table.sql")
	sqlBytes, err := os.ReadFile(sqlPath)
	if err != nil {
		return err
	}

	sqlStmt := string(sqlBytes)
	result, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	fmt.Println("Result: ", result)

	return nil
}
