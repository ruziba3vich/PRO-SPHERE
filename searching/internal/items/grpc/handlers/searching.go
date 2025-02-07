package handlers

import (
	"context"
	"search_service/genproto/searching"
	"search_service/internal/items/service"

	"go.uber.org/zap"
)

type (
	SearchingHandler struct {
		searchService *service.SearchService
		logger        *zap.Logger
		searching.UnimplementedSearchingServiceServer
	}
)

func NewSerachHandler(searchSer *service.SearchService, logger *zap.Logger) *SearchingHandler {
	return &SearchingHandler{
		searchService: searchSer,
		logger:        logger,
	}
}

func (s *SearchingHandler) Search(ctx context.Context, req *searching.QueryRequest) (*searching.SearchResponse, error) {
	return s.searchService.Search(ctx, req)
}
func (s *SearchingHandler) SearchImages(ctx context.Context, req *searching.QueryRequest) (*searching.ImageResponse, error) {
	return s.searchService.SearchImages(ctx, req)

}
func (s *SearchingHandler) SearchVideos(ctx context.Context, req *searching.VideoSearchRequest) (*searching.VideoSearchResponse, error) {
	return s.searchService.SearchVideos(ctx, req)
}
