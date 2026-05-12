package model

import (
	"time"

	"github.com/google/uuid"

	domain "github.com/tolu-c/rss-go/internal/model"
)

type CreateFeedFollowRequest struct {
	FeedID uuid.UUID `json:"feedId"`
}

type FeedFollowResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	UserID    uuid.UUID `json:"userId"`
	FeedID    uuid.UUID `json:"feedId"`
}

func NewFeedFollowResponse(ff domain.FeedFollow) FeedFollowResponse {
	return FeedFollowResponse{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}

func NewFeedFollowResponses(ffs []domain.FeedFollow) []FeedFollowResponse {
	out := make([]FeedFollowResponse, 0, len(ffs))
	for _, ff := range ffs {
		out = append(out, NewFeedFollowResponse(ff))
	}
	return out
}
