package model

import (
	"time"

	"github.com/google/uuid"

	domain "github.com/tolu-c/rss-go/internal/model"
)

type CreateFeedRequest struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type FeedResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

func NewFeedResponse(f domain.Feed) FeedResponse {
	return FeedResponse{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Url:       f.Url,
		UserID:    f.UserID,
	}
}

func NewFeedResponses(feeds []domain.Feed) []FeedResponse {
	out := make([]FeedResponse, 0, len(feeds))
	for _, f := range feeds {
		out = append(out, NewFeedResponse(f))
	}
	return out
}
