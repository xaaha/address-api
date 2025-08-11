package db_test

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/xaaha/address-api/internal/db"
)

// TODO: Complete this.
// we might need to either complete or delete the interface migrate.go
func TestCreateDBAndTables(t *testing.T) {
	currWd, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get working dir: %v", err)
	}
	defer os.Chdir(currWd)

	tempDir := t.TempDir()
	if err = os.Chdir(tempDir); err != nil {
		t.Fatalf("failed to chdir to temp dir: %v", err)
	}

	// Prepare migrations and JSON inside tempDir/db/migrations
	migrationsDir := filepath.Join(tempDir, "db", "migrations")
	os.MkdirAll(migrationsDir, 0755)

	os.WriteFile(filepath.Join(migrationsDir, "001_create_addresses_table.sql"), []byte(`
        CREATE TABLE addresses (
            id INTEGER PRIMARY KEY,
            name TEXT,
            address TEXT,
            phone TEXT,
            country_code TEXT,
            country TEXT
        );
    `), 0644)

	os.WriteFile(filepath.Join(migrationsDir, "002_insert_address.sql"), []byte(`
        INSERT INTO addresses (name, address, phone, country_code, country)
        VALUES (?, ?, ?, ?, ?);
    `), 0644)

	// Sample JSON
	sampleJSON := `[{
        "ID": 1,
        "Name": "Test Hotel",
        "Address": "123 Street",
        "Phone": "+123456789",
        "CountryCode": "AD",
        "Country": "Andorra"
    }]`
	os.WriteFile(filepath.Join(tempDir, "addresses.json"), []byte(sampleJSON), 0644)

	fmt.Println("This is temp dir: ", tempDir)
	if err := db.CreateDBAndTables(tempDir); err != nil {
		t.Fatalf("CreateDBAndTables() failed: %v", err)
	}

	// Assert DB has data
	sqlDB, _ := sql.Open("sqlite3", filepath.Join("internal", "db", "data.db"))
	defer sqlDB.Close()

	var count int
	if err := sqlDB.QueryRow("SELECT COUNT(*) FROM addresses").Scan(&count); err != nil {
		t.Fatalf("query failed: %v", err)
	}
	if count != 1 {
		t.Errorf("expected 1 address, got %d", count)
	}
}
