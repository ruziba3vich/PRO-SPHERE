package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/internal/items/repo"
	"go.uber.org/zap"
)

type FeedItemsService struct {
	feedRepo   repo.FeedsRepository
	fItemsRepo repo.FeedItemsRepository
	logger     *zap.Logger
	feedParser *gofeed.Parser
}

func NewFeedItemsService(feedRepo repo.FeedsRepository, fItemsRepo repo.FeedItemsRepository, logger *zap.Logger, parser *gofeed.Parser) *FeedItemsService {
	return &FeedItemsService{
		feedRepo:   feedRepo,
		fItemsRepo: fItemsRepo,
		logger:     logger,
		feedParser: parser,
	}
}

// fetchAndParseFeed retrieves feed content from a URL and parses it.
func (f *FeedItemsService) fetchAndParseFeed(ctx context.Context, feedID int) ([]*models.FeedItemResponse, error) {
	// Fetch feed details from the repository
	feed, err := f.feedRepo.GetFeed(ctx, feedID)
	if err != nil {
		f.logger.Error("Failed to get feed from repository", zap.Int("feedID", feedID), zap.Error(err))
		return nil, err
	}

	// Validate feed link
	if feed.BaseUrl == "" {
		return nil, errors.New("invalid feed link")
	}

	// Make HTTP request
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, feed.BaseUrl, nil)
	if err != nil {
		f.logger.Error("Failed to create HTTP request", zap.String("feedLink", feed.BaseUrl), zap.Error(err))
		return nil, err
	}

	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil || httpResponse.Body == nil {
		f.logger.Error("Failed to fetch feed", zap.String("feedLink", feed.BaseUrl), zap.Error(err))
		return nil, errors.New("failed to fetch feed or empty response")
	}
	defer httpResponse.Body.Close()

	// Parse feed content
	parsedFeed, err := f.feedParser.Parse(httpResponse.Body)
	if err != nil {
		f.logger.Error("Failed to parse feed content", zap.String("feedLink", feed.BaseUrl), zap.Error(err))
		return nil, err
	}

	// Transform parsed feed items into FeedItemResponse
	var feedItems []*models.FeedItemResponse
	for _, item := range parsedFeed.Items {
		feedItems = append(feedItems, &models.FeedItemResponse{
			FeedID:      feedID,
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			ImageURL:    item.Image.URL,
			PublishedAt: item.Published,
		})
	}

	return feedItems, nil
}

// FetchFeedItems retrieves and stores feed items in the database.
func (f *FeedItemsService) FetchFeedItems(ctx context.Context, feedID int) ([]*models.FeedItemResponse, error) {
	feedItems, err := f.fetchAndParseFeed(ctx, feedID)
	if err != nil {
		return nil, err
	}

	// Store fetched feed items asynchronously
	go f.storeFeedItems(feedID, feedItems)

	return feedItems, nil
}

// UpdateFeed updates the feed by fetching new items and storing them.
func (f *FeedItemsService) UpdateFeed(ctx context.Context, feedID int) ([]*models.FeedItemResponse, error) {
	feedItems, err := f.fetchAndParseFeed(ctx, feedID)
	if err != nil {
		return nil, err
	}

	// Store fetched feed items asynchronously
	go f.storeFeedItems(feedID, feedItems)

	return feedItems, nil
}

// storeFeedItems saves feed items into the database.
func (f *FeedItemsService) storeFeedItems(feedID int, feedItems []*models.FeedItemResponse) {
	for _, item := range feedItems {
		itemCtx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
		defer cancel()

		_, err := f.fItemsRepo.CreateFeedItem(itemCtx, &models.FeedItem{
			FeedID:      feedID,
			Title:       item.Title,
			Link:        item.Link,
			ImageURL:    item.ImageURL,
			Description: item.Description,
			PublishedAt: item.PublishedAt,
		})

		if err != nil {
			f.logger.Error("Failed to insert feed item into database", zap.String("title", item.Title), zap.Error(err))
		}
	}
}
