// Package feedfollows is the HTTP layer for feed-follow endpoints.
package feedfollows

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"

	apimodel "github.com/tolu-c/rss-go/internal/api/model"
	appff "github.com/tolu-c/rss-go/internal/app/feedfollows"
	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/pkg/middleware"
	"github.com/tolu-c/rss-go/internal/pkg/response"
)

type Handler struct {
	svc appff.FeedFollows
	mw  *middleware.Middleware
}

func New(svc appff.FeedFollows, mw *middleware.Middleware) *Handler {
	return &Handler{svc: svc, mw: mw}
}

func (h *Handler) Build(r chi.Router) {
	r.Post("/feed/follow", h.mw.Auth(h.create))
	r.Get("/feed/follows", h.mw.Auth(h.list))
	r.Delete("/feed/follow/{feedFollowID}", h.mw.Auth(h.delete))
}

func (h *Handler) create(w http.ResponseWriter, r *http.Request, user model.User) {
	var req apimodel.CreateFeedFollowRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Error parsing json: %s", err))
		return
	}

	follow, err := h.svc.Follow(r.Context(), user.ID, req.FeedID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to follow feed: %s", err))
		return
	}

	response.JSON(w, http.StatusOK, apimodel.NewFeedFollowResponse(follow))
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request, user model.User) {
	follows, err := h.svc.ListByUser(r.Context(), user.ID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to retrieve feed follows: %s", err))
		return
	}
	response.JSON(w, http.StatusOK, apimodel.NewFeedFollowResponses(follows))
}

func (h *Handler) delete(w http.ResponseWriter, r *http.Request, user model.User) {
	idStr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Invalid feed follow ID: %s", err))
		return
	}

	if err := h.svc.Delete(r.Context(), user.ID, feedFollowID); err != nil {
		response.Error(w, http.StatusBadRequest, fmt.Sprintf("Failed to delete feed follow: %s", err))
		return
	}

	response.JSON(w, http.StatusOK, struct{}{})
}
