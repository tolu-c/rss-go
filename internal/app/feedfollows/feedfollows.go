// Package feedfollows is the service layer for feed-follow operations.
package feedfollows

import (
	"context"
	"time"

	"github.com/google/uuid"

	"github.com/tolu-c/rss-go/internal/app"
	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/store"
)

type FeedFollows interface {
	Follow(ctx context.Context, userID, feedID uuid.UUID) (model.FeedFollow, error)
	ListByUser(ctx context.Context, userID uuid.UUID) ([]model.FeedFollow, error)
	Delete(ctx context.Context, userID, feedFollowID uuid.UUID) error
}

type service struct {
	dp app.Dependency
}

func New(dp app.Dependency) FeedFollows {
	return &service{dp: dp}
}

func (s *service) Follow(ctx context.Context, userID, feedID uuid.UUID) (model.FeedFollow, error) {
	now := time.Now().UTC()
	return s.dp.Store.FollowFeed(ctx, store.FollowFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    userID,
		FeedID:    feedID,
	})
}

func (s *service) ListByUser(ctx context.Context, userID uuid.UUID) ([]model.FeedFollow, error) {
	return s.dp.Store.GetFeedFollows(ctx, userID)
}

// Delete scopes the deletion to (feedFollowID, userID) so a request from
// one user cannot remove another user's follow even if the ID is known.
func (s *service) Delete(ctx context.Context, userID, feedFollowID uuid.UUID) error {
	return s.dp.Store.DeleteFeedFollow(ctx, store.DeleteFeedFollowParams{
		ID:     feedFollowID,
		UserID: userID,
	})
}
