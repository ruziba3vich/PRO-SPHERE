/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-27 03:58:18
 * @FilePath: /admin/internal/items/service/feedCategories.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"context"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/internal/items/repo"
	"go.uber.org/zap"
)

type (
	CategoriesService struct {
		storage repo.FeedCategoriesRepository
		logger  *zap.Logger
	}
)

func NewCategoriesService(storage repo.FeedCategoriesRepository, logger *zap.Logger) *CategoriesService {
	return &CategoriesService{
		storage: storage,
		logger:  logger,
	}
}

func (c *CategoriesService) CreateFeedCategory(ctx context.Context, iconURL string, translations []models.FeedCategoryTranslation) (*models.FeedCategory, error) {
	c.logger.Info("Creating a new feed category",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Any("translations", translations),
		zap.String("icon_url", iconURL),
	)
	imageDir := "./images/feedCategories/icons"
	err := os.MkdirAll(imageDir, os.ModePerm)
	if err != nil {
		c.logger.Error("Failed to create image directory", zap.String("directory", imageDir), zap.Error(err))
		return nil, err
	}

	// Fetch image and validate its type to determine the correct file extension
	imageID, err := SaveAndValidateImage(iconURL, imageDir, c.logger)
	if err != nil {
		c.logger.Error("Failed to save and validate image", zap.Error(err))
		return nil, err
	}

	// Proceed with creating the feed category in the database
	return c.storage.CreateFeedCategory(ctx, &models.FeedCategory{IconURL: iconURL, IconID: imageID, Translations: translations})
}

func (c *CategoriesService) GetFeedCategoryByID(ctx context.Context, id int64, lang string) (*models.FeedCategory, error) {
	c.logger.Info("Fetching feed category by ID",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Int64("id", id),
	)

	return c.storage.GetFeedCategoryByID(ctx, id, lang)
}

func (c *CategoriesService) UpdateFeedCategory(ctx context.Context, cat *models.FeedCategory) (*models.FeedCategory, error) {
	c.logger.Info("Updating feed category",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Int64("id", cat.ID),
		zap.String("name", cat.Translations[0].Name),
		zap.String("icon_url", cat.IconURL),
	)

	// Fetch the existing category's icon info
	iconInfo, err := c.storage.GetFeedCategoryIconInfo(ctx, int(cat.ID))
	if err != nil {
		c.logger.Error("Failed to fetch existing feed category icon info", zap.Error(err))
		return nil, err
	}

	// Check if the icon URL has changed
	if iconInfo.IconUrl != cat.IconURL {
		c.logger.Info("Icon URL has changed; downloading new image",
			zap.String("old_url", iconInfo.IconUrl),
			zap.String("new_url", cat.IconURL),
		)

		// Save and validate the new image
		imageDir := "./images/feedCategories/icons"
		newImageID, err := SaveAndValidateImage(cat.IconURL, imageDir, c.logger)
		if err != nil {
			c.logger.Error("Failed to save and validate new image", zap.Error(err))
			return nil, err
		}
		// Remove the out-dated image from the storage
		iconPath := "./images/feedCategories/icons/" + iconInfo.IconID
		go RemoveImage(iconPath, c.logger)

		// Update the image ID in the category object
		cat.IconID = newImageID
	} else {
		// If the URL hasn't changed, retain the existing IconID
		cat.IconID = iconInfo.IconID
	}

	// Update the category in the database
	updatedCategory, err := c.storage.UpdateFeedCategory(ctx, cat)
	if err != nil {
		c.logger.Error("Failed to update feed category", zap.Error(err))
		return nil, err
	}

	c.logger.Info("Feed category successfully updated", zap.Int64("id", updatedCategory.ID))
	return updatedCategory, nil
}

