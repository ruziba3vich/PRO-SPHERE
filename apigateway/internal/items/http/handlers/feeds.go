package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/feeds"
	"go.uber.org/zap"
)

// Feed Handlers

// // @Summary Fetch feed items by ID
// // @Description Retrieve a specific feed's items by its ID
// // @Tags Feeds
// // @Produce json
// // @Param id path int32 true "Feed ID"
// // @Success 200 {object} feeds.FetchFeedItemsResponse
// // @Failure 400 {object} map[string]string
// // @Failure 404 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /v1/admin/feeds/feed/fetch/{id} [get]
// func (h *AdminHandler) FetchFeedItems(c *gin.Context) {
// 	id, err := parseID(c, "id")
// 	if err != nil {
// 		return
// 	}

// 	req := &feeds.FetchFeedItemsRequest{FeedId: int64(id)}

// 	res, err := h.feeds.FetchFeedItems(c, req)
// 	if err != nil {
// 		h.logger.Error("Error fetching feed items", zap.Error(err))
// 		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK, res)
// }

// @Summary Create a new feed
// @Description Create a new feed with detailed information
// @Tags Feeds
// @Accept json
// @Produce json
// @Param feed body feeds.CreateFeedRequest true "Create Feed"
// @Success 201 {object} feeds.Feed
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/feed/ [post]
func (h *AdminHandler) CreateFeed(c *gin.Context) {
	var req feeds.CreateFeedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding feed request", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	feed, err := h.feeds.CreateFeed(c, &req)
	if err != nil {
		h.logger.Error("Error creating feed", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, feed)
}

// @Summary Get feed by ID
// @Description Retrieve a specific feed by its ID
// @Tags Feeds
// @Produce json
// @Param id path int32 true "Feed ID"
// @Success 200 {object} feeds.Feed
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/feed/{id} [get]
func (h *AdminHandler) GetFeedByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	req := &feeds.FeedRequest{Id: int64(id)}
	feed, err := h.feeds.GetFeed(c, req)
	if err != nil {
		h.logger.Error("Error retrieving feed", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, feed)
}

// @Summary Update feed
// @Description Update an existing feed
// @Tags Feeds
// @Accept json
// @Produce json
// @Param id path int32 true "Feed ID"
// @Param feed body feeds.UpdateFeedRequest true "Update Feed"
// @Success 200 {object} feeds.Feed
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/feed/{id} [put]
func (h *AdminHandler) UpdateFeed(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req feeds.UpdateFeedRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding feed update request", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = int64(id)
	feed, err := h.feeds.UpdateFeed(c, &req)
	if err != nil {
		h.logger.Error("Error updating feed", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, feed)
}

// @Summary Delete feed
// @Description Delete a feed by its ID
// @Tags Feeds
// @Param id path int32 true "Feed ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/feed/{id} [delete]
func (h *AdminHandler) DeleteFeed(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	req := &feeds.FeedRequest{Id: int64(id)}
	_, err = h.feeds.DeleteFeed(c, req)
	if err != nil {
		h.logger.Error("Error deleting feed", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get all feeds
// @Description Retrieve all feeds with pagination
// @Tags Feeds
// @Produce json
// @Param limit query int false "Limit of feeds per page" default(10)
// @Param page query int false "Page number" default(1)
// @Param priority query bool false "Priority config" default(true)
// @Param lang query string false "Lang" default(uz)
// @Success 200 {object} feeds.FeedsResponse
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/feed/all [get]
func (h *AdminHandler) GetAllFeeds(c *gin.Context) {
	// Parse query parameters with defaults
	limit := getQueryParamInt(c, "limit", 10)
	page := getQueryParamInt(c, "page", 1)
	lang := c.Query("lang")

	// Parse `priority` query parameter
	priority, err := strconv.ParseBool(c.DefaultQuery("priority", "true"))
	if err != nil {
		// If the query parameter is invalid, log and return a bad request error
		h.logger.Error("Invalid priority query parameter", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid priority value. Must be true or false."})
		return
	}

	req := &feeds.GetAllFeedsRequest{
		Limit:    limit,
		Page:     page,
		Priority: priority,
		Lang:     lang,
	}

	// Call the service layer to fetch feeds
	feeds, err := h.feeds.GetAllFeeds(c, req)
	if err != nil {
		h.logger.Error("Error retrieving all feeds", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the result
	c.IndentedJSON(http.StatusOK, feeds)
}

// AddFeedContent handles adding feed content.
// @Summary Add Feed Content
// @Description Adds new content to a feed.
// @Tags Feed
// @Accept json
// @Produce json
// @Param feedContent body feeds.FeedContent true "Feed Content"
// @Success 200 {object} feeds.FeedContent
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/content [post]
func (h *AdminHandler) AddFeedContent(c *gin.Context) {
	var req feeds.FeedContent
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newContent, err := h.feeds.AddFeedContent(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to add feed content", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add feed content"})
		return
	}

	c.JSON(http.StatusOK, newContent)
}

// GetFeedContent handles getting feed content.
// @Summary Get Feed Content
// @Description Gets exiting content by its id.
// @Tags Feed
// @Accept json
// @Produce json
// @Param id query string true "Feed Content"
// @Success 200 {object} feeds.FeedContent
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/content [get]
func (h *AdminHandler) GetFeedContent(c *gin.Context) {
	var req feeds.FeedRequest
	id := getQueryParamInt(c, "id", 0)
	req.Id = int64(id)

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Invalid request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newContent, err := h.feeds.GetFeedContent(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to get feed content", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get feed content"})
		return
	}

	c.JSON(http.StatusFound, newContent)
}

type UpdateFContent struct {
	Link       string `json:"link"`
	Lang       string `json:"lang"`
	CategoryId int    `json:"category_id"`
}

// UpdateFeedContent handles updating feed content.
// @Summary Update Feed Content
// @Description Updates new content to a feed.
// @Tags Feed
// @Accept json
// @Produce json
// @Param id path int true "feed content id"
// @Param feedContent body UpdateFContent true "Feed Content"
// @Success 200 {object} feeds.FeedContent
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/content/{id} [put]
func (h *AdminHandler) UpdateFeedContent(c *gin.Context) {
	// Get the ID from the path and convert it to int64
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.logger.Error("Invalid ID parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}
	var gRPCReq UpdateFContent

	if err := c.ShouldBindJSON(&gRPCReq); err != nil {
		h.logger.Error("Invalid request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedContent, err := h.feeds.UpdateFeedContent(c.Request.Context(), &feeds.FeedContent{
		Id:         id,
		Link:       gRPCReq.Link,
		Lang:       gRPCReq.Lang,
		CategoryId: int64(gRPCReq.CategoryId),
	})

	if err != nil {
		h.logger.Error("Failed to update feed content", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update feed content"})
		return
	}

	c.JSON(http.StatusOK, updatedContent)
}

// DeleteFeedContent handles deleting feed content.
// @Summary Delete Feed Content
// @Description Deletes exiting content by its id.
// @Tags Feed
// @Accept json
// @Produce json
// @Param id path int true "Feed Content id"
// @Success 200 {object} feeds.FeedContent
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/content/{id} [delete]
func (h *AdminHandler) DeleteFeedContent(c *gin.Context) {
	// Get the ID from the path and convert it to int64
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		h.logger.Error("Invalid ID parameter", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID parameter"})
		return
	}

	// Create the request object with the ID
	req := feeds.FeedRequest{Id: id}

	// Call the service layer to delete the feed content
	newContent, err := h.feeds.DeleteFeedContent(c.Request.Context(), &req)
	if err != nil {
		h.logger.Error("Failed to delete feed content", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete feed content"})
		return
	}

	// Return the deleted content
	c.JSON(http.StatusOK, newContent)
}

// GetAllFeedContent handles getting feed contents by feed Id.
// @Summary Get Feed Contents
// @Description Gets exiting contents belongs to Feed.
// @Tags Feed
// @Accept json
// @Produce json
// @Param feedId query string true "Feed ID"
// @Param categoryId query string false "Feed Category ID"c
// @Param lang query string false "Feed content language"
// @Success 200 {object} feeds.FeedContentsRespose
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/content/all [get]
func (h *AdminHandler) GetAllFeedContent(c *gin.Context) {
	// Parse query parameters
	feedID := getQueryParamInt(c, "feedId", 0)
	categoryId := getQueryParamInt(c, "categoryId", 0)
	lang := c.DefaultQuery("lang", "")

	// Validate `feedId`
	if feedID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "feedId is a required parameter",
		})
		return
	}

	// Call the storage layer
	contents, err := h.feeds.GetAllFeedContent(c.Request.Context(), &feeds.GetFeedContents{
		FeedId:     int64(feedID),
		CategoryId: int64(categoryId),
		Lang:       lang,
	})
	if err != nil {
		h.logger.Error("Failed to fetch feed contents", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch feed contents",
		})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"feedContents": contents,
	})
}
