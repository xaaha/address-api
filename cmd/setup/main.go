// Package setup sets runs the sqlite db migration before
package main

import (
	"log"

	"github.com/xaaha/address-api/internal/db"
)

func main() {
	err := db.CreateDBAndTables("./data")
	if err != nil {
		log.Println(err)
	}
}
