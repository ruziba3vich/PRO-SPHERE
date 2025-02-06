package storage

import (
	"context"
	"database/sql"

	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"

	"go.uber.org/zap"
)

type (
	FeedItemsStorage struct {
		db     *sql.DB
		logger *zap.Logger
		cfg    *config.Config
	}
)

func NewFeedItemsStorage(db *sql.DB, logger *zap.Logger, cfg *config.Config) *FeedItemsStorage {
	return &FeedItemsStorage{
		db:     db,
		logger: logger,
		cfg:    cfg,
	}
}

func (f *FeedItemsStorage) CreateFeedItem(ctx context.Context, item *models.FeedItem) (*models.FeedItemResponse, error) {

	query := `INSERT INTO feed_items (feed_id, title, description, image_url, published_at) 
			  VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`

	var createdItem models.FeedItemResponse
	err := f.db.QueryRowContext(ctx, query, item.FeedID, item.Title, item.Description, item.ImageURL, item.PublishedAt).
		Scan(&createdItem.ID, &createdItem.CreatedAt, &createdItem.UpdatedAt)

	if err != nil {
		f.logger.Error("Failed to create feed item", zap.Error(err))
		return nil, err
	}

	createdItem.Title = item.Title
	createdItem.Description = item.Description
	createdItem.ImageURL = item.ImageURL
	createdItem.PublishedAt = item.PublishedAt
	f.logger.Info("Created feed item", zap.Int("id", createdItem.ID))

	return &createdItem, nil
}

func (f *FeedItemsStorage) GetFeedItemByID(ctx context.Context, id int) (*models.FeedItemResponse, error) {
	query := `SELECT id, title, description, image_url, published_at, created_at, updated_at 
			  FROM feed_items WHERE id = $1`

	var item models.FeedItemResponse
	err := f.db.QueryRowContext(ctx, query, id).Scan(&item.ID, &item.Title, &item.Description, &item.ImageURL,
		&item.PublishedAt, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			f.logger.Warn("Feed item not found", zap.Int("id", id))
			return nil, nil
		}
		f.logger.Error("Failed to get feed item by ID", zap.Error(err))
		return nil, err
	}

	return &item, nil
}

func (f *FeedItemsStorage) UpdateFeedItem(ctx context.Context, updatedItem *models.UpdateItem) (*models.FeedItemResponse, error) {
	query := `UPDATE feed_items SET title = $1, description = $2, image_url = $3, updated_at = NOW() 
			  WHERE id = $4 RETURNING id, title, description, image_url, published_at, created_at, updated_at`

	var updated models.FeedItemResponse
	err := f.db.QueryRowContext(ctx, query, updatedItem.Title, updatedItem.Description, updatedItem.ImageURL, updatedItem.Id).
		Scan(&updated.ID, &updated.Title, &updated.Description, &updated.ImageURL, &updated.PublishedAt, &updated.CreatedAt, &updated.UpdatedAt)

	if err != nil {
		f.logger.Error("Failed to update feed item", zap.Error(err))
		return nil, err
	}

	f.logger.Info("Updated feed item", zap.Int("id", updatedItem.Id))
	return &updated, nil
}

func (f *FeedItemsStorage) DeleteFeedItem(ctx context.Context, id int) error {
	query := `DELETE FROM feed_items WHERE id = $1`

	res, err := f.db.ExecContext(ctx, query, id)
	if err != nil {
		f.logger.Error("Failed to delete feed item", zap.Error(err))
		return err
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		f.logger.Warn("No feed item deleted", zap.Int("id", id))
		return sql.ErrNoRows
	}

	f.logger.Info("Deleted feed item", zap.Int("id", id))
	return nil
}

// type GetAllFeedItemsFilter struct {
// 	Limit      int
// 	Page       int
// 	CategoryId int
// }

func (f *FeedItemsStorage) GetAllFeedItems(ctx context.Context, limit, page int) ([]*models.FeedItemResponse, error) {
	offset := (page - 1) * limit
	query := `SELECT id, title, description, image_url, published_at, created_at, updated_at 
			  FROM feed_items LIMIT $1 OFFSET $2`

	rows, err := f.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		f.logger.Error("Failed to get all feed items", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var items []*models.FeedItemResponse
	for rows.Next() {
		var item models.FeedItemResponse
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.ImageURL, &item.PublishedAt, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			f.logger.Error("Failed to scan feed item", zap.Error(err))
			return nil, err
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		f.logger.Error("Error iterating over feed items", zap.Error(err))
		return nil, err
	}

	return items, nil
}

func (f *FeedItemsStorage) GetAllFeedItemsByFeedId(ctx context.Context, feedId, limit, page int) ([]*models.FeedItemResponse, error) {
	offset := (page - 1) * limit
	query := `SELECT id, title, description, image_url, published_at, created_at, updated_at 
			  FROM feed_items WHERE feed_id = $1 LIMIT $2 OFFSET $3`

	rows, err := f.db.QueryContext(ctx, query, feedId, limit, offset)
	if err != nil {
		f.logger.Error("Failed to get feed items by feed ID", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var items []*models.FeedItemResponse
	for rows.Next() {
		var item models.FeedItemResponse
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.ImageURL, &item.PublishedAt, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			f.logger.Error("Failed to scan feed item", zap.Error(err))
			return nil, err
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		f.logger.Error("Error iterating over feed items", zap.Error(err))
		return nil, err
	}

	return items, nil
}

func (f *FeedItemsStorage) GetAllFeedItemsByFeedCategoryId(ctx context.Context, categoryId, limit, page int) ([]*models.FeedItemResponse, error) {
	offset := (page - 1) * limit
	query := `SELECT fi.id, fi.title, fi.description, fi.image_url, fi.published_at, fi.created_at, fi.updated_at
			  FROM feed_items fi
			  JOIN feeds f ON fi.feed_id = f.id
			  WHERE f.category_id = $1 LIMIT $2 OFFSET $3`

	rows, err := f.db.QueryContext(ctx, query, categoryId, limit, offset)
	if err != nil {
		f.logger.Error("Failed to get feed items by feed category ID", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var items []*models.FeedItemResponse
	for rows.Next() {
		var item models.FeedItemResponse
		err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.ImageURL, &item.PublishedAt, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			f.logger.Error("Failed to scan feed item", zap.Error(err))
			return nil, err
		}
		items = append(items, &item)
	}

	if err = rows.Err(); err != nil {
		f.logger.Error("Error iterating over feed items", zap.Error(err))
		return nil, err
	}

	return items, nil
}
