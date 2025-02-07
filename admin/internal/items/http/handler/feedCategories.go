/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-01 00:56:30
 * @FilePath: /admin/internal/items/http/handler/feedCategories.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package handlers

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"github.com/lib/pq"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
// 	"go.uber.org/zap"
// )

// type (
// 	CategoriesHandler struct {
// 		service *service.CategoriesService
// 		logger  *zap.Logger
// 	}
// )

// func NewCategoriesHandler(service *service.CategoriesService, logger *zap.Logger) *CategoriesHandler {
// 	return &CategoriesHandler{
// 		service: service,
// 		logger:  logger,
// 	}
// }

// // @Summary Create a new category
// // @Description Create a new feed category
// // @Tags Categories
// // @Accept json
// // @Produce json
// // @Param category body models.FeedCategory true "category"
// // @Success 201 {object} models.FeedCategory
// // @Failure 400 {object} map[string]string
// // @Failure 409 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feedcategories [post]
// func (h *CategoriesHandler) CreateCategory(c *gin.Context) {
// 	var req models.FeedCategory
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		h.logger.Error("Error occurred while binding category data", zap.Error(err))
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
// 		return
// 	}

// 	response, err := h.service.CreateFeedCategory(c, req.Name, req.IconURL)
// 	if err != nil {
// 		if pqErr, ok := err.(*pq.Error); ok {
// 			if pqErr.Code == "23505" {
// 				h.logger.Error("Duplicate category name", zap.Error(err))
// 				c.JSON(http.StatusConflict, gin.H{"error": "Category name already exists"})
// 				return
// 			}
// 		}

// 		h.logger.Error("Error occurred while creating category", zap.Error(err))
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
// 		return
// 	}

// 	c.JSON(http.StatusCreated, response)
// }

// // @Summary Get category by ID
// // @Description Get a feed category by its ID
// // @Tags Categories
// // @Produce json
// // @Param id path int true "category_id"
// // @Success 200 {object} models.FeedCategory
// // @Failure 404 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Failure 204 {object} map[string]string
// // @Router /feedcategories/{id} [get]
// func (h *CategoriesHandler) GetCategory(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		h.logger.Error("Error converting ID",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	response, err := h.service.GetFeedCategoryByID(c, int64(id))
// 	if response == nil {
// 		h.logger.Info("No categories found")
// 		c.Status(http.StatusNoContent)
// 		return
// 	}

// 	if err != nil {
// 		h.logger.Error("Error occurred while fetching category",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // @Summary Get all categories
// // @Description Get all feed categories with pagination
// // @Tags Categories
// // @Produce json
// // @Param limit query int false "limit"
// // @Param page query int false "page"
// // @Success 200 {array} models.FeedCategory
// // @Failure 500 {object} map[string]string
// // @Failure 204 {object} map[string]string
// // @Router /feedcategories [get]
// func (h *CategoriesHandler) GetAllCategories(c *gin.Context) {
// 	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
// 	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

// 	response, err := h.service.GetAllFeedCategories(c, limit, page)
// 	if response == nil {
// 		h.logger.Info("No category found")
// 		c.Status(http.StatusNoContent)
// 		return
// 	}

// 	if err != nil {
// 		h.logger.Error("Error occurred while fetching all categories",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // @Summary Update an existing category
// // @Description Update a feed category by its ID
// // @Tags Categories
// // @Accept json
// // @Produce json
// // @Param category body models.FeedCategory true "category"
// // @Success 200 {object} models.FeedCategory
// // @Failure 400 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feedcategories/{id} [put]
// func (h *CategoriesHandler) UpdateCategory(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		h.logger.Error("Error converting ID",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	var req models.FeedCategory
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		h.logger.Error("Error occurred while binding update data",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	response, err := h.service.UpdateFeedCategory(c, int64(id), req.Name, req.IconURL)
// 	if err != nil {
// 		h.logger.Error("Error occurred while updating category",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// // @Summary Delete a category by ID
// // @Description Delete a feed category by its ID
// // @Tags Categories
// // @Produce json
// // @Param id path int true "category_id"
// // @Success 200 {object} map[string]string
// // @Failure 404 {object} map[string]string
// // @Failure 500 {object} map[string]string
// // @Router /feedcategories/{id} [delete]
// func (h *CategoriesHandler) DeleteCategory(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		h.logger.Error("Error converting ID",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
// 		return
// 	}

// 	err = h.service.DeleteFeedCategory(c, int64(id))
// 	if err != nil {
// 		h.logger.Error("Error occurred while deleting category",
// 			zap.Error(err),
// 		)
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Category successfully deleted"})
// }
