package handlers

import (
	"context"
	"time"

	"github.com/projects/pro-sphere-backend/admin/genproto/genproto/feeds"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FeedServiceServer struct {
	feeds.UnimplementedFeedsServiceServer
	service *service.FeedsService
	logger  *zap.Logger
}

func NewFeedServiceServer(s *service.FeedsService, l *zap.Logger) *FeedServiceServer {
	return &FeedServiceServer{
		service: s,
		logger:  l,
	}
}
func feedTransalationToModels(tr []*feeds.FeedTransalation) []models.Translation {
	var trs []models.Translation
	for _, t := range tr {
		trs = append(trs, models.Translation{
			Id:          int(t.Id),
			Lang:        t.Lang,
			Title:       t.Title,
			Description: t.Description,
			FeedId:      int(t.FeedId),
		})
	}
	return trs
}
func (s *FeedServiceServer) CreateFeed(ctx context.Context, req *feeds.CreateFeedRequest) (*feeds.Feed, error) {
	var transalations []models.Translation
	for _, t := range req.Translation {
		transalations = append(transalations, models.Translation{
			Title:       t.Title,
			Lang:        t.Lang,
			Description: t.Description,
		})
	}
	createdFeed, err := s.service.CreateFeed(ctx, &models.Feed{
		LogoUrl:      req.LogoUrl,
		LogoUrlId:    req.LogoUrlId,
		Priority:     int(req.Priority),
		MaxItems:     int(req.MaxItems),
		BaseUrl:      req.BaseUrl,
		Translations: transalations,
	})

	if err != nil {
		s.logger.Error("Failed to create feed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to create feed: %v", err)
	}
	return convertFeedToProto(createdFeed), nil
}

func (s *FeedServiceServer) UpdateFeed(ctx context.Context, req *feeds.UpdateFeedRequest) (*feeds.Feed, error) {
	// Create a context with a 2-second timeout
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	updatedFeed, err := s.service.UpdateFeed(ctx, &models.Feed{
		Id:           int(req.Id),
		LogoUrl:      req.LogoUrl,
		LogoUrlId:    req.LogoUrlId,
		Priority:     int(req.Priority),
		MaxItems:     int(req.MaxItems),
		BaseUrl:      req.BaseUrl,
		Translations: feedTransalationToModels(req.Translation),
	})
	if err != nil {
		// Check if the error was caused by a context timeout
		if ctx.Err() == context.DeadlineExceeded {
			s.logger.Error("Update feed operation timed out", zap.Error(err))
			return nil, status.Errorf(codes.DeadlineExceeded, "Update feed operation timed out")
		}
		s.logger.Error("Failed to update feed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to update feed: %v", err)
	}

	return convertFeedToProto(updatedFeed), nil
}

func (s *FeedServiceServer) DeleteFeed(ctx context.Context, req *feeds.FeedRequest) (*feeds.EmptyResponse, error) {
	_, err := s.service.DeleteFeed(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete feed", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to delete feed: %v", err)
	}
	return &feeds.EmptyResponse{}, nil
}

func (s *FeedServiceServer) GetFeed(ctx context.Context, req *feeds.FeedRequest) (*feeds.Feed, error) {
	feed, err := s.service.GetFeed(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get feed", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "Feed not found: %v", err)
	}
	return convertFeedToProto(feed), nil
}

