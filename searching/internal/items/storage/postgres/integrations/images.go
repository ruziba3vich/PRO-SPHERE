/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-20 21:42:18
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-06 00:16:43
 * @FilePath: /searching/internal/items/storage/postgres/integrations/images.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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
	"strconv"

	"go.uber.org/zap"
)

func (s *SearchIntegration) SearchImages(ctx context.Context, req *searching.QueryRequest) (*searching.ImageResponse, error) {
	// Validate the query
	if req.Query == "" {
		return nil, fmt.Errorf("query cannot be empty")
	}

	// Encode the query and calculate the start index for pagination
	encodedQuery := url.QueryEscape(req.Query)
	start := (req.PageNumber-1)*req.PageSize + 1

	// Build the search URL
	searchURL := s.cfg.ImagesSearch()
	searchURL += "&q=" + encodedQuery
	searchURL += fmt.Sprintf("&start=%d", start)
	searchURL += fmt.Sprintf("&num=%d", req.PageSize)

	// Create an HTTP request
	httpRequest, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		s.logger.Error("Creating new request failed", zap.Error(err))
		return nil, err
	}

	// Send the HTTP request
	httpResponse, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		s.logger.Error("Sending request failed", zap.Error(err))
		return nil, err
	}
	defer httpResponse.Body.Close()

	s.logger.Debug("Request URL", zap.String("url", searchURL))
	if httpResponse.StatusCode == 429 {
		s.logger.Warn("Google API reached the limit for today!", zap.Int("status_code", httpResponse.StatusCode))
		return nil, fmt.Errorf("google api reached the limit for today")
	}

	// Check for non-OK HTTP response
	if httpResponse.StatusCode != http.StatusOK {
		s.logger.Warn("Request failed with non-OK status", zap.Int("status_code", httpResponse.StatusCode))
		return nil, fmt.Errorf("unexpected status code: %d", httpResponse.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		s.logger.Error("Reading from HTTP body failed", zap.Error(err))
		return nil, err
	}

	// Parse the API response
	parsedResponse := models.ImageSearchAPIResponse{}
	if err := json.Unmarshal(body, &parsedResponse); err != nil {
		s.logger.Error("Unmarshalling body failed", zap.Error(err))
		return nil, err
	}

	// Map the parsed response to the gRPC response format
	imageResults := []*searching.ImageResult{}
	for _, item := range parsedResponse.Items {
		imageResults = append(imageResults, &searching.ImageResult{
			Title:    item.Title,
			ImageUrl: item.Link,
			Width:    int32(item.Image.Width),
			Height:   int32(item.Image.Height),
		})
	}
	totalItems, err := strconv.Atoi(parsedResponse.SearchInfo.TotalResults)
	if err != nil {
		s.logger.Error("Failed to convert total items string into int using Atoi", zap.Error(err))
	}
	// Construct and return the final response
	return &searching.ImageResponse{
		TotalItems: int32(totalItems),
		Images:     imageResults,
	}, nil
}
