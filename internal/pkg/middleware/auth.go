// Package middleware holds HTTP middleware that needs construction-time
// dependencies (e.g. the store). Stateless middleware can be plain functions;
// stateful middleware lives on a Middleware struct so its dependencies are
// explicit.
package middleware

import (
	"fmt"
	"net/http"

	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/pkg/response"
	"github.com/tolu-c/rss-go/internal/pkg/security"
	"github.com/tolu-c/rss-go/internal/store"
)

// AuthedHandler is the signature for handlers that require an authenticated
// user. The middleware injects the resolved user as the third argument.
type AuthedHandler func(http.ResponseWriter, *http.Request, model.User)

// Middleware bundles dependencies needed by middleware that touches state.
// New middleware methods get added as needed without changing call sites.
type Middleware struct {
	store *store.Store
}

// New constructs a Middleware. The store is used by Auth to resolve API keys
// to users.
func New(s *store.Store) *Middleware {
	return &Middleware{store: s}
}

// Auth wraps an AuthedHandler so the underlying handler only runs when a
// valid API key resolves to an existing user. Failures short-circuit with
// 401 (missing/malformed key) or 400 (key not found).
func (m *Middleware) Auth(handler AuthedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := security.GetApiKey(r.Header)
		if err != nil {
			response.Error(w, http.StatusUnauthorized, fmt.Sprintf("Failed to get api key: %s", err))
			return
		}

		user, err := m.store.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
