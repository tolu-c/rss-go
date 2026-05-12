// Package feeds is the HTTP layer for feed-related endpoints.
package feeds

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	apimodel "github.com/tolu-c/rss-go/internal/api/model"
	appfeeds "github.com/tolu-c/rss-go/internal/app/feeds"
	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/pkg/middleware"
	"github.com/tolu-c/rss-go/internal/pkg/response"
)

type Handler struct {
	svc appfeeds.Feeds
	mw  *middleware.Middleware
}

func New(svc appfeeds.Feeds, mw *middleware.Middleware) *Handler {
	return &Handler{svc: svc, mw: mw}
}

func (h *Handler) Build(r chi.Router) {
	r.Post("/feed", h.mw.Auth(h.create))
	r.Get("/feeds", h.list)
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request, user model.User) {
	var req apimodel.CreateFeedRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	feed, err := h.svc.Create(r.Context(), req.Name, req.Url, user.ID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to create feed: %s", err))
		return
	}

	response.JSON(w, http.StatusCreated, apimodel.NewFeedResponse(feed))
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	feeds, err := h.svc.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to retrieve feeds: %s", err))
		return
	}
	response.JSON(w, http.StatusOK, apimodel.NewFeedResponses(feeds))
}
