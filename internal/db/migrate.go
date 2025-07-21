// Package db contains migration logic
package db

import (
	"fmt"
	"os"
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

	jsonData := string(jsonByte)
	fmt.Println(jsonData)

	return nil
}
