package data

import (
	"encoding/json"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var usStates = [...]string{
	"AL", "AK", "AZ", "AR", "CA", "CO", "CT", "DE", "FL", "GA",
	"HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD",
	"MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ",
	"NM", "NY", "NC", "ND", "OH", "OK", "OR", "PA", "RI", "SC",
	"SD", "TN", "TX", "UT", "VT", "VA", "WA", "WV", "WI", "WY",
}

// loop through all the json files in the `data/` directory
// then, if the file name != "United States"
// and the address section has address matching the united states
// remove the address from the array

func usAddress(addr Address) bool {
	fullAddress := strings.ToLower(addr.Address)

	if strings.Contains(fullAddress, "usa") || strings.Contains(fullAddress, "united states") {
		return true
	}

	for _, state := range usStates {
		if strings.Contains(fullAddress, " "+strings.ToLower(state)+" ") {
			return true
		}
	}

	zipPattern := regexp.MustCompile(`\b\d{5}(-\d{4})?\b`)
	return zipPattern.MatchString(fullAddress)
}

// rmJunkItem removes the polluted items from the arrAddr
func rmJunkItem(arrAddr []Address) []Address {
	cleaned := arrAddr[:0]
	for _, val := range arrAddr {
		if !usAddress(val) {
			cleaned = append(cleaned, val)
		}
	}
	return cleaned
}

// Cleanup reads the data dir
func Cleanup(dirPath string) error {
	return filepath.WalkDir(dirPath, func(path string, dirEntry fs.DirEntry, err error) error {
		if err != nil || dirEntry.IsDir() || filepath.Ext(path) != ".json" {
			return err
		}

		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		var address []Address
		if err = json.Unmarshal(data, &address); err != nil {
			return err
		}

		cleanedupArr := rmJunkItem(address)

		newData, err := json.MarshalIndent(cleanedupArr, "", "  ")
		if err != nil {
			return err
		}

		if err := os.WriteFile(path, newData, 0644); err != nil {
			return err
		}

		return nil
	})
}

// read the dir
// what if the dir is not there
// for each file in the dir, loop (user walk dir)
// read the file,
// what if the file is not there
// then using the json struct read the file array
// for each file, if the condition matches, remove the item
// once done move to the next file content.
