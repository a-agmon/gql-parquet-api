package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"

	"github.com/a-agmon/gql-parquet-api/graph/model"
)

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.db.GetUsers()
}

// GetUserByEmailDomain is the resolver for the getUserByEmailDomain field.
func (r *queryResolver) GetUserByEmailDomain(ctx context.Context, domain string) ([]*model.User, error) {
	return r.db.GetUsers()
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