func (s *FeedServiceServer) GetAllFeeds(ctx context.Context, req *feeds.GetAllFeedsRequest) (*feeds.FeedsResponse, error) {
	feedsList, err := s.service.GetAllFeeds(ctx, &req.Lang, int(req.Limit), int(req.Page))
	if err != nil {
		s.logger.Error("Failed to get all feeds", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to get all feeds: %v", err)
	}
	var protoFeeds []*feeds.Feed
	for _, feed := range feedsList {
		protoFeeds = append(protoFeeds, convertFeedToProto(&feed))
	}
	return &feeds.FeedsResponse{Feeds: protoFeeds}, nil
}

// Helper function to convert Feed model to protobuf Feed
func convertFeedToProto(feed *models.Feed) *feeds.Feed {
	return &feeds.Feed{
		Id:          int64(feed.Id),
		BaseUrl:     feed.BaseUrl,
		Priority:    int32(feed.Priority),
		LogoUrl:     feed.LogoUrl,
		LogoUrlId:   feed.LogoUrlId,
		Translation: convertTranslationToProto(feed.Translations, feed.Id),
	}
}

func convertTranslationToProto(tr []models.Translation, feedId int) []*feeds.FeedTransalation {
	var translations []*feeds.FeedTransalation
	for _, t := range tr {
		translations = append(translations, &feeds.FeedTransalation{
			Id:          int64(t.Id),
			FeedId:      int64(feedId),
			Title:       t.Title,
			Lang:        t.Lang,
			Description: t.Description,
		})
	}
	return translations
}

// Helper function to convert protobuf Feed to model Feed
func convertProtoToFeed(protoFeed *feeds.Feed) *models.Feed {
	var translations []models.Translation
	for _, tr := range protoFeed.Translation {
		translations = append(translations, models.Translation{
			Id:          int(tr.Id),
			Title:       tr.Title,
			Description: tr.Description,
			Lang:        tr.Lang,
		})
	}
	return &models.Feed{
		Id:           int(protoFeed.Id),
		BaseUrl:      protoFeed.BaseUrl,
		Priority:     int(protoFeed.Priority),
		LogoUrl:      protoFeed.LogoUrl,
		LogoUrlId:    protoFeed.LogoUrlId,
		Translations: translations,
		MaxItems:     int(protoFeed.MaxItems),
	}
}

func (s *FeedServiceServer) AddFeedContent(ctx context.Context, req *feeds.FeedContent) (*feeds.FeedContent, error) {
	newContent, err := s.service.AddFeedContent(ctx, convertProtoToFeedContent(req))
	if err != nil {
		s.logger.Error("Failed to add feed content", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to add feed content: %v", err)
	}
	return convertFeedContentToProto(newContent), nil
}

func (s *FeedServiceServer) UpdateFeedContent(ctx context.Context, req *feeds.FeedContent) (*feeds.FeedContent, error) {
	catId := int(req.CategoryId)
	updatedContent, err := s.service.UpdateFeedContent(ctx, int(req.Id), &req.Link, &req.Lang, &catId)
	if err != nil {
		s.logger.Error("Failed to update feed content", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to update feed content: %v", err)
	}
	return convertFeedContentToProto(updatedContent), nil
}

func (s *FeedServiceServer) DeleteFeedContent(ctx context.Context, req *feeds.FeedRequest) (*feeds.FeedContent, error) {
	err := s.service.DeleteFeedContent(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to delete feed content", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to delete feed content: %v", err)
	}
	return &feeds.FeedContent{Id: req.Id}, nil
}

func (s *FeedServiceServer) GetAllFeedContent(ctx context.Context, req *feeds.GetFeedContents) (*feeds.FeedContentsRespose, error) {
	catId := int(req.CategoryId)

	contents, err := s.service.GetAllFeedContent(ctx, int(req.FeedId), &req.Lang, &catId)
	if err != nil {
		s.logger.Error("Failed to get all feed contents", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "Failed to get all feed contents: %v", err)
	}
	var protoContents []*feeds.FeedContent
	for _, content := range contents {
		protoContents = append(protoContents, convertFeedContentToProto(&content))
	}
	return &feeds.FeedContentsRespose{Contents: protoContents}, nil
}

func (s *FeedServiceServer) GetFeedContent(ctx context.Context, req *feeds.FeedRequest) (*feeds.FeedContent, error) {
	content, err := s.service.GetFeedContent(ctx, int(req.Id))
	if err != nil {
		s.logger.Error("Failed to get feed content", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, "Feed content not found: %v", err)
	}
	return convertFeedContentToProto(content), nil
}

// Helper functions for FeedContent conversions
func convertFeedContentToProto(content *models.FeedContent) *feeds.FeedContent {
	return &feeds.FeedContent{
		Id:         int64(content.Id),
		FeedId:     int64(content.FeedId),
		CategoryId: int64(content.CategoryId),
		Lang:       content.Lang,
		Link:       content.Link,
	}
}

func convertProtoToFeedContent(protoContent *feeds.FeedContent) *models.FeedContent {
	return &models.FeedContent{
		Id:         int(protoContent.Id),
		FeedId:     int(protoContent.FeedId),
		CategoryId: int(protoContent.CategoryId),
		Lang:       protoContent.Lang,
		Link:       protoContent.Link,
	}
}
