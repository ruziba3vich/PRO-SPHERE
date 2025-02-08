/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-19 20:04:29
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-09 05:53:42
 * @FilePath: /searching/internal/items/storage/postgres/integrations/models.go
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
	"os"
	"search_service/genproto/searching"
	"search_service/internal/items/models"
	"strconv"

	"go.uber.org/zap"
)

func (s *SearchIntegration) Search(ctx context.Context, req *searching.QueryRequest) (*searching.SearchResponse, error) {
	encodedQuery := url.QueryEscape(req.Query)
	start := (req.PageNumber-1)*req.PageSize + 1

	searchURL := s.cfg.CustomSearch()
	searchURL += "&q=" + encodedQuery
	searchURL += fmt.Sprintf("&start=%d", start)
	searchURL += fmt.Sprintf("&num=%d", req.PageSize)

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, searchURL, nil)
	if err != nil {
		s.logger.Error("Creating new request failed", zap.Error(err))
		return nil, err
	}

	httpRes, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		s.logger.Error("Sending request failed", zap.Error(err))
		return nil, err
	}
	defer httpRes.Body.Close()

	if httpRes.StatusCode == 429 {
		s.logger.Warn("Google API reached the limit for today!", zap.Int("status_code", httpRes.StatusCode))
		return nil, fmt.Errorf("google api reached the limit for today")
	}

	if httpRes.StatusCode != http.StatusOK {
		s.logger.Warn("Request failed with status", zap.Int("status_code", httpRes.StatusCode))
		return nil, fmt.Errorf("unexpected status code: %d", httpRes.StatusCode)
	}
	body, err := io.ReadAll(httpRes.Body)
	if err != nil {
		s.logger.Error("Reading from HTTP body failed", zap.Error(err))
		return nil, err
	// fmt.Println(string(body))
	}

	searchRes := models.SearchElementsAPIResponse{}
	if err := json.Unmarshal(body, &searchRes); err != nil {
		s.logger.Error("Unmarshalling body failed", zap.Error(err))
		return nil, err
	}

	var results []*searching.Result
	for _, item := range searchRes.Items {
		// path, err := s.downloadAndCacheFavicon(item.DisplayLink)
		// if err != nil {
		// 	s.logger.Error(path, zap.Error(err))
		// }

		// fmt.Println()
		results = append(results, &searching.Result{
			Title:      item.Title,
			DirectLink: item.Link,
			Content:    item.Snippet,
			FavIconUrl: "https://" + item.DisplayLink + "/favicon.ico",
			PrimaryImageUrl: func() string {
				if len(item.PageMap.CseImage) == 0 {
					return ""
				}
				return item.PageMap.CseImage[0].Src
			}(),
			DisplayLink: item.DisplayLink,
			Thumbnails: func() []*searching.SearchThumbnail {
				thumbnailResults := []*searching.SearchThumbnail{}
				for _, thumbnail := range item.PageMap.CseThumbnail {
					// Convert string width/height to int64
					width, _ := strconv.ParseInt(thumbnail.Width, 10, 64)
					height, _ := strconv.ParseInt(thumbnail.Height, 10, 64)

					thumbnailResults = append(thumbnailResults, &searching.SearchThumbnail{
						Src:    thumbnail.Src,
						Width:  width,
						Height: height,
					})
				}
				return thumbnailResults
			}(),
		})
	}
	totalItems, err := strconv.Atoi(searchRes.SearchInformation.TotalResults)
	if err != nil {
		s.logger.Error("Failed to convert total items string into int using Atoi", zap.Error(err))
	}
	return &searching.SearchResponse{
		TotalItems: int32(totalItems),
		Results:    results}, nil
}

// func ProcessFavIcon()

func (s *SearchIntegration) downloadAndCacheFavicon(url string) (string, error) {
	// Path to save favicon
	faviconPath := "./static/favicons/" + url + ".ico"

	// Ensure the directory exists
	err := os.MkdirAll("./static/favicons", os.ModePerm) // Create the directory if it doesn't exist
	if err != nil {
		s.logger.Error("Failed to create directory", zap.Error(err))
		return "", err
	}

	// Check if the favicon already exists
	if _, err := os.Stat(faviconPath); !os.IsNotExist(err) {
		return faviconPath, nil // Return the cached favicon path if it exists
	}

	// Download the favicon if not cached
	resp, err := http.Get("https://" + url + "/favicon.ico")
	if err != nil {
		s.logger.Error("Failed to fetch favicon", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	// Save the favicon image to disk
	outFile, err := os.Create(faviconPath)
	if err != nil {
		s.logger.Error("Failed to create favicon file", zap.Error(err))
		return "", err
	}
	defer outFile.Close()

	// Copy the body (favicon data) into the file
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		s.logger.Error("Failed to save favicon file", zap.Error(err))
		return "", err
	}

	return faviconPath, nil
}
