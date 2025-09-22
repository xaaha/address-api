package graph

import (
	"context"
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xaaha/address-api/graph/model"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite3", ":memory:?_foreign_keys=on")
	if err != nil {
		t.Fatalf("failed to open in memory db: %v", err)
	}

	t.Cleanup(func() { db.Close() })

	// TODO: Put this in a utility func
	createAddrSQLFile := filepath.Join("db", "migrations", "001_create_addresses_table.sql")
	migration, err := os.ReadFile(createAddrSQLFile)
	if err != nil {
		t.Fatalf("failed to read migration file %v", err)
	}

	// what's the difference between this and CreateDBAndTables in migrate.go?
	// can't we jus use that?
	if _, err := db.Exec(string(migration)); err != nil {
		t.Fatalf("failed to execute migration: %v", err)
	}

	return db
}

func Test_queryResolver_CountryCode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		country *string
		want    []*model.CountryInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var r queryResolver
			got, gotErr := r.CountryCode(context.Background(), tt.country)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CountryCode() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CountryCode() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_queryResolver_AddressesByCountryCode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		countryCode string
		count       *int32
		want        []*model.Address
		wantErr     bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: construct the receiver type.
			var r queryResolver
			got, gotErr := r.AddressesByCountryCode(context.Background(), tt.countryCode, tt.count)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddressesByCountryCode() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddressesByCountryCode() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("AddressesByCountryCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
