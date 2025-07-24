// Package data loads the json data to the sqlite db
package data

import (
	"encoding/json"
	"fmt"
	"os"
)

type Address struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country-code"`
	Country     string `json:"country"`
}

func ReadJSON() error {
	// TODO: Loop over dir and rad all the files
	file := "data/Afghanistan.json"
	jsonByte, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	var testAddress []Address
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
