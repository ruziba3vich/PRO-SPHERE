/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:39:30
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-22 06:14:18
 * @FilePath: /admin/internal/items/gRPC/handlers/feeds_category.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package handlers

import (
	"context"

	"github.com/projects/pro-sphere-backend/admin/genproto/genproto/feeds"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/status"
)

type (
	FeedsCategoryHandler struct {
		categoryService *service.CategoriesService
		logger          *zap.Logger
		feeds.UnimplementedCategoriesServiceServer
	}
)

func NewCategoryHandler(catSer *service.CategoriesService, logger *zap.Logger) *FeedsCategoryHandler {
	return &FeedsCategoryHandler{
		categoryService: catSer,
		logger:          logger,
	}
}

// CreateFeedCategory creates a new feed category
func (h *FeedsCategoryHandler) CreateFeedCategory(ctx context.Context, req *feeds.CreateFeedCategoryRequest) (*feeds.FeedCategory, error) {
	// Convert the request into the internal FeedCategory model
	category := &models.FeedCategory{
		IconURL:      req.IconUrl,
		Translations: ConvertProtoToFeedCategoryTranslation(req.Translations),
	}

	// Call the CategoriesService to create the category
	newCategory, err := h.categoryService.CreateFeedCategory(ctx, category.IconURL, category.Translations)
	if err != nil {
		h.logger.Error("Failed to create feed category", zap.Error(err))
		return nil, status.Errorf(status.Code(err), "failed to create feed category: %v", err)
	}

	// Convert the internal FeedCategory model back to protobuf
	return ConvertFeedCategoryToProto(newCategory), nil
}

// GetFeedCategoryByID retrieves a feed category by its ID
func (h *FeedsCategoryHandler) GetFeedCategoryByID(ctx context.Context, req *feeds.GetFeedCategoryByIDRequest) (*feeds.FeedCategory, error) {
	// Call the CategoriesService to fetch the category by ID
	category, err := h.categoryService.GetFeedCategoryByID(ctx, req.Id, req.Lang)
	if err != nil {
		h.logger.Error("Failed to fetch feed category", zap.Error(err))
		return nil, status.Errorf(status.Code(err), "failed to fetch feed category: %v", err)
	}

	// Convert the internal FeedCategory model to protobuf and return it
	return ConvertFeedCategoryToProto(category), nil
}

// UpdateFeedCategory updates an existing feed category
func (h *FeedsCategoryHandler) UpdateFeedCategory(ctx context.Context, req *feeds.FeedCategory) (*feeds.FeedCategory, error) {
	// Convert the request into the internal FeedCategory model
	category := ConvertProtoToFeedCategory(req)

	// Call the CategoriesService to update the category
	updatedCategory, err := h.categoryService.UpdateFeedCategory(ctx, category)
	if err != nil {
		h.logger.Error("Failed to update feed category", zap.Error(err))
		return nil, status.Errorf(status.Code(err), "failed to update feed category: %v", err)
	}

	// Convert the internal FeedCategory model back to protobuf and return it
	return ConvertFeedCategoryToProto(updatedCategory), nil
}

// DeleteFeedCategory deletes a feed category by its ID
func (h *FeedsCategoryHandler) DeleteFeedCategory(ctx context.Context, req *feeds.DeleteFeedCategoryRequest) (*feeds.EmptyResponse, error) {
	// Call the CategoriesService to delete the category
	err := h.categoryService.DeleteFeedCategory(ctx, req.Id)
	if err != nil {
		h.logger.Error("Failed to delete feed category", zap.Error(err))
		return nil, status.Errorf(status.Code(err), "failed to delete feed category: %v", err)
	}

	// Return an empty response after successful deletion
	return &feeds.EmptyResponse{}, nil
}

// GetAllFeedCategories retrieves all feed categories with pagination
func (h *FeedsCategoryHandler) GetAllFeedCategories(ctx context.Context, req *feeds.GetAllFeedCategoriesRequest) (*feeds.FeedCategoriesResponse, error) {
	// Call the CategoriesService to fetch all feed categories
	categories, err := h.categoryService.GetAllFeedCategories(ctx, int(req.Limit), int(req.Page), req.Lang)
	if err != nil {
		h.logger.Error("Failed to fetch all feed categories", zap.Error(err))
		return nil, status.Errorf(status.Code(err), "failed to fetch all feed categories: %v", err)
	}

	// Convert the internal FeedCategory models to protobuf
	var feedCategories []*feeds.FeedCategory
	for _, category := range categories {
		feedCategories = append(feedCategories, ConvertFeedCategoryToProto(category))
	}

	// Return the list of categories in the FeedCategoriesResponse
	return &feeds.FeedCategoriesResponse{Categories: feedCategories}, nil
}
