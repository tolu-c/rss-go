// Package feeds is the service layer for feed-related operations.
package feeds

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/tolu-c/rss-go/internal/app"
	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/store"
)

type Feeds interface {
	Create(ctx context.Context, name, url string, userID uuid.UUID) (model.Feed, error)
	List(ctx context.Context) ([]model.Feed, error)
}

type service struct {
	dp app.Dependency
}

func New(dp app.Dependency) Feeds {
	return &service{dp: dp}
}

func (s *service) Create(ctx context.Context, name, url string, userID uuid.UUID) (model.Feed, error) {
	now := time.Now().UTC()
	return s.dp.Store.CreateFeed(ctx, store.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    userID,
	})
}

func (s *service) List(ctx context.Context) ([]model.Feed, error) {
	return s.dp.Store.GetFeeds(ctx)
}
