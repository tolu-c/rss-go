// Package app is the service layer — pure business logic with no HTTP or
// transport concerns. Each feature lives in its own subpackage (app/users,
// app/feeds, …) and exports an interface that the api layer depends on.
//
// The Dependency struct is the canonical DI container. Every service
// constructor takes one of these. New shared dependencies (loggers, caches,
// external API clients) get added as fields here and become available to
// every service without changing constructors.
package app

import (
	"github.com/tolu-c/rss-go/internal/pkg/environment"
	"github.com/tolu-c/rss-go/internal/store"
)

// Dependency bundles everything a service might need. Keep it small until
// it needs to grow — premature fields are clutter.
type Dependency struct {
	Env   *environment.Config
	Store *store.Store
}

// InitDp builds a Dependency. It exists as a named function (rather than a
// struct literal at the call site) so future construction logic — wiring a
// logger, opening a connection pool, validating env — has a single home.
func InitDp(env *environment.Config, s *store.Store) Dependency {
	return Dependency{
		Env:   env,
		Store: s,
	}
}
