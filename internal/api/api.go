// Package api is the HTTP composition root. It owns the chi router and
// delegates feature routing to subpackages via their Build() methods.
package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"github.com/tolu-c/rss-go/internal/api/feedfollows"
	"github.com/tolu-c/rss-go/internal/api/feeds"
	"github.com/tolu-c/rss-go/internal/api/users"
	"github.com/tolu-c/rss-go/internal/app"
	appfeedfollows "github.com/tolu-c/rss-go/internal/app/feedfollows"
	appfeeds "github.com/tolu-c/rss-go/internal/app/feeds"
	appusers "github.com/tolu-c/rss-go/internal/app/users"
	"github.com/tolu-c/rss-go/internal/pkg/middleware"
	"github.com/tolu-c/rss-go/internal/pkg/response"
)

// Handler aggregates all per-feature handlers and the shared router.
// New() wires everything; Build() registers routes; Router() returns the
// configured chi router for the http.Server to use.
type Handler struct {
	router      *chi.Mux
	users       *users.Handler
	feeds       *feeds.Handler
	feedfollows *feedfollows.Handler
}

// New wires the api layer. It builds each feature's service, the
// middleware bundle, and the per-feature handler — all from a single
// Dependency. Adding a new feature is a four-line change here.
func New(dp app.Dependency) *Handler {
	mw := middleware.New(dp.Store)

	return &Handler{
		router:      newRouter(),
		users:       users.New(appusers.New(dp), mw),
		feeds:       feeds.New(appfeeds.New(dp), mw),
		feedfollows: feedfollows.New(appfeedfollows.New(dp), mw),
	}
}

// Build mounts every feature's routes under /v1 and the top-level utility
// routes (health, err) at the same level. Each feature's Build receives the
// /v1 router and is responsible for its own subroute prefix.
func (h *Handler) Build() {
	h.router.Route("/v1", func(r chi.Router) {
		r.Get("/health", handlerReadiness)
		r.Get("/err", handlerError)

		h.users.Build(r)
		h.feeds.Build(r)
		h.feedfollows.Build(r)
	})
}

// Router returns the underlying chi router. http.Server uses it directly.
func (h *Handler) Router() http.Handler {
	return h.router
}

// newRouter builds the chi router with the standard middleware stack.
func newRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	return r
}

// handlerReadiness is the trivial liveness probe. Lives here because it has
// no business value — not worth a feature folder.
func handlerReadiness(w http.ResponseWriter, _ *http.Request) {
	response.JSON(w, http.StatusOK, struct{}{})
}

// handlerError is a deliberate 400 used to verify error plumbing in dev.
func handlerError(w http.ResponseWriter, _ *http.Request) {
	response.Error(w, http.StatusBadRequest, "Something went wrong")
}
