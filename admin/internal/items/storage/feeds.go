/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:02:16
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2025-01-05 19:56:15
 * @FilePath: /admin/internal/items/storage/feeds.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package storage

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"

	"go.uber.org/zap"
)

type FeedsStorage struct {
	db     *sql.DB
	logger *zap.Logger
	cfg    *config.Config
	cache  *redis.Client
}

func NewFeedsStorage(db *sql.DB, logger *zap.Logger, cfg *config.Config, client *redis.Client) *FeedsStorage {
	return &FeedsStorage{
		db:     db,
		logger: logger,
		cfg:    cfg,
		cache:  client,
	}
}

func (f *FeedsStorage) CreateFeed(ctx context.Context, feed *models.Feed) (*models.Feed, error) {
	tx, err := f.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		f.logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()
	// start inserting base fields of feed
	query := `
		INSERT INTO feeds(priority, max_items, base_url, logo_url, logo_url_id) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at, last_refreshed
	`

	row := tx.QueryRowContext(ctx, query, feed.Priority, feed.MaxItems, feed.BaseUrl, feed.LogoUrl, feed.LogoUrlId)
	if err := row.Err(); err != nil {
		f.logger.Error("Failed to create feed", zap.Error(err))
		return nil, err
	}

	if err := row.Scan(&feed.Id, &feed.CreatedAt, &feed.UpdatedAt, &feed.LastRefreshed); err != nil {
		f.logger.Error("Failed to scan feed creation fields", zap.Error(err))
		return nil, err
	}

	// start inserting translations for a created feed
	for i, translation := range feed.Translations {
		query = `
			INSERT INTO feed_translations (feed_id, lang, title, description)
    			VALUES ($1, $2, $3, $4)	RETURNING id
		`

		row := tx.QueryRowContext(ctx, query, feed.Id, translation.Lang, translation.Title, translation.Description)
		if err := row.Err(); err != nil {
			f.logger.Error("Inserting tranlation failed at index",
				zap.Int("index", i),
				zap.String("lang", translation.Lang),
				zap.String("title", translation.Title),
				zap.Error(err))
			return nil, err
		}

		err = row.Scan(&translation.Id)
		if err != nil {
			f.logger.Error("Failed to scan transalation id", zap.Error(err))
		}
		feed.Translations[i].Id = translation.Id
	}

	if err := tx.Commit(); err != nil {
		f.logger.Error("Failed to commit feed creation transaction", zap.Error(err))
		return nil, err
	}

	if err := f.cacheFeed(ctx, *feed); err != nil {
		f.logger.Error("Failed to cache feed after creation", zap.Error(err))
	}

	return feed, nil
}

// GetFeed retrieves a feed from the cache or database.
func (f *FeedsStorage) GetFeed(ctx context.Context, feedID int) (*models.Feed, error) {
	key := fmt.Sprintf("feeds:%d", feedID)
	cachedFeed, err := f.cache.Get(ctx, key).Result()
	if err == nil {
		var feed models.Feed
		if err := json.Unmarshal([]byte(cachedFeed), &feed); err == nil {
			return &feed, nil
		}
		f.logger.Error("Failed to unmarshal cached feed", zap.Error(err))
	}

	query := `
		SELECT id, priority, max_items, base_url, logo_url, logo_url_id, last_refreshed, created_at, updated_at
		FROM feeds WHERE id = $1
	`
	feed := &models.Feed{}
	row := f.db.QueryRowContext(ctx, query, feedID)
	if err := row.Scan(&feed.Id, &feed.Priority, &feed.MaxItems, &feed.BaseUrl, &feed.LogoUrl, &feed.LogoUrlId, &feed.LastRefreshed, &feed.CreatedAt, &feed.UpdatedAt); err != nil {
		f.logger.Error("Failed to fetch feed from database", zap.Error(err))
		return nil, err
	}

	feed.Translations, err = f.getFeedTranslations(ctx, feedID)
	if err != nil {
		f.logger.Error("Failed to fetch feed translations", zap.Error(err))
		return nil, err
	}

	if err := f.cacheFeed(ctx, *feed); err != nil {
		f.logger.Error("Failed to cache feed", zap.Error(err))
	}

	return feed, nil
}

