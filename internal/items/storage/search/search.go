package search

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/ruziba3vich/prosphere/internal/items/config"
)

type (
	SearchStorage struct {
		postgres *sql.DB
		logger   *log.Logger
		cfg      *config.Config
	}

	SearchRequest struct {
		ApiKey            string `json:"api_key"` // Capitalized field name
		Query             string `json:"query"`
		SearchDepth       string `json:"search_depth"`        // Capitalized field name
		IncludeAnswer     bool   `json:"include_answer"`      // Capitalized field name
		IncludeImages     bool   `json:"include_images"`      // Capitalized field name
		IncludeRawContent bool   `json:"include_raw_content"` // Capitalized field name
		MaxResults        int    `json:"max_results"`         // Capitalized field name
	}

	SearchResponse struct {
		Query        string   `json:"query"`
		Answer       string   `json:"answer"`
		Images       []string `json:"images"`
		Results      []Result `json:"results"`
		ResponseTime float32  `json:"response_time"`
	}

	Result struct {
		Title   string  `json:"title"`
		Url     string  `json:"url"`
		Content string  `json:"content"`
		Score   float32 `json:"score"`
	}
)

func NewSearchStorage(postgres *sql.DB, logger *log.Logger, config *config.Config) *SearchStorage {
	return &SearchStorage{
		postgres: postgres,
		logger:   logger,
		cfg:      config,
	}
}

func (s *SearchStorage) SearchElement(ctx context.Context, req *SearchRequest) (*SearchResponse, error) {
	s.fill(req)
	requestBody, err := json.Marshal(req)
	if err != nil {
		s.logger.Println("Error marshaling request data:", err)
		return nil, err
	}

	response, err := http.Post(s.cfg.Search.Api, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		s.logger.Println("Error sending POST request:", err)
		return nil, err
	}
	defer response.Body.Close()

	// pp.Println(req)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		s.logger.Println("Error reading response body:", err)
		return nil, err
	}
	// pp.Println(string(responseBody))

	var searchResponse SearchResponse
	err = json.Unmarshal(responseBody, &searchResponse)
	if err != nil {
		s.logger.Println("Error unmarshaling response data:", err)
		return nil, err
	}
	return &searchResponse, nil
}

func (s *SearchStorage) fill(req *SearchRequest) {
	req.ApiKey = s.cfg.Search.Token
	req.SearchDepth = "basic"
	req.IncludeAnswer = true
	req.IncludeImages = true
	req.IncludeRawContent = false
	req.MaxResults = 10
}
