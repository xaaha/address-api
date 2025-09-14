package graph_test

import (
	"context"
	"testing"

	"github.com/xaaha/address-api/graph/model"
)

func Test_queryResolver_CountryCode(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		country *string
		want    []string
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
