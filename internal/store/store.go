// Package store is the data layer. It wraps the sqlc-generated queries package
// and exposes a single Store struct with methods per entity. Higher layers
// depend on *Store, never on the underlying queries package.
package store

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/tolu-c/rss-go/internal/store/queries"
)

// Store is the data-access entry point. It owns the raw DB connection so it
// can manage lifecycle (ping/close), and the sqlc-generated Queries struct
// for actual query execution.
type Store struct {
	db      *sql.DB
	queries *queries.Queries
}

// New opens a Postgres connection at dbURL, verifies it with a ping, and
// returns a Store ready for use. The caller is responsible for calling Close.
func New(dbURL string) (*Store, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("store: open: %w", err)
	}
	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("store: ping: %w", err)
	}
	return &Store{
		db:      db,
		queries: queries.New(db),
	}, nil
}

// Close releases the underlying DB connection.
func (s *Store) Close() error {
	return s.db.Close()
}
