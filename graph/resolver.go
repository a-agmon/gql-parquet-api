package graph

import (
	"github.com/a-agmon/gql-parquet-api/foundation/data"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *data.Store
}

func NewResolver(db *data.Store) *Resolver {
	return &Resolver{db: db}
}
