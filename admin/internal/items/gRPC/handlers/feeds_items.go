package handlers

import (
	"context"

	"github.com/projects/pro-sphere-backend/admin/genproto/genproto/feeds"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/internal/items/storage"
	"go.uber.org/zap"
)

type (
	FeedsItemsHandler struct {
		feedItems *storage.FeedItemsStorage
		logger    *zap.Logger
		feeds.UnimplementedFeedItemsServiceServer
	}
)

func NewFeedItemsHandler(feedItems *storage.FeedItemsStorage, logger *zap.Logger) *FeedsItemsHandler {
	return &FeedsItemsHandler{
		feedItems: feedItems,
		logger:    logger,
	}
}

func (f *FeedsItemsHandler) CreateFeedItem(ctx context.Context, req *feeds.FeedItem) (*feeds.FeedItemResponse, error) {
	feedItem := &models.FeedItem{
		FeedID:      int(req.FeedId),
		Title:       req.Title,
		ImageURL:    req.ImageUrl,
		Description: req.Description,
		PublishedAt: req.PublishedAt,
	}

	createdItem, err := f.feedItems.CreateFeedItem(ctx, feedItem)
	if err != nil {
		f.logger.Error("Failed to create feed item", zap.Error(err))
		return nil, err
	}

	return &feeds.FeedItemResponse{
		Id:          int32(createdItem.ID),
		FeedId:      int32(createdItem.FeedID),
		Title:       createdItem.Title,
		Description: createdItem.Description,
		ImageUrl:    createdItem.ImageURL,
		PublishedAt: createdItem.PublishedAt,
		CreatedAt:   createdItem.CreatedAt,
		UpdatedAt:   createdItem.UpdatedAt,
	}, nil
}

func (f *FeedsItemsHandler) DeleteFeedItem(ctx context.Context, req *feeds.DeleteFeedItemRequest) (*feeds.EmptyResponse, error) {
	err := f.feedItems.DeleteFeedItem(ctx, int(req.Id))
	if err != nil {
		f.logger.Error("Failed to delete feed item", zap.Error(err))
		return nil, err
	}
	return &feeds.EmptyResponse{}, nil
}

func (f *FeedsItemsHandler) GetAllFeedItems(ctx context.Context, req *feeds.GetAllFeedItemsRequest) (*feeds.FeedItemsResponse, error) {
	feedItems, err := f.feedItems.GetAllFeedItems(ctx, int(req.Limit), int(req.Page))
	if err != nil {
		f.logger.Error("Failed to get all feed items", zap.Error(err))
		return nil, err
	}

	feedItemResponses := make([]*feeds.FeedItemResponse, len(feedItems))
	for i, item := range feedItems {
		feedItemResponses[i] = &feeds.FeedItemResponse{
			Id:          int32(item.ID),
			FeedId:      int32(item.FeedID),
			Title:       item.Title,
			Description: item.Description,
			ImageUrl:    item.ImageURL,
			PublishedAt: item.PublishedAt,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		}
	}

	return &feeds.FeedItemsResponse{Items: feedItemResponses}, nil
}

func (f *FeedsItemsHandler) GetFeedItemByID(ctx context.Context, req *feeds.GetFeedItemByIDRequest) (*feeds.FeedItemResponse, error) {
	feedItem, err := f.feedItems.GetFeedItemByID(ctx, int(req.Id))
	if err != nil {
		f.logger.Error("Failed to get feed item by ID", zap.Error(err))
		return nil, err
	}

	return &feeds.FeedItemResponse{
		Id:          int32(feedItem.ID),
		FeedId:      int32(feedItem.FeedID),
		Title:       feedItem.Title,
		Description: feedItem.Description,
		ImageUrl:    feedItem.ImageURL,
		PublishedAt: feedItem.PublishedAt,
		CreatedAt:   feedItem.CreatedAt,
		UpdatedAt:   feedItem.UpdatedAt,
	}, nil
}

func (f *FeedsItemsHandler) UpdateFeedItem(ctx context.Context, req *feeds.UpdateItem) (*feeds.FeedItemResponse, error) {
	updateItem := &models.UpdateItem{
		Id:          int(req.Id),
		Title:       req.Title,
		Description: req.Description,
		PublishedAt: req.PublishedAt,
		ImageURL:    req.ImageUrl,
	}

	updatedItem, err := f.feedItems.UpdateFeedItem(ctx, updateItem)
	if err != nil {
		f.logger.Error("Failed to update feed item", zap.Error(err))
		return nil, err
	}

	return &feeds.FeedItemResponse{
		Id:          int32(updatedItem.ID),
		FeedId:      int32(updatedItem.FeedID),
		Title:       updatedItem.Title,
		Description: updatedItem.Description,
		ImageUrl:    updatedItem.ImageURL,
		PublishedAt: updatedItem.PublishedAt,
		CreatedAt:   updatedItem.CreatedAt,
		UpdatedAt:   updatedItem.UpdatedAt,
	}, nil
}
