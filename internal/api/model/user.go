// Package model holds the API-layer DTOs — the JSON wire shapes for
// requests and responses. These are deliberately separate from the domain
// entities in internal/model so the HTTP contract can evolve independently
// of the internal data model.
package model

import (
	"time"

	"github.com/google/uuid"

	domain "github.com/tolu-c/rss-go/internal/model"
)

// CreateUserRequest is the body shape for POST /v1/user.
type CreateUserRequest struct {
	Name string `json:"name"`
}

// UserResponse is the JSON shape returned to API callers.
type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

// NewUserResponse converts a domain User into its wire representation.
// The conversion is intentionally explicit so adding/removing/renaming
// API fields cannot accidentally leak from changes to the domain type.
func NewUserResponse(u domain.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Name:      u.Name,
		ApiKey:    u.ApiKey,
	}
}
