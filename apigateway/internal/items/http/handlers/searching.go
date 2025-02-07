package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/searching"
	"go.uber.org/zap"
)

type (
	SearchHandler struct {
		client searching.SearchingServiceClient
		logger *zap.Logger
	}
)

func NewSearchHandler(client searching.SearchingServiceClient, logger *zap.Logger) *SearchHandler {
	return &SearchHandler{
		client: client,
		logger: logger,
	}
}

// @Summary Search elements
// @Description Search elements using Google Search Custom JSON API
// @Tags Searching
// @Accept  json
// @Produce  json
// @Param query query string true "Search query"
// @Param page_number query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} searching.SearchResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/search [get]
func (h *SearchHandler) Search(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	pageNumberStr := c.DefaultQuery("page_number", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "page_number must be a positive integer"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "page_size must be a positive integer"})
		return
	}

	req := &searching.QueryRequest{
		Query:      query,
		PageNumber: int32(pageNumber),
		PageSize:   int32(pageSize),
	}

	response, err := h.client.Search(c, req)
	if err != nil {
		h.logger.Error("Error occurred while searching elements", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

// @Summary Search videos on YouTube
// @Description Search videos using YouTube Data API
// @Tags Searching
// @Accept  json
// @Produce  json
// @Param query query string true "Search query"
// @Param max_results query int false "Maximum number of results"
// @Param next_page_token query string false "Token for fetching the next page of results"
// @Success 200 {object} searching.VideoSearchResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/search/youtube [get]
func (h *SearchHandler) SearchVideo(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	maxResultsStr := c.DefaultQuery("max_results", "10")
	maxResults, err := strconv.Atoi(maxResultsStr)
	if err != nil || maxResults < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "max_results must be a positive integer"})
		return
	}

	nextPageToken := c.Query("next_page_token")

	req := &searching.VideoSearchRequest{
		Query:         query,
		MaxResults:    int32(maxResults),
		NextPageToken: nextPageToken,
	}

	response, err := h.client.SearchVideos(c, req)
	if err != nil {
		h.logger.Error("Error occurred while searching videos", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}

// @Summary Search images
// @Description Search images using Google Custom Search API
// @Tags Searching
// @Accept  json
// @Produce  json
// @Param query query string true "Search query"
// @Param page_number query int false "Page number"
// @Param page_size query int false "Page size"
// @Success 200 {object} searching.ImageResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/search/images [get]
func (h *SearchHandler) SearchImages(c *gin.Context) {
	query := c.Query("query")
	if query == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Query parameter is required"})
		return
	}

	pageNumberStr := c.DefaultQuery("page_number", "1")
	pageSizeStr := c.DefaultQuery("page_size", "10")

	pageNumber, err := strconv.Atoi(pageNumberStr)
	if err != nil || pageNumber < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "page_number must be a positive integer"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "page_size must be a positive integer"})
		return
	}

	req := &searching.QueryRequest{
		Query:      query,
		PageNumber: int32(pageNumber),
		PageSize:   int32(pageSize),
	}

	response, err := h.client.SearchImages(c, req)
	if err != nil {
		h.logger.Error("Error occurred while searching images", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, response)
}
