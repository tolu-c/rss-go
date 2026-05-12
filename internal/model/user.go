// Package model holds the domain entities of the application — the "true"
// shapes that flow through the service layer. They are deliberately free of
// HTTP and database concerns: no JSON tags, no SQL tags. Conversion to/from
// DB rows happens in the store layer; conversion to/from API DTOs happens
// in the api layer.
package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	ApiKey    string
}
