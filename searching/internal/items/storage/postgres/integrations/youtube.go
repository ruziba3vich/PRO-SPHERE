package integrations

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"search_service/genproto/searching"
	"search_service/internal/items/models"

	"go.uber.org/zap"
)

func (s *SearchIntegration) SearchYouTubeVideos(ctx context.Context, req *searching.VideoSearchRequest) (*searching.VideoSearchResponse, error) {
	// Convert the proto request to the models request
	modelsReq := &models.VideoSearchRequest{
		Query:         req.Query,
		NextPageToken: req.NextPageToken,
		MaxResults:    fmt.Sprintf("%d", req.MaxResults), // Assuming maxResults is an int32, we convert to string
	}

	// Call the function to load YouTube data
	youtubeRes, err := s.LoadYoutubeVideoContent(ctx, modelsReq)
	if err != nil {
		s.logger.Error("Error occurred while loading YouTube videos", zap.Error(err))
		return nil, err
	}

	// Create the response to return
	videoItems := []*searching.VideoResult{}
	for _, item := range youtubeRes.Items {
		resItem := &searching.VideoResult{
			Title:        item.Snippet.Title,
			ChannelTitle: item.Snippet.ChannelTitle,
			ChannelUrl:   fmt.Sprintf("https://www.youtube.com/channel/%s", item.Snippet.ChannelID),
			VideoUrl:     fmt.Sprintf("https://www.youtube.com/watch?v=%s", item.ID.VideoID),
			PublishTime:  item.Snippet.PublishTime,
			Thumbnails: &searching.VideoThumbnails{
				Default: &searching.SearchThumbnail{
					Src:    item.Snippet.Thumbnails.Default.Src,
					Width:  int64(item.Snippet.Thumbnails.Default.Width),
					Height: int64(item.Snippet.Thumbnails.Default.Height),
				},
				Medium: &searching.SearchThumbnail{
					Src:    item.Snippet.Thumbnails.Medium.Src,
					Width:  int64(item.Snippet.Thumbnails.Medium.Width),
					Height: int64(item.Snippet.Thumbnails.Medium.Height),
				},
				High: &searching.SearchThumbnail{
					Src:    item.Snippet.Thumbnails.High.Src,
					Width:  int64(item.Snippet.Thumbnails.High.Width),
					Height: int64(item.Snippet.Thumbnails.High.Height),
				},
			},
		}
		videoItems = append(videoItems, resItem)
	}

	// Return the final response
	return &searching.VideoSearchResponse{
		NextPageToken: youtubeRes.NextPageToken,
		Videos:        videoItems,
		TotalItems:    int32(youtubeRes.PageInfo.TotalResults),
	}, nil
}

func (s *SearchIntegration) LoadYoutubeVideoContent(ctx context.Context, req *models.VideoSearchRequest) (*models.YouTubeSearchResponse, error) {
	encodedQuery := url.QueryEscape(req.Query)

	searchURL := s.cfg.YoutubeSearch() + "&part=snippet"
	searchURL += "&q=" + encodedQuery
	searchURL += "&maxResults=" + req.MaxResults

	if req.NextPageToken != "" {
		searchURL += "&pageToken=" + req.NextPageToken
	}
	fmt.Println(searchURL)
	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		s.logger.Error("Creating YouTube request failed", zap.Error(err))
		return nil, err
	}

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		s.logger.Error("Sending YouTube request failed", zap.Error(err))
		return nil, err
	}
	defer httpRes.Body.Close()
	if httpRes.StatusCode == 429 {
		s.logger.Warn("Google API reached the limit for today!", zap.Int("status_code", httpRes.StatusCode))
		return nil, fmt.Errorf("google api reached the limit for today")
	}

	if httpRes.StatusCode != http.StatusOK {
		s.logger.Warn("YouTube request failed with status", zap.Int("status_code", httpRes.StatusCode))
		return nil, fmt.Errorf("unexpected status code: %d", httpRes.StatusCode)
	}

	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		s.logger.Error("Reading YouTube response body failed", zap.Error(err))
		return nil, err
	}

	var youtubeRes models.YouTubeSearchResponse
	if err := json.Unmarshal(body, &youtubeRes); err != nil {
		s.logger.Error("Unmarshalling YouTube response failed", zap.Error(err))
		return nil, err
	}
	s.logger.Debug("YouTube API Request URL", zap.String("url", searchURL))

	return &youtubeRes, nil
}
