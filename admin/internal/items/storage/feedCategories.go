package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"go.uber.org/zap"
)

// CategoriesStorage struct definition for database and cache access
type CategoriesStorage struct {
	db     *sql.DB
	cache  *redis.Client
	logger *zap.Logger
}

func NewFeedCategoriesStorage(db *sql.DB, logger *zap.Logger, cache *redis.Client) *CategoriesStorage {
	return &CategoriesStorage{
		db:     db,
		logger: logger,
		cache:  cache,
	}
}

// cacheFeedCategory caches the feed category in Redis
func (c *CategoriesStorage) cacheFeedCategory(ctx context.Context, category *models.FeedCategory) error {
	key := fmt.Sprintf("feed_category:%d", category.ID)
	data, err := json.Marshal(category)
	if err != nil {
		c.logger.Error("Failed to marshal feed category for caching", zap.Error(err))
		return err
	}

	err = c.cache.Set(ctx, key, data, 24*time.Hour).Err() // Cache for 24 hours
	if err != nil {
		c.logger.Error("Failed to set feed category in Redis", zap.Error(err))
		return err
	}

	c.logger.Info("Cached feed category in Redis", zap.String("key", key))
	return nil
}

// GetFeedCategoryByID retrieves a feed category by its ID and optional language.
func (c *CategoriesStorage) GetFeedCategoryByID(ctx context.Context, id int64, lang string) (*models.FeedCategory, error) {
	// Generate Redis cache key (include lang if specified)
	cacheKey := fmt.Sprintf("feed_category:%d", id)
	if lang != "" {
		cacheKey = fmt.Sprintf("%s:%s", cacheKey, lang)
	}

	// Attempt to retrieve the category from Redis cache
	var category models.FeedCategory
	cachedData, err := c.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(cachedData), &category); err == nil {
			c.logger.Info("Feed category retrieved from cache", zap.Int64("id", id), zap.String("lang", lang))
			return &category, nil
		}
		c.logger.Error("Failed to unmarshal cached feed category", zap.Error(err))
	}

	// Query the database if cache miss occurs
	var query string
	var args []interface{}
	var translations []models.FeedCategoryTranslation

	if lang == "" {
		// Fetch all translations
		query = `
            SELECT fc.id, fc.icon_url, fc.icon_id, fct.lang, fct.name
            FROM feed_categories fc
            LEFT JOIN feed_category_translations fct ON fc.id = fct.feed_category_id
            WHERE fc.id = $1`
		args = append(args, id)
	} else {
		// Fetch a specific translation
		query = `
            SELECT fc.id, fc.icon_url, fc.icon_id, fct.lang, fct.name
            FROM feed_categories fc
            LEFT JOIN feed_category_translations fct ON fc.id = fct.feed_category_id AND fct.lang = $2
            WHERE fc.id = $1`
		args = append(args, id, lang)
	}

	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		c.logger.Error("Database error while retrieving feed category", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	// Parse results
	for rows.Next() {
		var translation models.FeedCategoryTranslation
		if err := rows.Scan(&category.ID, &category.IconURL, &category.IconID, &translation.Lang, &translation.Name); err != nil {
			c.logger.Error("Failed to scan feed category row", zap.Error(err))
			return nil, err
		}
		translations = append(translations, translation)
	}

	if rows.Err() != nil {
		c.logger.Error("Error iterating over feed category rows", zap.Error(err))
		return nil, rows.Err()
	}

	category.Translations = translations

	// Default Name Fallback (Optional)
	if len(category.Translations) == 0 {
		category.Translations = append(category.Translations, models.FeedCategoryTranslation{
			Lang: "en", Name: "Default Category Name", // Add default fallback if no translation found
		})
	}

	// Cache the result for future use
	if err := c.cacheFeedCategory(ctx, &category); err != nil {
		c.logger.Error("Failed to cache feed category", zap.Error(err))
	}

	return &category, nil
}

