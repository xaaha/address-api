// Package data loads the json data to the sqlite db
package data

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
)

// Address represents the JSON object in address file
type Address struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phone       string `json:"phone"`
	CountryCode string `json:"country-code"`
	Country     string `json:"country"`
}

// ReadJSON reads all json address from the provided directory
func ReadJSON(dir string) ([]Address, error) {
	var all []Address

	err := filepath.WalkDir(dir, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil || dirEntry.IsDir() || filepath.Ext(path) != ".json" {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var address []Address
		if err := json.Unmarshal(data, &address); err != nil {
			return err
		}

		all = append(all, address...)
		return nil
	})

	return all, err
}
