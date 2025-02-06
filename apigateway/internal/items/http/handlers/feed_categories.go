package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/feeds"
	"go.uber.org/zap"
)

// AdminHandler handles admin-related operations for feeds, feed items, and categories
type AdminHandler struct {
	feeds     feeds.FeedsServiceClient
	feedItems feeds.FeedItemsServiceClient
	feedCat   feeds.CategoriesServiceClient
	logger    *zap.Logger
}

// NewAdminHandler creates a new instance of AdminHandler
func NewAdminHandler(
	feeds feeds.FeedsServiceClient,
	feedItems feeds.FeedItemsServiceClient,
	feedCat feeds.CategoriesServiceClient,
	logger *zap.Logger,
) *AdminHandler {
	return &AdminHandler{
		feeds:     feeds,
		feedItems: feedItems,
		feedCat:   feedCat,
		logger:    logger,
	}
}

// Feed Category Handlers

// @Summary Serving feed category icon
// @Description Serving feed category icon by its name and extension
// @Tags Feed Categories
// @Accept json
// @Produce json
// @Param icon_name query string false "icon uuid with extension" default(.png)
// @Success 201 {object} feeds.FeedCategory
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/categories/icon [get]
func (h *AdminHandler) ServeFeedCategoryIcon(c *gin.Context) {
	imageName := c.DefaultQuery("icon_name", "")
	if imageName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image name is required"})
		return
	}

	imageDir := "./images/feedCategories/icons"
	imagePath := fmt.Sprintf("%s/%s", imageDir, imageName)

	if _, err := os.Stat(imagePath); err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to access the file"})
		}
		return
	}

	c.File(imagePath)
}

// @Summary Create a new feed category
// @Description Create a new feed category with name and icon URL
// @Tags Feed Categories
// @Accept json
// @Produce json
// @Param category body feeds.CreateFeedCategoryRequest true "Create Feed Category"
// @Success 201 {object} feeds.FeedCategory
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/categories/ [post]
func (h *AdminHandler) CreateFeedCategory(c *gin.Context) {
	var req feeds.CreateFeedCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding feed category request", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	category, err := h.feedCat.CreateFeedCategory(c, &req)
	if err != nil {
		h.logger.Error("Error creating feed category", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusCreated, category)
}

// @Summary Get feed category by ID
// @Description Retrieve a specific feed category by its ID
// @Tags Feed Categories
// @Produce json
// @Param id path int64 true "Feed Category ID"
// @Success 200 {object} feeds.FeedCategory
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/categories/{id} [get]
func (h *AdminHandler) GetFeedCategoryByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	req := &feeds.GetFeedCategoryByIDRequest{Id: int64(id)}
	category, err := h.feedCat.GetFeedCategoryByID(c, req)
	if err != nil {
		h.logger.Error("Error retrieving feed category", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}

// @Summary Update feed category
// @Description Update an existing feed category
// @Tags Feed Categories
// @Accept json
// @Produce json
// @Param id path int64 true "Feed Category ID"
// @Param category body feeds.FeedCategory true "Update Feed Category"
// @Success 200 {object} feeds.FeedCategory
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/categories/{id} [put]
func (h *AdminHandler) UpdateFeedCategory(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	var req feeds.FeedCategory
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Error binding feed category update request", zap.Error(err))
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.Id = id
	category, err := h.feedCat.UpdateFeedCategory(c, &req)
	if err != nil {
		h.logger.Error("Error updating feed category", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, category)
}

// @Summary Delete feed category
// @Description Delete a feed category by its ID
// @Tags Feed Categories
// @Param id path int64 true "Feed Category ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/categories/{id} [delete]
func (h *AdminHandler) DeleteFeedCategory(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}

	req := &feeds.DeleteFeedCategoryRequest{Id: int64(id)}
	_, err = h.feedCat.DeleteFeedCategory(c, req)
	if err != nil {
		h.logger.Error("Error deleting feed category", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

// @Summary Get all feed categories
// @Description Retrieve all feed categories with pagination
// @Tags Feed Categories
// @Produce json
// @Param limit query int false "Limit of categories per page" default(10)
// @Param page query int false "Page number" default(1)
// @Param lang query string false "lang" default(uz)
// @Success 200 {object} feeds.FeedCategoriesResponse
// @Failure 500 {object} map[string]string
// @Router /v1/admin/feeds/categories/all [get]
func (h *AdminHandler) GetAllFeedCategories(c *gin.Context) {
	limit := getQueryParamInt(c, "limit", 10)
	page := getQueryParamInt(c, "page", 1)
	lang := c.Query("lang")

	req := &feeds.GetAllFeedCategoriesRequest{
		Limit: limit,
		Page:  page,
		Lang:  lang,
	}

	categories, err := h.feedCat.GetAllFeedCategories(c, req)
	if err != nil {
		h.logger.Error("Error retrieving all feed categories", zap.Error(err))
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, categories)
}
