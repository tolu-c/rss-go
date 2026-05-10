// Package users is the service layer for user-related operations.
// The exported Users interface is what api handlers depend on; the
// unexported service struct is the concrete implementation.
package users

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/tolu-c/rss-go/internal/app"
	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/store"
)

// Users defines the user-facing operations. Interface lives here (not in
// the api layer) so any caller — HTTP, CLI, background job — can satisfy
// its dependency on user logic without knowing about the implementation.
type Users interface {
	Create(ctx context.Context, name string) (model.User, error)
}

type service struct {
	dp app.Dependency
}

// New returns a Users implementation wired to dp. Returning the interface
// (not *service) prevents callers from bypassing the constructor.
func New(dp app.Dependency) Users {
	return &service{dp: dp}
}

// Create assigns an ID and timestamps to a new user and persists it.
// The store handles api_key generation via SQL (sha256 of randomness).
func (s *service) Create(ctx context.Context, name string) (model.User, error) {
	now := time.Now().UTC()
	return s.dp.Store.CreateUser(ctx, store.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	})
}