// CreateFeedCategory creates a new feed category and caches it
func (c *CategoriesStorage) CreateFeedCategory(ctx context.Context, category *models.FeedCategory) (*models.FeedCategory, error) {
	// Begin a transaction
	tx, err := c.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		c.logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, err
	}

	// Rollback transaction on failure
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	// Insert the category into the database
	query := `
        INSERT INTO feed_categories (icon_url, icon_id)
        VALUES ($1, $2) RETURNING id`
	err = tx.QueryRowContext(ctx, query, category.IconURL, category.IconID).Scan(&category.ID)
	if err != nil {
		c.logger.Error("Failed to create feed category", zap.Error(err))
		return nil, err
	}

	// Insert the translations
	for _, translation := range category.Translations {
		_, err := tx.ExecContext(ctx, `
            INSERT INTO feed_category_translations (feed_category_id, lang, name)
            VALUES ($1, $2, $3)`,
			category.ID, translation.Lang, translation.Name)
		if err != nil {
			c.logger.Error("Failed to insert feed category translation", zap.Error(err))
			return nil, err
		}
	}

	// Commit the transaction
	if err := tx.Commit(); err != nil {
		c.logger.Error("Failed to commit transaction", zap.Error(err))
		return nil, err
	}

	// Cache the newly created category
	if err := c.cacheFeedCategory(ctx, category); err != nil {
		c.logger.Error("Failed to cache feed category after creation", zap.Error(err))
	}

	return category, nil
}

// UpdateFeedCategory updates an existing feed category and caches the updated version
func (c *CategoriesStorage) UpdateFeedCategory(ctx context.Context, category *models.FeedCategory) (*models.FeedCategory, error) {
	// Update the category in the database
	_, err := c.db.ExecContext(ctx, `
        UPDATE feed_categories
        SET icon_url = $1, icon_id = $2 
        WHERE id = $3`,
		category.IconURL, category.IconID, category.ID)
	if err != nil {
		c.logger.Error("Failed to update feed category", zap.Error(err))
		return nil, err
	}

	// Update the translations
	for _, translation := range category.Translations {
		_, err := c.db.ExecContext(ctx, `
            INSERT INTO feed_category_translations (feed_category_id, lang, name)
            VALUES ($1, $2, $3)
            ON CONFLICT (feed_category_id, lang)
            DO UPDATE SET name = EXCLUDED.name`,
			category.ID, translation.Lang, translation.Name)
		if err != nil {
			c.logger.Error("Failed to update feed category translation", zap.Error(err))
			return nil, err
		}
	}

	// Cache the updated category
	if err := c.cacheFeedCategory(ctx, category); err != nil {
		c.logger.Error("Failed to cache updated feed category", zap.Error(err))
	}

	return category, nil
}

// DeleteFeedCategory deletes a feed category and clears the cache
func (c *CategoriesStorage) DeleteFeedCategory(ctx context.Context, id int64) error {
	// Delete the feed category from the database
	_, err := c.db.ExecContext(ctx, `
        DELETE FROM feed_categories WHERE id = $1`, id)
	if err != nil {
		c.logger.Error("Failed to delete feed category", zap.Error(err))
		return err
	}

	// Delete the category translations
	_, err = c.db.ExecContext(ctx, `
        DELETE FROM feed_category_translations WHERE feed_category_id = $1`, id)
	if err != nil {
		c.logger.Error("Failed to delete feed category translations", zap.Error(err))
		return err
	}

	// Delete the cache
	cacheKey := fmt.Sprintf("feed_category:%d", id)
	if err := c.cache.Del(ctx, cacheKey).Err(); err != nil {
		c.logger.Error("Failed to delete cache for feed category", zap.Error(err))
	}

	return nil
}

