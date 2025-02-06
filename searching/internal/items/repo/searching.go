package repo

import (
	"context"
	"search_service/genproto/searching"
)

type (
	SearchRepository interface {
		Search(ctx context.Context, req *searching.QueryRequest) (*searching.SearchResponse, error)
		SearchYouTubeVideos(ctx context.Context, req *searching.VideoSearchRequest) (*searching.VideoSearchResponse, error)
		SearchImages(ctx context.Context, req *searching.QueryRequest) (*searching.ImageResponse, error)
	}
)
