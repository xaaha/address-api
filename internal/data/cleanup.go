package data

import (
	"io/fs"
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

func readDir(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil || dirEntry.IsDir() || filepath.Ext(path) != ".json" {
			return err
		}
	})
}
