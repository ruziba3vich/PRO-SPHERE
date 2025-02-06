package repo

import (
	"context"

	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
)

type (
	FeedsRepository interface {
		CreateFeed(ctx context.Context, feed *models.Feed) (*models.Feed, error)
		UpdateFeed(ctx context.Context, feed *models.Feed) (*models.Feed, error)
		DeleteFeed(ctx context.Context, feedID int) error
		GetFeed(ctx context.Context, feedID int) (*models.Feed, error)
		GetAllFeeds(ctx context.Context, lang *string, limit, page int) ([]models.Feed, error)
		GetFeedLogoInfo(ctx context.Context, feedId int) (*models.FeedLogoInfo, error)

		AddFeedContent(ctx context.Context, content *models.FeedContent) (*models.FeedContent, error)
		UpdateFeedContent(ctx context.Context, id int, link *string, lang *string, categoryId *int) (*models.FeedContent, error)
		DeleteFeedContent(ctx context.Context, id int) error
		GetAllFeedContent(ctx context.Context, feedId int, lang *string, categoryId *int) ([]models.FeedContent, error)
		GetFeedContent(ctx context.Context, id int) (*models.FeedContent, error)
	}
)