func (f *FeedsStorage) UpdateFeed(ctx context.Context, feed *models.Feed) (*models.Feed, error) {
	f.logger.Info("Starting UpdateFeed operation", zap.Int("FeedID", feed.Id))

	tx, err := f.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		f.logger.Error("Failed to begin transaction", zap.Error(err))
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			f.logger.Error("Panic recovered during transaction", zap.Any("panic", p))
			tx.Rollback()
			panic(p)
		} else if err != nil {
			f.logger.Error("Transaction rollback due to error", zap.Error(err))
			tx.Rollback()
		} else {
			f.logger.Info("Transaction committed successfully")
			tx.Commit()
		}
	}()

	feed.UpdatedAt = time.Now().GoString()
	f.logger.Info("Setting UpdatedAt for feed", zap.String("UpdatedAt", feed.UpdatedAt))

	query := `
        UPDATE feeds
        SET priority = $1, max_items = $2, base_url = $3, logo_url = $4, 
            logo_url_id = $5, updated_at = NOW()
        WHERE id = $6
    `
	f.logger.Info("Executing main feed update query", zap.String("Query", query))
	_, err = tx.ExecContext(ctx, query, feed.Priority, feed.MaxItems, feed.BaseUrl, feed.LogoUrl, feed.LogoUrlId, feed.Id)
	if err != nil {
		f.logger.Error("Failed to update feed", zap.Error(err))
		return nil, err
	}

	f.logger.Info("Successfully executed main feed update query", zap.Int("FeedID", feed.Id))

	for i, translation := range feed.Translations {
		f.logger.Info("Processing feed translation", zap.Int("Index", i), zap.String("Lang", translation.Lang))

		query = `
            INSERT INTO feed_translations (feed_id, lang, title, description)
            VALUES ($1, $2, $3, $4)
            ON CONFLICT (feed_id, lang)
            DO UPDATE SET title = EXCLUDED.title, description = EXCLUDED.description
        `
		f.logger.Info("Executing translation upsert query", zap.String("Query", query), zap.String("Lang", translation.Lang))
		_, err = tx.ExecContext(ctx, query, feed.Id, translation.Lang, translation.Title, translation.Description)
		if err != nil {
			f.logger.Error("Failed to upsert feed translation", zap.Int("FeedID", feed.Id), zap.String("Lang", translation.Lang), zap.Error(err))
			return nil, err
		}
		f.logger.Info("Successfully upserted feed translation", zap.String("Lang", translation.Lang))
	}

	f.logger.Info("Caching updated feed", zap.Int("FeedID", feed.Id))
	if err := f.cacheFeed(ctx, *feed); err != nil {
		f.logger.Error("Failed to cache updated feed", zap.Error(err))
	} else {
		f.logger.Info("Successfully cached updated feed", zap.Int("FeedID", feed.Id))
	}

	f.logger.Info("UpdateFeed operation completed successfully", zap.Int("FeedID", feed.Id))
	return feed, nil
}

// DeleteFeed removes a feed and all related translations and contents.
func (f *FeedsStorage) DeleteFeed(ctx context.Context, feedID int) error {
	tx, err := f.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		f.logger.Error("Failed to begin transaction", zap.Error(err))
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		}
	}()

	query := `DELETE FROM feeds WHERE id = $1`
	_, err = tx.ExecContext(ctx, query, feedID)
	if err != nil {
		f.logger.Error("Failed to delete feed", zap.Error(err))
		return err
	}

	if err := tx.Commit(); err != nil {
		f.logger.Error("Failed to commit transaction", zap.Error(err))
		return err
	}

	key := fmt.Sprintf("feeds:%d", feedID)
	if err := f.cache.Del(ctx, key).Err(); err != nil {
		f.logger.Error("Failed to delete feed from cache", zap.Error(err))
	}

	return nil
}

