// Package main in Server is for server work like db migration and stuff
package main

import (
	"fmt"

	"github.com/xaaha/address-api/internal/db"
)

func main() {
	// fmt.Println("GO Girls!")
	err := db.CreateDBAndTables("data")
	if err != nil {
		fmt.Println("Error occurred in server/main.go ", err)
	}
}
