// Package db contains migration logic
package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

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
	jsonByte, err := os.ReadFile("internal/db/Afghanistan.json")
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
	}

	return nil
}

// CreateDBAndTables creates sqlite tables and db
func CreateDBAndTables() error {
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		return err
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS address (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        name TEXT
				address TEXT
				phone TEXT
				country_code TEXT
				country TEXT
    );
	`
	result, err := db.Exec(sqlStmt)
	if err != nil {
		return err
	}
	fmt.Println("Result: ", result)

	return nil
}