// GetAllFeeds retrieves all feeds with translations, optionally filtering by language and paginated.
func (f *FeedsStorage) GetAllFeeds(ctx context.Context, lang *string, limit, page int) ([]models.Feed, error) {
	var query string
	var rows *sql.Rows
	var err error

	offset := (page - 1) * limit

	if lang != nil {
		query = `
				SELECT f.id, f.priority, f.max_items, f.base_url, f.logo_url, f.logo_url_id, f.last_refreshed, f.created_at, f.updated_at
				FROM feeds f
				JOIN feed_translations ft ON f.id = ft.feed_id
				WHERE ft.lang = $1
				ORDER BY f.priority
				LIMIT $2 OFFSET $3
			`
		rows, err = f.db.QueryContext(ctx, query, *lang, limit, offset)
	} else {
		query = `
				SELECT id, priority, max_items, base_url, logo_url, logo_url_id, last_refreshed, created_at, updated_at
				FROM feeds
				ORDER BY priority
				LIMIT $1 OFFSET $2
			`
		rows, err = f.db.QueryContext(ctx, query, limit, offset)
	}

	if err != nil {
		f.logger.Error("Failed to fetch feeds", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var feeds []models.Feed
	for rows.Next() {
		feed := models.Feed{}
		if err := rows.Scan(&feed.Id, &feed.Priority, &feed.MaxItems, &feed.BaseUrl, &feed.LogoUrl, &feed.LogoUrlId, &feed.LastRefreshed, &feed.CreatedAt, &feed.UpdatedAt); err != nil {
			f.logger.Error("Failed to scan feed", zap.Error(err))
			return nil, err
		}
		feed.Translations, err = f.getFeedTranslations(ctx, feed.Id)
		if err != nil {
			f.logger.Error("Failed to fetch feed translations", zap.Error(err))
			return nil, err
		}
		feeds = append(feeds, feed)
	}

	return feeds, nil
}

// getFeedTranslations retrieves translations for a given feed.
func (f *FeedsStorage) getFeedTranslations(ctx context.Context, feedID int) ([]models.Translation, error) {
	query := `
		SELECT id, lang, title, description
		FROM feed_translations WHERE feed_id = $1
	`
	rows, err := f.db.QueryContext(ctx, query, feedID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var translations []models.Translation
	for rows.Next() {
		translation := models.Translation{}
		if err := rows.Scan(&translation.Id, &translation.Lang, &translation.Title, &translation.Description); err != nil {
			return nil, err
		}
		translations = append(translations, translation)
	}

	return translations, nil
}

// cacheFeed caches the feed in Redis.
func (f *FeedsStorage) cacheFeed(ctx context.Context, feed models.Feed) error {
	key := fmt.Sprintf("feeds:%d", feed.Id)
	data, err := json.Marshal(feed)
	if err != nil {
		f.logger.Error("Failed to marshal feed for caching", zap.Error(err))
		return err
	}

	if err := f.cache.Set(ctx, key, data, 12*time.Hour).Err(); err != nil {
		f.logger.Error("Failed to set feed in cache", zap.Error(err))
		return err
	}

	return nil
}

func (c *FeedsStorage) GetFeedLogoInfo(ctx context.Context, id int) (*models.FeedLogoInfo, error) {
	cacheKey := fmt.Sprintf("feed_logo_info:%d", id)

	// Try to get from Redis cache
	cachedData, err := c.cache.Get(ctx, cacheKey).Result()
	if err == nil {
		c.logger.Info("Cache hit for FeedLogoInfo", zap.Int("id", id))
		var logoInfo models.FeedLogoInfo
		if err := json.Unmarshal([]byte(cachedData), &logoInfo); err == nil {
			return &logoInfo, nil
		}
		c.logger.Error("Failed to unmarshal cached FeedLogoInfo", zap.Error(err))
	}

	// Cache miss: Fetch from database
	var logoInfo models.FeedLogoInfo
	query := `SELECT logo_url, logo_url_id FROM feeds WHERE id = $1`
	err = c.db.QueryRowContext(ctx, query, id).Scan(&logoInfo.LogoUrl, &logoInfo.LogoUrlId)
	if err != nil {
		c.logger.Error("Failed to fetch FeedCategoryIconInfo from database", zap.Error(err))
		return nil, err
	}

	// Serialize and store in Redis
	data, err := json.Marshal(logoInfo)
	if err == nil {
		err = c.cache.Set(ctx, cacheKey, data, 10*time.Minute).Err()
		if err != nil {
			c.logger.Error("Failed to cache FeedLogoInfo", zap.Error(err))
		}
	} else {
		c.logger.Error("Failed to serialize FeedLogoInfo for caching", zap.Error(err))
	}

	return &logoInfo, nil
}
