package data_test

import (
	"testing"

	"github.com/xaaha/address-api/internal/data"
)

func TestReadJSON(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dir     string
		want    []data.Address
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := data.ReadJSON(tt.dir)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ReadJSON() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ReadJSON() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ReadJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
