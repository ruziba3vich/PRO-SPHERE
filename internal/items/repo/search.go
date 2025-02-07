package repo

import (
	"context"

	"github.com/ruziba3vich/prosphere/internal/items/storage/search"
)

type (
	SearchRepository interface {
		SearchElement(context.Context, *search.SearchRequest) (*search.SearchResponse, error)
		// SearchByPhoto()
		// SearchByVideo()
	}
)
