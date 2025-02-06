package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/feeds"
	"go.uber.org/zap"
)

// Feed Items Handlers

// @Summary Create a new feed item
// @Description Create a new feed item
// @Tags Feed Items
// @Accept json
// @Produce json
// @Param item body feeds.CreateFeedItem true "Create Feed Item"
// @Success 201 {object} feeds.FeedItem
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/items/ [post]
func (h *AdminHandler) CreateFeedItem(c *gin.Context) {
	var req feeds.CreateFeedItem
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding feed item request", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := h.feedItems.CreateFeedItem(c, &req)
	if err != nil {
		h.logger.Error("Error creating feed item", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, item)
}

// @Summary Get feed item by ID
// @Description Retrieve a specific feed item by its ID
// @Tags Feed Items
// @Produce json
// @Param id path int32 true "Feed Item ID"
// @Success 200 {object} feeds.FeedItem
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/items/{id} [get]
func (h *AdminHandler) GetFeedItemByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req := &feeds.FeedItemRequest{FeedItemId: int64(id)}
	item, err := h.feedItems.GetFeedItem(c, req)
	if err != nil {
		h.logger.Error("Error retrieving feed item", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, item)
}

// @Summary Update an existing feed item
// @Description Update an existing feed item by its ID
// @Tags Feed Items
// @Accept json
// @Produce json
// @Param id path int32 true "Feed Item ID"
// @Param item body feeds.FeedItemsUpdate true "Update Feed Item"
// @Success 200 {object} feeds.FeedItem
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/items/{id} [put]
func (h *AdminHandler) UpdateFeedItem(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req feeds.FeedItemsUpdate
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding feed item update request", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.FeedId = int64(id)
	updatedItem, err := h.feedItems.UpdateFeedItem(c, &req)
	if err != nil {
		h.logger.Error("Error updating feed item", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, updatedItem)
}

// @Summary Get all feed items by feed id
// @Description Retrieve all feed items belong to feed with pagination support
// @Tags Feed Items
// @Produce json
// @Param limit query int32 true "Limit number of items"
// @Param page query int32 true "Page number"
// @Param feed_id path int32 true "feed id"
// @Success 200 {object} feeds.FeedItemsResponse
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/items/all/{feed_id} [get]
func (h *AdminHandler) GetAllFeedItemsByFeed(c *gin.Context) {
	// Get query parameters for pagination (default to 10 items per page, page 1)
	limit := getQueryParamInt(c, "limit", 10)
	page := getQueryParamInt(c, "page", 1)
	feedId, err := parseID(c, "feed_id")
	if err != nil {
		h.logger.Error("Faild to get feed id from path", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the request to get all feed items
	req := &feeds.GetFeedItemsByFeed{
		Limit:  int64(limit),
		Page:   int64(page),
		FeedId: int64(feedId),
	}

	// Call the corresponding gRPC handler method
	resp, err := h.feedItems.GetAllFeedItemsByFeed(c, req)
	if err != nil {
		h.logger.Error("Error retrieving all feed items", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Return the response to the client
	c.IndentedJSON(http.StatusOK, resp)
}

// @Summary Delete a feed item
// @Description Delete a feed item by its ID
// @Tags Feed Items
// @Produce json
// @Param id path string true "Feed Item ID"
// @Success 200 {string} string "Feed item deleted successfully"
// @Failure 400 {object} map[string]string "Invalid ID"
// @Failure 404 {object} map[string]string "Feed item not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /v1/admin/feeds/items/{id} [delete]
func (h *AdminHandler) DeleteFeedItem(c *gin.Context) {
	// Retrieve the feed item ID from the URL path
	idStr := c.Param("id")

	// Validate the ID (you can add more validation if needed)
	if idStr == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Convert the string ID to int32
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	// Create the request for deleting the feed item
	req := &feeds.FeedItemRequest{
		FeedItemId: int64(id), // Convert the id to int32
	}

	// Call the corresponding gRPC handler method
	_, err = h.feedItems.DeleteFeedItem(c, req)
	if err != nil {
		// Handle error - item not found or server error
		if err.Error() == "not found" {
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Feed item not found"})
		} else {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Successfully deleted the item, return a success message
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Feed item deleted successfully"})
}

func parseID(c *gin.Context, p string) (int32, error) {
	idStr := c.Param(p) // Retrieves the parameter from the URL path
	if idStr == "" {
		return 0, fmt.Errorf("parameter %s is required", p)
	}

	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("invalid parameter %s: %v", p, err)
	}

	return int32(id), nil
}

func getQueryParamInt(c *gin.Context, param string, defaultValue int32) int32 {
	valueStr := c.Query(param) // Retrieves the query parameter
	if valueStr == "" {
		return defaultValue
	}

	id, err := strconv.ParseInt(valueStr, 10, 64)

	if err != nil {
		// Log or handle the error if necessary, but return the default value
		return defaultValue
	}

	return int32(id)
}
