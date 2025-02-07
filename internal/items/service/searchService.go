package service

import (
	"context"
	"log"
	"time"

	"github.com/ruziba3vich/prosphere/internal/items/repo"
	"github.com/ruziba3vich/prosphere/internal/items/storage/search"
)

type (
	SearchService struct {
		storage repo.SearchRepository
		logger  *log.Logger
	}
)

func NewSearchService(storage repo.SearchRepository, logger *log.Logger) repo.SearchRepository {
	return &SearchService{
		storage: storage,
		logger:  logger,
	}
}

func (s *SearchService) SearchElement(ctx context.Context, req *search.SearchRequest) (*search.SearchResponse, error) {
	s.logger.Printf("-- RECEIVED A NEW REQUEST INTO SearchService -- %s\n", time.Now().String())
	return s.storage.SearchElement(ctx, req)
}
