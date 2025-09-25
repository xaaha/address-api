package graph

import (
	"context"
	"database/sql"
	"os"
	"reflect"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/xaaha/address-api/graph/model"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()

	testDb, err := sql.Open("sqlite3", ":memory:?_foreign_keys=on")
	if err != nil {
		t.Fatalf("failed to open in memory db: %v", err)
	}

	t.Cleanup(func() { testDb.Close() })

	createAddrSQLFile := "../db/migrations/001_create_addresses_table.sql"
	migration, err := os.ReadFile(createAddrSQLFile)
	if err != nil {
		t.Fatalf("failed to read migration file %v", err)
	}

	if _, err := testDb.Exec(string(migration)); err != nil {
		t.Fatalf("failed to execute migration: %v", err)
	}

	return testDb
}

func Test_queryResolver_CountryCode(t *testing.T) {
	tests := []struct {
		name          string
		countryToFind *string
		seedData      []model.CountryInfo
		want          []*model.CountryInfo
		wantErr       bool
		errContains   string
	}{
		{
			name:          "Get all unique countries when input is nil",
			countryToFind: nil,
			seedData: []model.CountryInfo{
				{Country: "United States", Code: "US"},
				{Country: "Canada", Code: "CA"},
				{Country: "United States", Code: "US"},
			},
			want: []*model.CountryInfo{
				{Country: "United States", Code: "US"},
				{Country: "Canada", Code: "CA"},
			},
			wantErr: false,
		},
		{
			name:          "Find specific country that exists",
			countryToFind: func() *string { s := "Canada"; return &s }(),
			seedData: []model.CountryInfo{
				{Country: "United States", Code: "US"},
				{Country: "Canada", Code: "CA"},
			},
			want: []*model.CountryInfo{
				{Country: "Canada", Code: "CA"},
			},
			wantErr: false,
		},
		{
			name:          "Return a helpful error for a country that does not exist",
			countryToFind: func() *string { s := "Mexico"; return &s }(),
			seedData: []model.CountryInfo{
				{Country: "United States", Code: "US"},
				{Country: "Canada", Code: "CA"},
			},
			want:        nil,
			wantErr:     true,
			errContains: "No matching country found for 'Mexico'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testDb := newTestDB(t)

			for _, data := range tt.seedData {
				_, err := testDb.Exec(
					"INSERT INTO address (country, country_code) VALUES (?, ?)",
					data.Country,
					data.Code,
				)
				if err != nil {
					t.Fatalf("failed to seed the data %v", err)
				}
			}

			resolver := &queryResolver{
				Resolver: &Resolver{DB: testDb},
			}

			got, gotErr := resolver.CountryCode(context.Background(), tt.countryToFind)
			if (gotErr != nil) != tt.wantErr {
				t.Fatalf("CountryCode() error = %v, wantErr %v", gotErr, tt.wantErr)
			}
			if tt.wantErr {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountryCode() got = %v, want %v", got, tt.want)
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
