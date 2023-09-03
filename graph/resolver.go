package graph

import (
	"github.com/a-agmon/gql-parquet-api/pkg/data"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *data.DAO
}

func NewResolver(db *data.DAO) *Resolver {
	return &Resolver{db: db}
}
