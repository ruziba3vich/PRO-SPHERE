// /*
//   - @Author: javohir-a abdusamatovjavohir@gmail.com
//   - @Date: 2024-10-08 23:48:14
//   - @LastEditors: javohir-a abdusamatovjavohir@gmail.com
//   - @LastEditTime: 2024-11-23 00:37:06
//   - @FilePath: /sphere_posts/internal/items/http/handler/feedHandler.go
//   - @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
//     */
package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/projects/pro-sphere-backend/admin/genproto/genproto/feeds"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
// 	"go.uber.org/zap"
// )

// type (
// 	FeedsHandler struct {
// 		service *service.FeedsService
// 		logger  *zap.Logger
// 	}
// )

// func NewFeedsHandler(service *service.FeedsService, logger *zap.Logger) *FeedsHandler {
// 	return &FeedsHandler{
// 		service: service,
// 		logger:  logger,
// 	}
// }

// // @Summary Create a new feed
// // @Description Create a new feed
// // @Tags Feeds
// // @Accept json
// // @Produce json
// // @Param feed body models.Feed true "feed"
// // @Success 201 {object} feeds.FeedResponse
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feeds [post]
// func (h *FeedsHandler) CreateFeed(c *gin.Context) {
// 	var req feeds.Feed
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		h.logger.Error("Error occurred while binding feed data",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response, err := h.service.CreateFeed(c, &req)
// 	if err != nil {
// 		h.logger.Error("Error occurred while creating feed",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response)
// }

// // @Summary Get feed by ID
// // @Description Get a feed by its ID
// // @Tags Feeds
// // @Produce json
// // @Param id path int true "feed_id"
// // @Success 200 {object} feeds.FeedResponse
// // @Failure 404 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Failure 204 {object} map[string]string
// // @Router /feeds/{id} [get]
// func (h *FeedsHandler) GetFeed(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		h.logger.Error("Error converting ID",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	response, err := h.service.GetFeed(c, id)
// 	if response == nil {
// 		h.logger.Info("No feed found")
// 		c.Status(http.StatusNoContent)
// 		return
// 	}
// 	if err != nil {
// 		h.logger.Error("Error occurred while fetching feed",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Feed not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // @Summary Get all feeds
// // @Description Get all feeds with pagination
// // @Tags Feeds
// // @Produce json
// // @Param limit query int false "limit"
// // @Param page query int false "page"
// // @Success 200 {array} feeds.FeedResponse
// // @Failure 500 {object} map[string]string
// // @Failure 204 {object} map[string]string
// // @Router /feeds [get]
// func (h *FeedsHandler) GetAllFeeds(c *gin.Context) {
// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

// 	response, err := h.service.GetAllFeeds(c, limit, page)
// 	if response == nil {
// 		h.logger.Info("No feeds found")
// 		c.Status(http.StatusNoContent)
// 		return
// 	}
// 	if err != nil {
// 		h.logger.Error("Error occurred while fetching all feeds",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // @Summary Update an existing feed
// // @Description Update a feed by its ID
// // @Tags Feeds
// // @Accept json
// // @Produce json
// // @Param feed body feeds.UpdateFeed true "feed"
// // @Success 200 {object} feeds.FeedResponse
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feeds [put]
// func (h *FeedsHandler) UpdateFeed(c *gin.Context) {
// 	var req feeds.UpdateFeed
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		h.logger.Error("Error occurred while binding update data",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response, err := h.service.UpdateFeed(c, &req)
// 	if err != nil {
// 		h.logger.Error("Error occurred while updating feed",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // @Summary Delete a feed by ID
// // @Description Delete a feed by its ID
// // @Tags Feeds
// // @Produce json
// // @Param id path int true "feed_id"
// // @Success 200 {object} map[string]string
// // @Failure 404 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feeds/{id} [delete]
// func (h *FeedsHandler) DeleteFeed(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		h.logger.Error("Error converting ID",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	_, err = h.service.DeleteFeed(c, id)
// 	if err != nil {
// 		h.logger.Error("Error occurred while deleting feed",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Feed not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Feed successfully deleted"})
// }
