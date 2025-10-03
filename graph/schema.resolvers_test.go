package graph

import (
	"context"
	"errors"
	"testing"

	"github.com/xaaha/address-api/graph/model"
)

// for each query and/or mutation, check that when the

type mockAddressRepository struct {
	mockGetCountryCode            func(ctx context.Context, country *string) ([]*model.CountryInfo, error)
	mockGetAddressesByCountryCode func(ctx context.Context, countryCode string, count *int32) ([]*model.Address, error)
}

func (m *mockAddressRepository) GetCountryCode(
	ctx context.Context,
	country *string,
) ([]*model.CountryInfo, error) {
	return m.mockGetCountryCode(ctx, country)
}

func (m *mockAddressRepository) GetAddressesByCountryCode(
	ctx context.Context,
	countrycode string,
	count *int32,
) ([]*model.Address, error) {
	return m.mockGetAddressesByCountryCode(ctx, countrycode, count)
}

func Test_queryResolver_CountryCode(t *testing.T) {
	t.Run("Successfully Returns the country code from the repo", func(t *testing.T) {
		mockRepo := &mockAddressRepository{
			mockGetCountryCode: func(_ context.Context, _ *string) ([]*model.CountryInfo, error) {
				return []*model.CountryInfo{{Country: "Testnation", Code: "TN"}}, nil
			},
		}

		resolver := &queryResolver{
			Resolver: &Resolver{Repo: mockRepo},
		}

		got, err := resolver.CountryCode(context.Background(), nil)
		if err != nil {
			t.Fatalf("issue occured on calling resolver.CountryCode: %v", err)
		}
		if got[0].Country != "Testnation" {
			t.Fatalf("expected 'Testnation' got  %v", got[0].Country)
		}
	})

	t.Run("Correctly returns an error from the repository", func(t *testing.T) {
		// Setup the mock to return an error.
		mockRepo := &mockAddressRepository{
			mockGetCountryCode: func(_ context.Context, _ *string) ([]*model.CountryInfo, error) {
				return nil, errors.New("a simulated database error")
			},
		}

		resolver := &queryResolver{
			Resolver: &Resolver{Repo: mockRepo},
		}

		_, err := resolver.CountryCode(context.Background(), nil)
		if err != nil {
			if err.Error() != "a simulated database error" {
				t.Fatalf(
					"expected error to be %v, but got %v",
					"a simulated database error",
					err.Error(),
				)
			}
		}
	})
}

func Test_queryResolver_AddressesByCountryCode(t *testing.T) {
	int32Ptr := func(i int32) *int32 { return &i }

	t.Run("Successfully returns addresses from the repository", func(t *testing.T) {
		mockRepo := &mockAddressRepository{
			mockGetAddressesByCountryCode: func(_ context.Context, countryCode string, _ *int32) ([]*model.Address, error) {
				if countryCode != "US" {
					return nil, errors.New("test setup error: expected countryCode 'US'")
				}
				return []*model.Address{
					{
						ID:          "123",
						Name:        "Test Address",
						CountryCode: func(str string) *string { return &str }("US"),
					},
				}, nil
			},
		}

		resolver := &queryResolver{
			Resolver: &Resolver{Repo: mockRepo},
		}

		got, err := resolver.AddressesByCountryCode(context.Background(), "US", int32Ptr(5))
		if err != nil {
			t.Fatalf("AddressesByCountryCode returned an unexpected error: %v", err)
		}
		if len(got) != 1 {
			t.Fatalf("expected 1 address, but got %d", len(got))
		}
		if got[0].ID != "123" {
			t.Fatalf("expected address ID to be '123', but got '%s'", got[0].ID)
		}
	})

	t.Run("Correctly returns an error from the repository", func(t *testing.T) {
		mockRepo := &mockAddressRepository{
			mockGetAddressesByCountryCode: func(_ context.Context, _ string, _ *int32) ([]*model.Address, error) {
				return nil, errors.New("a simulated database error")
			},
		}

		resolver := &queryResolver{
			Resolver: &Resolver{Repo: mockRepo},
		}

		_, err := resolver.AddressesByCountryCode(context.Background(), "CA", int32Ptr(10))

		if err == nil {
			t.Fatal("expected an error, but got nil")
		}
		expectedErr := "a simulated database error"
		if err.Error() != expectedErr {
			t.Fatalf("expected error message '%s', but got '%s'", expectedErr, err.Error())
		}
	})
}
