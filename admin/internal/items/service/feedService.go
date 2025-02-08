package service

import (
	"context"
	"fmt"
	"os"

	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/internal/items/repo"
	"go.uber.org/zap"
)

var imageDir = "./images/feed/logos"

type (
	FeedsService struct {
		storage repo.FeedsRepository
		logger  *zap.Logger
	}
)

func NewFeedsService(storage repo.FeedsRepository, logger *zap.Logger) *FeedsService {
	return &FeedsService{
		storage: storage,
		logger:  logger,
	}
}

func (f *FeedsService) CreateFeed(ctx context.Context, feed *models.Feed) (*models.Feed, error) {
	// Log the start of the feed creation process
	f.logger.Info("Starting feed creation process",
		zap.String("Base url", feed.BaseUrl),
		zap.Int("Priority", feed.Priority),
		zap.String("Base url", feed.BaseUrl),
		zap.String("logo_url", feed.LogoUrl),
		zap.Any("translations", feed.Translations),
	)

	// Ensure the image directory exists
	if err := os.MkdirAll(imageDir, os.ModePerm); err != nil {
		f.logger.Error("Failed to create image directory",
			zap.String("directory", imageDir),
			zap.Error(err),
		)
		return nil, fmt.Errorf("unable to create image directory %s: %w", imageDir, err)
	}

	// Save and validate the feed logo
	logoUrlId, err := SaveAndValidateImage(feed.LogoUrl, imageDir, f.logger)
	if err != nil {
		f.logger.Error("Failed to save or validate feed logo",
			zap.String("logo_url", feed.LogoUrl),
			zap.String("directory", imageDir),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to save or validate logo: %w", err)
	}
	feed.LogoUrlId = logoUrlId

	// Attempt to create the feed in storage
	createdFeed, err := f.storage.CreateFeed(ctx, feed)
	if err != nil {
		f.logger.Error("Failed to create feed in storage",
			zap.Any("feed_data", feed),
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to create feed in storage: %w", err)
	}

	// Log successful feed creation
	f.logger.Info("Successfully created feed",
		zap.Any("created_feed", createdFeed),
	)

	return createdFeed, nil
}

func (f *FeedsService) UpdateFeed(ctx context.Context, feed *models.Feed) (*models.Feed, error) {
	f.logger.Info("Starting feed update process",
		zap.Int("FeedId", feed.Id),
		zap.String("Base url", feed.BaseUrl),
		zap.Int("Priority", feed.Priority),
		zap.String("Base url", feed.BaseUrl),
		zap.String("logo_url", feed.LogoUrl),
		zap.Any("translations", feed.Translations),
	)
	logoInfo, err := f.storage.GetFeedLogoInfo(ctx, feed.Id)
	if err != nil {
		f.logger.Error("Failed to fetch logo info",
			zap.Int("FeedId", feed.Id),
			zap.Error(err))
		return nil, err
	}
	f.logger.Info("Feed logo info", zap.String("Logo url", logoInfo.LogoUrl), zap.String("Logo id", logoInfo.LogoUrlId))
	// Check if the icon URL has changed
	if logoInfo.LogoUrl != feed.LogoUrl {
		f.logger.Info("Icon URL has changed; downloading new image",
			zap.String("old_url", logoInfo.LogoUrl),
			zap.String("new_url", feed.LogoUrl),
		)

		// Save and validate the new image
		newImageID, err := SaveAndValidateImage(feed.LogoUrl, imageDir, f.logger)
		if err != nil {
			f.logger.Error("Failed to save and validate new image", zap.Error(err))
			return nil, err
		}
		// Remove the out-dated image from the storage
		iconPath := imageDir + "/" + logoInfo.LogoUrlId
		go RemoveImage(iconPath, f.logger)

		// Update the image ID in the feedegory object
		feed.LogoUrlId = newImageID
	} else {
		// If the URL hasn't changed, retain the existing LogoUrlId
		feed.LogoUrlId = logoInfo.LogoUrlId
	}

	updatedFeed, err := f.storage.UpdateFeed(ctx, feed)
	if err != nil {
		f.logger.Error("Failed to update feed", zap.Error(err))
		return nil, err
	}

	return updatedFeed, nil
}

func (f *FeedsService) DeleteFeed(ctx context.Context, feedID int) (*models.Feed, error) {
	f.logger.Info("Starting deleting feed process",
		zap.Int("FeedID", feedID),
	)

	logoInfo, err := f.storage.GetFeedLogoInfo(ctx, feedID)
	if err != nil {
		f.logger.Error("Failed to get logo info", zap.Error(err))
	} else {
		RemoveImage(imageDir+"/"+logoInfo.LogoUrlId, f.logger)
	}

	if err := f.storage.DeleteFeed(ctx, feedID); err != nil {
		return nil, err
	}

	return &models.Feed{Id: feedID}, nil
}

func (f *FeedsService) GetFeed(ctx context.Context, feedID int) (*models.Feed, error) {
	// Fetch the feed from the storage repository.
	feed, err := f.storage.GetFeed(ctx, feedID)
	if err != nil {
		f.logger.Error("Failed to fetch feed by ID", zap.Int("FeedID", feedID), zap.Error(err))
		return nil, fmt.Errorf("failed to fetch feed with ID %d: %w", feedID, err)
	}

	return feed, nil
}

func (f *FeedsService) GetAllFeeds(ctx context.Context, lang *string, limit, page int) ([]models.Feed, error) {
	// Fetch all feeds with pagination and language filter.
	feeds, err := f.storage.GetAllFeeds(ctx, lang, limit, page)
	if err != nil {
		f.logger.Error("Failed to fetch all feeds", zap.Int("limit", limit), zap.Int("page", page), zap.Error(err))
		return nil, fmt.Errorf("failed to fetch all feeds: %w", err)
	}

	return feeds, nil
}

func (f *FeedsService) AddFeedContent(ctx context.Context, content *models.FeedContent) (*models.FeedContent, error) {
	// Validate the feed link before adding content
	ok, err := ValidateFeedLink(ctx, content.Link)
	if err != nil {
		f.logger.Error("Invalid feed link", zap.String("link", content.Link), zap.Error(err))
		return nil, fmt.Errorf("invalid feed link: %w", err)
	}
	if !ok {
		f.logger.Warn("Feed link validation failed", zap.String("link", content.Link))
		return nil, fmt.Errorf("feed link validation failed for: %s", content.Link)
	}

	// Add content to the storage
	createdContent, err := f.storage.AddFeedContent(ctx, content)
	if err != nil {
		f.logger.Error("Failed to add feed content", zap.Any("content", content), zap.Error(err))
		return nil, fmt.Errorf("failed to add feed content: %w", err)
	}

	return createdContent, nil
}

func (f *FeedsService) UpdateFeedContent(ctx context.Context, id int, link *string, lang *string, feedegoryId *int) (*models.FeedContent, error) {
	// Update the content of a specific feed.
	if *link == "" {
		link = nil
	}
	if *lang == "" {
		lang = nil
	}
	if *feedegoryId == 0 {
		feedegoryId = nil
	}

	if link == nil && lang == nil && feedegoryId == nil {
		return nil, fmt.Errorf("all fileds are empty: must be one of them is valid")
	}

	updatedContent, err := f.storage.UpdateFeedContent(ctx, id, link, lang, feedegoryId)
	if err != nil {
		f.logger.Error("Failed to update feed content", zap.Int("ContentID", id), zap.Error(err))
		return nil, fmt.Errorf("failed to update feed content with ID %d: %w", id, err)
	}

	return updatedContent, nil
}

func (f *FeedsService) DeleteFeedContent(ctx context.Context, id int) error {
	// Delete the content of a specific feed.
	err := f.storage.DeleteFeedContent(ctx, id)
	if err != nil {
		f.logger.Error("Failed to delete feed content", zap.Int("ContentID", id), zap.Error(err))
		return fmt.Errorf("failed to delete feed content with ID %d: %w", id, err)
	}

	return nil
}

func (f *FeedsService) GetAllFeedContent(ctx context.Context, feedId int, lang *string, feedegoryId *int) ([]models.FeedContent, error) {
	// Fetch all content for a specific feed.
	if *feedegoryId == 0 {
		feedegoryId = nil
	}
	if *lang == "" {
		lang = nil
	}

	content, err := f.storage.GetAllFeedContent(ctx, feedId, lang, feedegoryId)
	if err != nil {
		f.logger.Error("Failed to fetch all feed content", zap.Int("FeedID", feedId), zap.Error(err))
		return nil, fmt.Errorf("failed to fetch content for feed with ID %d: %w", feedId, err)
	}

	return content, nil
}

func (f *FeedsService) GetFeedContent(ctx context.Context, id int) (*models.FeedContent, error) {
	// Fetch content for a specific feed by content ID.
	content, err := f.storage.GetFeedContent(ctx, id)
	if err != nil {
		f.logger.Error("Failed to fetch feed content by ID", zap.Int("ContentID", id), zap.Error(err))
		return nil, fmt.Errorf("failed to fetch feed content with ID %d: %w", id, err)
	}

	return content, nil
}
