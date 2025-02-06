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

type (
	FeedItemsService struct {
		feedRepo   repo.FeedsRepository
		fItemsRepo repo.FeedItemsRepository
		logger     *zap.Logger
		feedParser *gofeed.Parser
	}
)

func NewFeedItemsService(feedRepo repo.FeedsRepository, fItemsRepo repo.FeedItemsRepository, logger *zap.Logger, parser *gofeed.Parser) *FeedItemsService {
	return &FeedItemsService{
		feedParser: parser,
		feedRepo:   feedRepo,
		fItemsRepo: fItemsRepo,
		logger:     logger,
	}
}

func (f *FeedItemsService) FetchFeedItems(ctx context.Context, feedID int) ([]*models.FeedItemResponse, error) {
	// Fetch feed details from the repository
	// feed, err := f.feedRepo.GetFeed(ctx, feedID)
	// if err != nil {
	// 	f.logger.Error("Error occurred during getting feed", zap.Error(err))
	// 	return nil, err
	// }
	f.logger.Info("Fetched feed from database", zap.Any("", "feed"))
	// Create HTTP request to fetch the feed
	httpRequest, err := http.NewRequest(http.MethodGet, "feed.Link", nil)
	if err != nil {
		f.logger.Error("Error during creating new HTTP request", zap.Error(err))
		return nil, err
	}

	// Execute the HTTP request
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		f.logger.Error("Error while getting HTTP response", zap.Error(err))
		return nil, err
	}

	if httpResponse == nil || httpResponse.Body == nil {
		return nil, errors.New("invalid http response or empty body")
	}

	defer httpResponse.Body.Close()
	// Parse the feed using gofeed
	parsedFeed, err := f.feedParser.Parse(httpResponse.Body)
	if err != nil {
		f.logger.Error("Error while parsing feed items", zap.Error(err))
		return nil, err
	}
	f.logger.Info("Feed service: ", zap.Any("Feed Items", parsedFeed.String()))

	// Map parsed feed items to your FeedItemResponse model
	var feedItems []*models.FeedItemResponse
	for _, item := range parsedFeed.Items {
		feedItem := &models.FeedItemResponse{
			FeedID:      feedID,
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			ImageURL:    item.Image.URL,
			PublishedAt: item.Published,
			CreatedAt:   "",
			UpdatedAt:   "",
		}

		feedItems = append(feedItems, feedItem)
	}
	go func() {
		for _, item := range feedItems {
			itemCtx, _ := context.WithTimeout(context.Background(), time.Millisecond*200)

			if _, err := f.fItemsRepo.CreateFeedItem(itemCtx, &models.FeedItem{
				FeedID:      feedID,
				Title:       item.Title,
				ImageURL:    item.ImageURL,
				Description: item.Description,
				PublishedAt: item.PublishedAt,
			}); err != nil {
				f.logger.Error("Error while Inserting feed item into Database",
					zap.Error(err))
			}

		}
	}()

	return feedItems, nil
}

// func (f *FeedItemsService) SaveFetchedItems() (models.Feed, error) {

// }

func (f *FeedItemsService) UpdateFeed(ctx context.Context, feedID int) ([]*models.FeedItemResponse, error) {
	// Fetch feed details from the repository
	// feed, err := f.feedRepo.GetFeed(ctx, feedID)
	// if err != nil {
	// 	f.logger.Error("Error occurred during getting feed", zap.Error(err))
	// 	return nil, err
	// }
	f.logger.Info("Fetched feed from database", zap.Any("", "feed"))
	// Create HTTP request to fetch the feed
	httpRequest, err := http.NewRequest(http.MethodGet, "feed.Link", nil)
	if err != nil {
		f.logger.Error("Error during creating new HTTP request", zap.Error(err))
		return nil, err
	}

	// Execute the HTTP request
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		f.logger.Error("Error while getting HTTP response", zap.Error(err))
		return nil, err
	}

	if httpResponse == nil || httpResponse.Body == nil {
		return nil, errors.New("invalid http response or empty body")
	}

	defer httpResponse.Body.Close()
	// Parse the feed using gofeed
	parsedFeed, err := f.feedParser.Parse(httpResponse.Body)
	if err != nil {
		f.logger.Error("Error while parsing feed items", zap.Error(err))
		return nil, err
	}
	f.logger.Info("Feed service: ", zap.Any("Feed Items", parsedFeed.String()))

	// Map parsed feed items to your FeedItemResponse model
	var feedItems []*models.FeedItemResponse
	for _, item := range parsedFeed.Items {
		feedItem := &models.FeedItemResponse{
			FeedID:      feedID,
			Title:       item.Title,
			Link:        item.Link,
			Description: item.Description,
			ImageURL:    item.Image.URL,
			PublishedAt: item.Published,
			CreatedAt:   "",
			UpdatedAt:   "",
		}

		feedItems = append(feedItems, feedItem)
	}
	go func() {
		for _, item := range feedItems {
			itemCtx, _ := context.WithTimeout(context.Background(), time.Millisecond*200)

			if _, err := f.fItemsRepo.CreateFeedItem(itemCtx, &models.FeedItem{
				FeedID:      feedID,
				Title:       item.Title,
				ImageURL:    item.ImageURL,
				Description: item.Description,
				PublishedAt: item.PublishedAt,
			}); err != nil {
				f.logger.Error("Error while Inserting feed item into Database",
					zap.Error(err))
			}

		}
	}()

	return feedItems, nil
}
