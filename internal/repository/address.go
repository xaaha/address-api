// Package repository has the db access logic for graphql queries
package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/xaaha/address-api/graph/model"
)

// AddressRepositoryInterface defines the contract for our address repository
type AddressRepositoryInterface interface {
	GetCountryCode(
		ctx context.Context,
		country *string,
	) ([]*model.CountryInfo, error)

	GetAddressesByCountryCode(
		ctx context.Context,
		countryCode string,
		count *int32,
	) ([]*model.Address, error)
}

// AddressRepository struct for refactoring schema.resolvers.go
type AddressRepository struct {
	DB *sql.DB
}

// NewAddressRepository is the constructor function for AddressRepository
func NewAddressRepository(db *sql.DB) *AddressRepository { return &AddressRepository{DB: db} }

// GetCountryCode has the db logic for ContryCode resolver
func (r *AddressRepository) GetCountryCode(
	ctx context.Context,
	country *string,
) ([]*model.CountryInfo, error) {
	// ideally the db logic should be separated to perhaps internal/repository/address_repo.go
	// but for this small app, I am going to leave this here
	db := r.DB

	var query string
	var args []any

	const baseQuery = "SELECT DISTINCT country, country_code FROM address"

	if country != nil {
		query = baseQuery + " WHERE country = ? ORDER BY country"
		processedCountry := strings.TrimSpace(*country)
		args = append(args, processedCountry)
	} else {
		query = baseQuery + " ORDER BY country"
	}

	rows, err := db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("database query failed: %w", err)
	}
	defer rows.Close()

	var results []*model.CountryInfo

	for rows.Next() {
		var info model.CountryInfo

		if err := rows.Scan(&info.Country, &info.Code); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, &info)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	// If the user searched for a specific country AND we found no results...
	if country != nil && len(results) == 0 {
		// ...return a helpful error instead of just an empty list.
		errMsg := fmt.Sprintf(
			"No matching country found for '%s'. Omit the argument to get a list of all countries and codes.",
			*country,
		)
		return nil, errors.New(errMsg)
	}

	return results, nil
}

// GetAddressesByCountryCode has the db logic for ContryCode resolver
func (r *AddressRepository) GetAddressesByCountryCode(
	ctx context.Context,
	countryCode string,
	count *int32,
) ([]*model.Address, error) {
	db := r.DB

	// country code should be Uppercase.
	processedCode := strings.ToUpper(countryCode)
	const defaultLimit int32 = 5
	const maxLimit int32 = 50

	limit := defaultLimit
	if count != nil {
		limit = min(*count, maxLimit)
		if limit <= 0 {
			limit = defaultLimit
		}
	}

	query := `
        SELECT id, name, full_address, phone, country_code, country
        FROM address
        WHERE country_code = ?
        ORDER BY RANDOM()
        LIMIT ?`

	rows, err := db.QueryContext(ctx, query, processedCode, limit)
	if err != nil {
		return nil, fmt.Errorf("error on db query:  %v", err)
	}
	defer rows.Close()

	var results []*model.Address
	for rows.Next() {
		var addr model.Address

		err := rows.Scan(
			&addr.ID,
			&addr.Name,
			&addr.FullAddress,
			&addr.Phone,
			&addr.CountryCode,
			&addr.Country,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		results = append(results, &addr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}
	return results, nil
}
