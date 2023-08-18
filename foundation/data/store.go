package data

import "github.com/a-agmon/gql-parquet-api/graph/model"

type Store struct{}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) GetUsers() ([]*model.User, error) {
	return []*model.User{
		{
			ID:        "1",
			Email:     "some@domain.com",
			FirstName: "Some",
			LastName:  "Name",
			AccountID: "1",
		},
		{
			ID:        "2",
			Email:     "some@domain2",
			FirstName: "Some2",
			LastName:  "Name2",
			AccountID: "2",
		},
	}, nil
}
