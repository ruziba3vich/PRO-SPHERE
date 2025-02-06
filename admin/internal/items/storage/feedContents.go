package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"go.uber.org/zap"
)

func (f *FeedsStorage) AddFeedContent(ctx context.Context, content *models.FeedContent) (*models.FeedContent, error) {
	// Insert feed content
	query := `
		INSERT INTO feed_contents (feed_id, lang, link, category_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := f.db.QueryRowContext(ctx, query, content.FeedId, content.Lang, content.Link, content.CategoryId).Scan(&content.Id)
	if err != nil {
		f.logger.Error("Failed to insert feed content", zap.Error(err))
		return nil, err
	}

	// Cache the added feed content
	if err = f.cacheFeedContent(ctx, *content); err != nil {
		f.logger.Error("Failed to cache feed content after creation", zap.Error(err))
		return nil, err
	}

	f.logger.Info("Feed content added and cached successfully", zap.Int("content_id", content.Id))
	return content, nil
}

// GetFeedContent fetches feed content by ID from the database, using Redis cache if available.
func (f *FeedsStorage) GetFeedContent(ctx context.Context, id int) (*models.FeedContent, error) {
	// Check cache first
	cacheKey := fmt.Sprintf("feed_content:%d", id)
	cachedData, err := f.cache.Get(ctx, cacheKey).Result()
	if err == nil && cachedData != "" {
		var cachedContent models.FeedContent
		if err := json.Unmarshal([]byte(cachedData), &cachedContent); err == nil {
			f.logger.Info("Cache hit for feed content", zap.Int("id", id))
			return &cachedContent, nil
		}
		f.logger.Warn("Failed to unmarshal cached feed content", zap.Error(err))
	}

	// Cache miss: fetch from the database
	query := `
		SELECT id, feed_id, lang, link, category_id
		FROM feed_contents
		WHERE id = $1
	`
	content := models.FeedContent{}
	err = f.db.QueryRowContext(ctx, query, id).Scan(
		&content.Id, &content.FeedId, &content.Lang, &content.Link, &content.CategoryId,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			f.logger.Warn("Feed content not found", zap.Int("id", id))
			return nil, nil // or return an error like sql.ErrNoRows, depending on your needs
		}
		f.logger.Error("Failed to query feed content", zap.Error(err))
		return nil, err
	}

	// Cache the result for future requests
	if err = f.cacheFeedContent(ctx, content); err != nil {
		f.logger.Warn("Failed to cache feed content", zap.Error(err))
	}

	f.logger.Info("Feed content fetched successfully from database", zap.Int("id", id))
	return &content, nil
}

// UpdateFeedContent updates a specific feed content by its ID
func (f *FeedsStorage) UpdateFeedContent(ctx context.Context, id int, link *string, lang *string, categoryId *int) (*models.FeedContent, error) {
	// Start building the update query
	query := "UPDATE feed_contents SET "
	var args []interface{}
	argCount := 0

	// Add fields to update dynamically
	if link != nil {
		argCount++
		query += fmt.Sprintf("link = $%d", argCount)
		args = append(args, *link)

	}

	if lang != nil {
		if argCount > 0 {
			query += ", "
		}
		argCount++
		query += fmt.Sprintf("lang = $%d", argCount)
		args = append(args, *lang)
	}
	if categoryId != nil {
		if argCount > 0 {
			query += ", "
		}
		argCount++
		query += fmt.Sprintf("category_id = $%d", argCount)
		args = append(args, *categoryId)
	}

	// Add the WHERE clause and RETURNING
	query += fmt.Sprintf(" WHERE id = $%d RETURNING id, lang, link, category_id", argCount+1)
	args = append(args, id)
	// log.Println("Query: ", query)
	// Execute the query and fetch the updated values
	updatedContent := models.FeedContent{}
	err := f.db.QueryRowContext(ctx, query, args...).Scan(&updatedContent.Id, &updatedContent.Lang, &updatedContent.Link, &updatedContent.CategoryId)
	if err != nil {
		f.logger.Error("Failed to update feed content", zap.Error(err))
		return &updatedContent, err
	}

	// Update cache
	if err := f.cacheFeedContent(ctx, updatedContent); err != nil {
		f.logger.Error("Failed to update feed content in cache", zap.Error(err))
		return &updatedContent, err
	}

	f.logger.Info("Feed content updated successfully", zap.Int("id", id))
	return &updatedContent, nil
}

// DeleteFeedContent deletes a specific feed content by its ID
func (f *FeedsStorage) DeleteFeedContent(ctx context.Context, id int) error {
	// Define the delete query
	query := `DELETE FROM feed_contents WHERE id = $1`

	// Execute the delete query
	_, err := f.db.ExecContext(ctx, query, id)
	if err != nil {
		f.logger.Error("Failed to delete feed content", zap.Error(err))
		return err
	}

	key := fmt.Sprintf("feed_content:%d", id)

	if err := f.cache.Del(ctx, key).Err(); err != nil {
		f.logger.Error("Failed to delete feed content in redis cache", zap.Error(err))
	}

	return nil
}

func (f *FeedsStorage) GetAllFeedContent(ctx context.Context, feedId int, lang *string, categoryId *int) ([]models.FeedContent, error) {
	// Base query for fetching feed content based on feed_id
	query := `
		SELECT id, feed_id, lang, link, category_id
		FROM feed_contents
		WHERE feed_id = $1
	`

	var args []interface{}
	args = append(args, feedId) // Always filter by feed_id

	var argCount int

	// Add filter for lang if provided
	if lang != nil {
		argCount++
		query += fmt.Sprintf(" AND lang = $%d", argCount+1)
		args = append(args, *lang)
	}

	// Add filter for category_id if provided
	if categoryId != nil {
		argCount++
		query += fmt.Sprintf(" AND category_id = $%d", argCount+1)
		args = append(args, *categoryId)
	}

	// Execute the query
	rows, err := f.db.QueryContext(ctx, query, args...)
	if err != nil {
		f.logger.Error("Failed to query feed contents", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	// Scan the results into a slice of FeedContent
	var contents []models.FeedContent
	for rows.Next() {
		var content models.FeedContent
		err := rows.Scan(&content.Id, &content.FeedId, &content.Lang, &content.Link, &content.CategoryId)
		if err != nil {
			f.logger.Error("Failed to scan feed content", zap.Error(err))
			return nil, err
		}
		contents = append(contents, content)
	}

	// Check for row scanning errors
	if err := rows.Err(); err != nil {
		f.logger.Error("Failed to process feed content rows", zap.Error(err))
		return nil, err
	}

	// Return the fetched feed content
	return contents, nil
}

// cacheFeedContent caches the feed content in Redis
func (f *FeedsStorage) cacheFeedContent(ctx context.Context, content models.FeedContent) error {
	key := fmt.Sprintf("feed_content:%d", content.Id)
	data, err := json.Marshal(content)
	if err != nil {
		f.logger.Error("Failed to marshal feed content for caching", zap.Error(err))
		return err
	}

	err = f.cache.Set(ctx, key, data, 12*time.Hour).Err() // Cache for 12 hours
	if err != nil {
		f.logger.Error("Failed to set feed content in Redis", zap.Error(err))
		return err
	}

	f.logger.Info("Cached feed content in Redis", zap.String("key", key))
	return nil
}