// GetAllFeedCategories retrieves all feed categories with pagination, language filtering, and caching.
func (c *CategoriesStorage) GetAllFeedCategories(ctx context.Context, page, limit int, lang string) ([]*models.FeedCategory, error) {
	// Generate cache key for pagination and language (page, limit, lang)
	cacheKey := fmt.Sprintf("feed_categories:page:%d:limit:%d:lang:%s", page, limit, lang)

	// Try to retrieve the feed categories from Redis cache
	cachedData, err := c.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		var categories []*models.FeedCategory
		if err := json.Unmarshal([]byte(cachedData), &categories); err == nil {
			if len(categories) > 0 {
				c.logger.Info("Feed categories retrieved from cache", zap.String("cacheKey", cacheKey), zap.Int("page", page), zap.Int("limit", limit), zap.String("lang", lang))
				return categories, nil
			}
			c.logger.Warn("Empty categories found in cache", zap.String("cacheKey", cacheKey))
		} else {
			c.logger.Error("Failed to unmarshal cached feed categories", zap.Error(err))
		}
	} else {
		c.logger.Info("No cached feed categories found", zap.String("cacheKey", cacheKey))
	}

	// Query the database to fetch the feed categories for the requested page and language
	offset := (page - 1) * limit
	args := []interface{}{}

	log.Println(limit, offset, lang)

	var query string
	// If a language is provided, filter by that language
	if lang != "" {
		query = `
        SELECT fc.id, fc.icon_url, fc.icon_id, fct.lang, fct.name
        FROM feed_categories fc
        LEFT JOIN feed_category_translations fct ON fc.id = fct.feed_category_id
        WHERE fct.lang = $1
        LIMIT $2 OFFSET $3`
		args = append(args, lang, limit, offset)
		c.logger.Info("Args", zap.Any("With lang", args))
	} else {
		// If no language is provided, fetch all translations
		query = `
        SELECT fc.id, fc.icon_url, fc.icon_id, fct.lang, fct.name
        FROM feed_categories fc
        LEFT JOIN feed_category_translations fct ON fc.id = fct.feed_category_id
        LIMIT $1 OFFSET $2`
		// Only pass limit and offset
		args = append(args, limit, offset)
		c.logger.Info("Args", zap.Any("without lang", args))
	}
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err == sql.ErrNoRows {
		c.logger.Error("No categories found to match the requirements", zap.Error(err))
		return nil, err
	}

	if err != nil {
		c.logger.Error("Database error while retrieving feed categories", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	count := 0
	// Parse the results into a list of FeedCategory objects
	var categories []*models.FeedCategory
	for rows.Next() {
		count++
		var category models.FeedCategory
		var translation models.FeedCategoryTranslation
		if err := rows.Scan(&category.ID, &category.IconURL, &category.IconID, &translation.Lang, &translation.Name); err != nil {
			c.logger.Error("Failed to scan feed category row", zap.Error(err))
			return nil, err
		}
		// Default empty translations if none exist
		if translation.Name == "" {
			translation.Name = "Default Name" // Default if missing
		}
		category.Translations = append(category.Translations, translation)
		categories = append(categories, &category)
	}

	log.Println("count: ", count)
	if err := rows.Err(); err != nil {
		c.logger.Error("Error iterating over feed category rows", zap.Error(err))
		return nil, err
	}
	if len(categories) == 0 {
		c.logger.Info("No feed categories found in database", zap.Int("page", page), zap.Int("limit", limit), zap.String("lang", lang))
	}

	// Cache the fetched categories for future requests
	data, err := json.Marshal(categories)
	if err != nil {
		c.logger.Error("Failed to marshal feed categories for caching", zap.Error(err))
		return nil, err
	}

	// Set the cache with a TTL of 24 hours (can adjust based on your needs)
	if err := c.cache.Set(ctx, cacheKey, data, 24*time.Hour).Err(); err != nil {
		c.logger.Error("Failed to cache feed categories", zap.Error(err))
	}

	c.logger.Info("Cached feed categories in Redis", zap.String("cacheKey", cacheKey))

	return categories, nil
}

func (c *CategoriesStorage) GetFeedCategoryIconInfo(ctx context.Context, id int) (*models.FeedIconInfo, error) {
	cacheKey := fmt.Sprintf("feed_category_icon_info:%d", id)

	// Try to get from Redis cache
	cachedData, err := c.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		c.logger.Info("Cache hit for FeedCategoryIconInfo", zap.Int("id", id))
		var iconInfo models.FeedIconInfo
		if err := json.Unmarshal([]byte(cachedData), &iconInfo); err == nil {
			return &iconInfo, nil
		}
		c.logger.Error("Failed to unmarshal cached FeedCategoryIconInfo", zap.Error(err))
	}

	// Cache miss: Fetch from database
	var iconInfo models.FeedIconInfo
	query := `SELECT icon_url, icon_id FROM feed_categories WHERE id = $1`
	err = c.db.QueryRowContext(ctx, query, id).Scan(&iconInfo.IconUrl, &iconInfo.IconID)
	if err != nil {
		c.logger.Error("Failed to fetch FeedCategoryIconInfo from database", zap.Error(err))
		return nil, err
	}

	// Serialize and store in Redis
	data, err := json.Marshal(iconInfo)
	if err == nil {
		err = c.cache.Set(ctx, cacheKey, data, 10*time.Minute).Err()
		if err != nil {
			c.logger.Error("Failed to cache FeedCategoryIconInfo", zap.Error(err))
		}
	} else {
		c.logger.Error("Failed to serialize FeedCategoryIconInfo for caching", zap.Error(err))
	}

	return &iconInfo, nil
}
