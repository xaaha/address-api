// Package setup sets runs the sqlite db migration before
package setup

import (
	"log"

	"github.com/xaaha/address-api/internal/db"
)

func main() {
	err := db.CreateDBAndTables()
	if err != nil {
		log.Println(err)
	}
}
