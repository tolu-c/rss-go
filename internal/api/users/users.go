// Package users is the HTTP layer for user-related endpoints. The Handler
// translates between JSON wire shapes and the users service; it contains
// no business logic.
package users

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	appusers "github.com/tolu-c/rss-go/internal/app/users"
	apimodel "github.com/tolu-c/rss-go/internal/api/model"
	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/pkg/middleware"
	"github.com/tolu-c/rss-go/internal/pkg/response"
)

// Handler bundles the dependencies a user-feature endpoint needs.
type Handler struct {
	svc appusers.Users
	mw  *middleware.Middleware
}

func New(svc appusers.Users, mw *middleware.Middleware) *Handler {
	return &Handler{svc: svc, mw: mw}
}

// Build registers user routes onto r. The feature owns its full URL space
// so the composition root just calls Build(r) and is done.
func (h *Handler) Build(r chi.Router) {
	r.Route("/user", func(r chi.Router) {
		r.Post("/", h.create)
		r.Get("/", h.mw.Auth(h.getByApiKey))
	})
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request) {
	var req apimodel.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	user, err := h.svc.Create(r.Context(), req.Name)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to create user: %s", err))
		return
	}

	response.JSON(w, http.StatusCreated, apimodel.NewUserResponse(user))
}

// getByApiKey is auth-protected — the middleware injects the resolved user
// directly. The handler just shapes the response.
func (h *Handler) getByApiKey(w http.ResponseWriter, _ *http.Request, user model.User) {
	response.JSON(w, http.StatusOK, apimodel.NewUserResponse(user))
}
