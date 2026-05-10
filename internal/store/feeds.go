package store

import (
	"context"

	"github.com/tolu-c/rss-go/internal/model"
	"github.com/tolu-c/rss-go/internal/store/queries"
)

type CreateFeedParams = queries.CreateFeedParams

func (s *Store) CreateFeed(ctx context.Context, arg CreateFeedParams) (model.Feed, error) {
	row, err := s.queries.CreateFeed(ctx, arg)
	if err != nil {
		return model.Feed{}, err
	}
	return feedFromRow(row), nil
}

func (s *Store) GetFeeds(ctx context.Context) ([]model.Feed, error) {
	rows, err := s.queries.GetFeeds(ctx)
	if err != nil {
		return nil, err
	}
	feeds := make([]model.Feed, 0, len(rows))
	for _, row := range rows {
		feeds = append(feeds, feedFromRow(row))
	}
	return feeds, nil
}

func feedFromRow(f queries.Feed) model.Feed {
	return model.Feed{
		ID:        f.ID,
		CreatedAt: f.CreatedAt,
		UpdatedAt: f.UpdatedAt,
		Name:      f.Name,
		Url:       f.Url,
		UserID:    f.UserID,
	}
}
