// /*
//   - @Author: javohir-a abdusamatovjavohir@gmail.com
//   - @Date: 2024-11-22 23:02:16
//   - @LastEditors: javohir-a abdusamatovjavohir@gmail.com
//   - @LastEditTime: 2024-11-22 23:26:02
//   - @FilePath: /pro-sphere-backend/admin/internal/items/http/handler/feedFetcher.go
//   - @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
//     */
package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
// 	"go.uber.org/zap"
// )

// type (
// 	FeedsFetcherHandler struct {
// 		service *service.FeedItemsService
// 		logger  *zap.Logger
// 	}
// )

// func NewFeedsFetcherHandler(service *service.FeedItemsService, logger *zap.Logger) *FeedsFetcherHandler {
// 	return &FeedsFetcherHandler{
// 		service: service,
// 		logger:  logger,
// 	}
// }

// // @Summary This endpoint is used to test the feed fetcher
// // @Description Fetcher
// // @Tags Feed Fetcher (test)
// // @Accept json
// // @Produce json
// // @Param id path int true "Feed ID"
// // @Success 200 {object} []models.FeedItemResponse
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feedTest/{id} [get]
// func (f *FeedsFetcherHandler) FetchByID(c *gin.Context) {
// 	// Parse feed ID from URL
// 	feedIDStr := c.Param("id")
// 	feedID, err := strconv.Atoi(feedIDStr)
// 	if err != nil {
// 		f.logger.Error("Invalid feed ID", zap.Error(err))
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid feed ID"})
// 		return
// 	}

// 	// Fetch feed items from the service
// 	feedItems, err := f.service.FetchFeedItems(c.Request.Context(), feedID)
// 	if err != nil {
// 		f.logger.Error("Error fetching feed items", zap.Error(err))
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch feed items"})
// 		return
// 	}

// 	// Return the feed items as JSON response
// 	c.JSON(http.StatusOK, feedItems)
// }
