package db

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/xaaha/address-api/internal/data"
)

// Should expose Connect() and return *sql.DB

func CreateDB() (*sql.DB, error) {
	dbLocation := filepath.Join("internal", "db", "data.db")
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func ExecSQLFile(db *sql.DB, path string) error {
	sqlBytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlBytes))

	return err
}

func InsertAddress(db *sql.DB, addr data.Address) error {
	sqlStmt := `
 INSERT INTO address(id, name, address, phone, country_code, country)
 VALUES ()
 `
	_, err := db.Exec(
		sqlStmt,
		addr.ID,
		addr.Name,
		addr.Address,
		addr.Phone,
		addr.CountryCode,
		addr.Country,
	)
	return err
}
