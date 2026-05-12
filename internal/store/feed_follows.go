package store

import (
	"context"

	"github.com/google/uuid"

	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/store/queries"
)

type FollowFeedParams = queries.FollowFeedParams
type DeleteFeedFollowParams = queries.DeleteFeedFollowParams

func (s *Store) FollowFeed(ctx context.Context, arg FollowFeedParams) (model.FeedFollow, error) {
	row, err := s.queries.FollowFeed(ctx, arg)
	if err != nil {
		return model.FeedFollow{}, err
	}
	return feedFollowFromRow(row), nil
}

func (s *Store) GetFeedFollows(ctx context.Context, userID uuid.UUID) ([]model.FeedFollow, error) {
	rows, err := s.queries.GetFeedFollows(ctx, userID)
	if err != nil {
		return nil, err
	}
	follows := make([]model.FeedFollow, 0, len(rows))
	for _, row := range rows {
		follows = append(follows, feedFollowFromRow(row))
	}
	return follows, nil
}

func (s *Store) DeleteFeedFollow(ctx context.Context, arg DeleteFeedFollowParams) error {
	return s.queries.DeleteFeedFollow(ctx, arg)
}

func feedFollowFromRow(ff queries.FeedFollow) model.FeedFollow {
	return model.FeedFollow{
		ID:        ff.ID,
		CreatedAt: ff.CreatedAt,
		UpdatedAt: ff.UpdatedAt,
		UserID:    ff.UserID,
		FeedID:    ff.FeedID,
	}
}
