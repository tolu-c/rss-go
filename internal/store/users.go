package store

import (
	"context"

	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/store/queries"
)

// CreateUserParams is re-exported so callers don't need to import the queries
// subpackage. The store is the only place that talks to it.
type CreateUserParams = queries.CreateUserParams

func (s *Store) CreateUser(ctx context.Context, arg CreateUserParams) (model.User, error) {
	row, err := s.queries.CreateUser(ctx, arg)
	if err != nil {
		return model.User{}, err
	}
	return userFromRow(row), nil
}

func (s *Store) GetUserByApiKey(ctx context.Context, apiKey string) (model.User, error) {
	row, err := s.queries.GetUserByApiKey(ctx, apiKey)
	if err != nil {
		return model.User{}, err
	}
	return userFromRow(row), nil
}

// userFromRow translates a DB row into a domain entity. Lives at the layer
// boundary: nothing above the store knows queries.User exists.
func userFromRow(u queries.User) model.User {
	return model.User{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
	}
}