func (c *CategoriesService) DeleteFeedCategory(ctx context.Context, id int64) error {
	c.logger.Info("Deleting feed category by ID",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Int64("id", id),
	)

	// Step 1: Retrieve the icon information associated with the category
	iconInfo, err := c.storage.GetFeedCategoryIconInfo(ctx, int(id))
	if err != nil {
		c.logger.Error("Failed to fetch icon information for feed category",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return fmt.Errorf("failed to retrieve feed category icon info: %w", err)
	}

	// Step 2: Delete the feed category from the database
	if err := c.storage.DeleteFeedCategory(ctx, id); err != nil {
		c.logger.Error("Failed to delete feed category from database",
			zap.Int64("id", id),
			zap.Error(err),
		)
		return fmt.Errorf("failed to delete feed category: %w", err)
	}
	iconPath := "./images/feedCategories/icons/" + iconInfo.IconID
	// Step 3: Attempt to delete the icon image from storage
	go RemoveImage(iconPath, c.logger)

	return nil
}

func RemoveImage(iconPath string, logger *zap.Logger) {
	if iconPath != "" {
		err := os.Remove(iconPath)
		if err != nil {
			logger.Warn("Failed to delete icon image file",
				zap.String("path", iconPath),
				zap.Error(err),
			)
			// Log and proceed, as this should not block the deletion of the category
		} else {
			logger.Info("Successfully deleted icon image file",
				zap.String("path", iconPath),
			)
		}
	}
}

func (c *CategoriesService) GetAllFeedCategories(ctx context.Context, limit, page int, lang string) ([]*models.FeedCategory, error) {
	c.logger.Info("Fetching all feed categories",
		zap.String("time", time.Now().Format(time.RFC3339)),
		zap.Int("limit", limit),
		zap.Int("page", page),
		zap.String("lang", lang),
	)

	cats, err := c.storage.GetAllFeedCategories(ctx, page, limit, lang)
	if err != nil {
		c.logger.Error("Failed to getch fata from storage layer", zap.Error(err))
		return nil, err
	}

	return cats, err
}

// saveAndValidateImage downloads an image, validates its type, and saves it with a proper file extension.
func SaveAndValidateImage(imgURL, imageDir string, logger *zap.Logger) (string, error) {

	req, err := http.NewRequest(http.MethodGet, imgURL, nil)
	if err != nil {
		logger.Error("Failed to create HTTP request for image", zap.Error(err))
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		logger.Error("Failed to fetch image", zap.String("url", imgURL), zap.Error(err))
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		logger.Error("Failed to download image", zap.String("url", imgURL), zap.Int("status", res.StatusCode))
		return "", fmt.Errorf("failed to fetch image: status code %d", res.StatusCode)
	}

	// Validate the Content-Type header
	contentType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "image/") {
		logger.Error("Invalid content type", zap.String("content_type", contentType))
		return "", fmt.Errorf("invalid content type: %s", contentType)
	}

	// Determine file extension based on content type
	exts, err := mime.ExtensionsByType(contentType)
	if err != nil || len(exts) == 0 {
		logger.Error("Failed to determine file extension", zap.String("content_type", contentType))
		return "", fmt.Errorf("unknown file extension for content type: %s", contentType)
	}
	fileExtension := exts[0]

	// Generate a unique image ID with the correct extension
	imageID := uuid.NewString() + fileExtension

	// Ensure the directory exists
	err = os.MkdirAll(imageDir, os.ModePerm)
	if err != nil {
		logger.Error("Failed to create image directory", zap.String("directory", imageDir), zap.Error(err))
		return "", err
	}

	// Save the image to disk
	filePath := filepath.Join(imageDir, imageID)
	file, err := os.Create(filePath)
	if err != nil {
		logger.Error("Failed to create file for image", zap.String("path", filePath), zap.Error(err))
		return "", err
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		logger.Error("Failed to save image to file", zap.String("path", filePath), zap.Error(err))
		return "", err
	}

	logger.Info("Image successfully saved", zap.String("file_path", filePath))
	return imageID, nil
}
