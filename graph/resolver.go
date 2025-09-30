package graph

import (
	"github.com/xaaha/address-api/internal/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver servers as a dependency  injection for queries in graph/schema.resolvers.go
type Resolver struct {
	Repo repository.AddressRepositoryInterface
}
