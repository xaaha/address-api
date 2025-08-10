package db_test

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3" // needed for sqlite3 driver
	"github.com/xaaha/address-api/internal/data"
	"github.com/xaaha/address-api/internal/db"
)

func createTempSQLFile(t *testing.T, content string) string {
	t.Helper()
	tmpFile := filepath.Join(t.TempDir(), "test.sql")
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("failed to write temp sql file: %v", err)
	}
	return tmpFile
}

func TestCreateDB(t *testing.T) {
	t.Run("creates db successfully", func(t *testing.T) {
		tempDir := t.TempDir()
		oldWd, _ := os.Getwd()
		defer os.Chdir(oldWd)
		os.Chdir(tempDir)

		if err := os.MkdirAll(filepath.Join("internal", "db"), 0o755); err != nil {
			t.Fatalf("failed to create folders: %v", err)
		}

		d, err := db.CreateDB()
		if err != nil {
			t.Fatalf("CreateDB() returned error: %v", err)
		}
		if d == nil {
			t.Fatal("CreateDB() returned nil db")
		}

		if _, err := d.Exec(`CREATE TABLE test (id INTEGER);`); err != nil {
			t.Fatalf("failed to create table: %v", err)
		}

		expectedPath := filepath.Join("internal", "db", "data.db")
		if _, err := os.Stat(expectedPath); err != nil {
			t.Errorf("expected db file at %s, got error: %v", expectedPath, err)
		}
	})
}

func TestExecSQLFile(t *testing.T) {
	t.Run("executes schema creation SQL", func(t *testing.T) {
		d, err := sql.Open("sqlite3", ":memory:")
		if err != nil {
			t.Fatalf("failed to open in-memory DB: %v", err)
		}
		schema := `CREATE TABLE addresses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			address TEXT,
			phone TEXT,
			country_code TEXT,
			country TEXT
		);`
		sqlFile := createTempSQLFile(t, schema)

		if err := db.ExecSQLFile(d, sqlFile); err != nil {
			t.Fatalf("ExecSQLFile() failed: %v", err)
		}
	})

	t.Run("returns error for missing SQL file", func(t *testing.T) {
		d, _ := sql.Open("sqlite3", ":memory:")
		err := db.ExecSQLFile(d, "nonexistent.sql")
		if err == nil {
			t.Fatal("ExecSQLFile() expected error, got nil")
		}
	})
}

func TestInsertAddress(t *testing.T) {
	t.Run("inserts one address successfully", func(t *testing.T) {
		d, _ := sql.Open("sqlite3", ":memory:")
		schema := `CREATE TABLE addresses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			address TEXT,
			phone TEXT,
			country_code TEXT,
			country TEXT
		);`
		if _, err := d.Exec(schema); err != nil {
			t.Fatalf("failed to create schema: %v", err)
		}

		insertSQL := `INSERT INTO addresses (name, address, phone, country_code, country) VALUES (?, ?, ?, ?, ?);`
		sqlFile := createTempSQLFile(t, insertSQL)

		addr := data.Address{
			Name:        "Glaner Hotel Café",
			Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
			Phone:       "+376879444",
			CountryCode: "AD",
			Country:     "Andorra",
		}

		if err := db.InsertAddress(d, addr, sqlFile); err != nil {
			t.Fatalf("InsertAddress() failed: %v", err)
		}

		var count int
		if err := d.QueryRow(`SELECT COUNT(*) FROM addresses`).Scan(&count); err != nil {
			t.Fatalf("query count failed: %v", err)
		}
		if count != 1 {
			t.Fatalf("expected 1 row, got %d", count)
		}
	})

	t.Run("returns error for bad SQL file path", func(t *testing.T) {
		d, _ := sql.Open("sqlite3", ":memory:")
		addr := data.Address{}
		err := db.InsertAddress(d, addr, "badpath.sql")
		if err == nil {
			t.Fatal("InsertAddress() expected error, got nil")
		}
	})
}

func TestInsertAddressesInBulk(t *testing.T) {
	t.Run("inserts multiple addresses successfully", func(t *testing.T) {
		d, _ := sql.Open("sqlite3", ":memory:")
		schema := `CREATE TABLE addresses (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			address TEXT,
			phone TEXT,
			country_code TEXT,
			country TEXT
		);`
		if _, err := d.Exec(schema); err != nil {
			t.Fatalf("failed to create schema: %v", err)
		}

		insertSQL := `INSERT INTO addresses (name, address, phone, country_code, country) VALUES (?, ?, ?, ?, ?);`
		sqlFile := createTempSQLFile(t, insertSQL)

		addresses := []data.Address{
			{
				Name:        "Glaner Hotel Café",
				Address:     "Carrer de na Maria Pla, 19-21, AD500 Andorra la Vella, Andorra",
				Phone:       "+376879444",
				CountryCode: "AD",
				Country:     "Andorra",
			},
			{
				Name:        "Hotel Magic",
				Address:     "Av. Doctor Mitjavila, 3, AD500 Andorra la Vella, Andorra",
				Phone:       "+376876900",
				CountryCode: "AD",
				Country:     "Andorra",
			},
		}

		if err := db.InsertAddressesInBulk(d, addresses, sqlFile); err != nil {
			t.Fatalf("InsertAddressesInBulk() failed: %v", err)
		}

		var count int
		if err := d.QueryRow(`SELECT COUNT(*) FROM addresses`).Scan(&count); err != nil {
			t.Fatalf("query count failed: %v", err)
		}
		if count != 2 {
			t.Fatalf("expected 2 rows, got %d", count)
		}
	})

	t.Run("returns error for bad SQL file path", func(t *testing.T) {
		d, _ := sql.Open("sqlite3", ":memory:")
		addresses := []data.Address{}
		err := db.InsertAddressesInBulk(d, addresses, "badpath.sql")
		if err == nil {
			t.Fatal("InsertAddressesInBulk() expected error, got nil")
		}
	})
}
