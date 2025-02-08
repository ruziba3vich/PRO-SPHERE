/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-19 16:00:12
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-11-20 23:51:16
 * @FilePath: /searching/internal/items/service/searching.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"search_service/genproto/searching"
	"search_service/internal/items/repo"
	"time"

	"go.uber.org/zap"
)

type (
	SearchService struct {
		storage repo.SearchRepository
		logger  *zap.Logger
	}
)

func NewSearchService(storage repo.SearchRepository, logger *zap.Logger) *SearchService {
	return &SearchService{
		storage: storage,
		logger:  logger,
	}
}

func (s *SearchService) Search(ctx context.Context, req *searching.QueryRequest) (*searching.SearchResponse, error) {
	s.logger.Info("Received new search request",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Any("request", req),
	)

	return s.storage.Search(ctx, req)
}

func (s *SearchService) SearchVideos(ctx context.Context, req *searching.VideoSearchRequest) (*searching.VideoSearchResponse, error) {
	s.logger.Info("Received new video search request",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Any("request", req),
	)

	return s.storage.SearchYouTubeVideos(ctx, req)
}

func (s *SearchService) SearchImages(ctx context.Context, req *searching.QueryRequest) (*searching.ImageResponse, error) {
	s.logger.Info("Received new image search request",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Any("request", req),
	)

	return s.storage.SearchImages(ctx, req)
}
