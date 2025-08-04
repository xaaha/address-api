package db_test

import (
	"database/sql"
	"testing"

	"github.com/xaaha/address-api/internal/db"
)

func TestCreateDB(t *testing.T) {
	tests := []struct {
		name    string // description of this test case
		want    *sql.DB
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := db.CreateDB()
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("CreateDB() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("CreateDB() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("CreateDB() = %v, want %v", got, tt.want)
			}
		})
	}
}
