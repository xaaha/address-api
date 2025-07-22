// Package main is the entry point
package main

import (
	"log"

	"github.com/xaaha/address-api/internal/db"
)

func main() {
	// err := db.ReadJSON()
	// if err != nil {
	// 	log.Println(err)
	// }

	err := db.CreateDBAndTables()
	if err != nil {
		log.Println(err)
	}
}
