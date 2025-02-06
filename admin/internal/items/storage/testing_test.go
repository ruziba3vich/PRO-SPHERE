package storage

import (
	"context"
	"testing"

	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
	"github.com/projects/pro-sphere-backend/admin/pkg/redisconn"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestCreateFeedCategory(t *testing.T) {
	// Setup PostgreSQL database connection
	dbConfig := &config.Config{
		Database: config.DatabaseConfig{
			User:     "postgres",
			Password: "1702",
			Host:     "localhost",
			Port:     "5432",
			DBName:   "posts_db",
		},
	}
	db, err := ConnectDB(dbConfig)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %s", err)
	}
	defer db.Close()

	// Setup Redis client
	redisConfig := &config.RedisConfig{
		Host: "localhost",
		Port: "6379",
	}
	rdb := redisconn.ConnectDB(redisConfig)

	// Flush Redis to ensure no pre-existing cache data
	err = rdb.FlushAll(context.Background()).Err()
	if err != nil {
		t.Fatalf("Failed to flush Redis cache: %s", err)
	}

	// Set up logger
	logger := zap.NewExample()

	// Instantiate CategoriesStorage
	categoriesStorage := &CategoriesStorage{
		db:     db,
		cache:  rdb,
		logger: logger,
	}

	// Define test category
	category := &models.FeedCategory{
		ID:      1,
		IconURL: "https://example.com/icon.png",
		Translations: []models.FeedCategoryTranslation{
			{Lang: "en", Name: "Test Category"},
		},
	}

	// Test the CreateFeedCategory method
	createdCategory, err := categoriesStorage.CreateFeedCategory(context.Background(), category)
	assert.Nil(t, err)
	assert.NotNil(t, createdCategory)
	assert.Equal(t, category.ID, createdCategory.ID)

	// Verify that the category is added to Redis cache
	cacheKey := "feed_category:1"
	cachedCategory, err := rdb.Get(context.Background(), cacheKey).Result()
	assert.Nil(t, err)
	assert.NotEmpty(t, cachedCategory)
}

func TestDeleteFeedCategory(t *testing.T) {
	// Setup PostgreSQL database connection
	dbConfig := &config.Config{
		Database: config.DatabaseConfig{
			User:     "postgres",
			Password: "1702",
			Host:     "localhost",
			Port:     "5432",
			DBName:   "posts_db",
		},
	}
	db, err := ConnectDB(dbConfig)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %s", err)
	}
	defer db.Close()

	// Setup Redis client
	redisConfig := &config.RedisConfig{
		Host: "localhost",
		Port: "6379",
	}
	rdb := redisconn.ConnectDB(redisConfig)

	// Set up logger
	logger := zap.NewExample()

	// Instantiate CategoriesStorage
	categoriesStorage := &CategoriesStorage{
		db:     db,
		cache:  rdb,
		logger: logger,
	}

	// Ensure there is no pre-existing data
	err = rdb.FlushAll(context.Background()).Err()
	if err != nil {
		t.Fatalf("Failed to flush Redis cache: %s", err)
	}

	// Add a category to delete
	category := &models.FeedCategory{
		ID:      6,
		IconURL: "https://example.com/icon.png",
		Translations: []models.FeedCategoryTranslation{
			{Lang: "en", Name: "Test Category to Delete"},
		},
	}
	_, err = categoriesStorage.CreateFeedCategory(context.Background(), category)
	if err != nil {
		t.Fatalf("Failed to create category: %s", err)
	}

	// Test the DeleteFeedCategory method
	err = categoriesStorage.DeleteFeedCategory(context.Background(), 6)
	assert.Nil(t, err)

	// Verify that the category is deleted from Redis cache
	cacheKey := "feed_category:6"
	_, err = rdb.Get(context.Background(), cacheKey).Result()
	assert.Error(t, err) // Should return an error as the cache should be empty
}

func TestGetAllFeedCategories(t *testing.T) {
	// Setup PostgreSQL database connection
	dbConfig := &config.Config{
		Database: config.DatabaseConfig{
			User:     "postgres",
			Password: "1702",
			Host:     "localhost",
			Port:     "5432",
			DBName:   "posts_db",
		},
	}
	db, err := ConnectDB(dbConfig)
	if err != nil {
		t.Fatalf("Failed to connect to PostgreSQL: %s", err)
	}
	defer db.Close()

	// Setup Redis client
	redisConfig := &config.RedisConfig{
		Host: "localhost",
		Port: "6379",
	}
	rdb := redisconn.ConnectDB(redisConfig)

	// Set up logger
	logger := zap.NewExample()

	// Instantiate CategoriesStorage
	categoriesStorage := &CategoriesStorage{
		db:     db,
		cache:  rdb,
		logger: logger,
	}

	// Ensure no pre-existing data in Redis
	err = rdb.FlushAll(context.Background()).Err()
	if err != nil {
		t.Fatalf("Failed to flush Redis cache: %s", err)
	}

	// Create sample categories in the database
	category1 := &models.FeedCategory{
		ID:      3,
		IconURL: "https://example.com/icon1.png",
		Translations: []models.FeedCategoryTranslation{
			{Lang: "en", Name: "Category 1"},
		},
	}
	_, err = categoriesStorage.CreateFeedCategory(context.Background(), category1)
	if err != nil {
		t.Fatalf("Failed to create category: %s", err)
	}

	category2 := &models.FeedCategory{
		ID:      4,
		IconURL: "https://example.com/icon2.png",
		Translations: []models.FeedCategoryTranslation{
			{Lang: "en", Name: "Category 2"},
		},
	}
	_, err = categoriesStorage.CreateFeedCategory(context.Background(), category2)
	if err != nil {
		t.Fatalf("Failed to create category: %s", err)
	}

	// Test GetAllFeedCategories method
	categories, err := categoriesStorage.GetAllFeedCategories(context.Background(), 1, 10, "en")
	assert.Nil(t, err)
	assert.NotNil(t, categories)
	assert.Len(t, categories, 2)
	assert.Equal(t, "Category 1", categories[0].Translations[0].Name)
}

func TestCacheFeedCategory(t *testing.T) {
	// Setup Redis client
	redisConfig := &config.RedisConfig{
		Host: "localhost",
		Port: "6379",
	}
	rdb := redisconn.ConnectDB(redisConfig)

	// Flush Redis to ensure no pre-existing cache data
	err := rdb.FlushAll(context.Background()).Err()
	if err != nil {
		t.Fatalf("Failed to flush Redis cache: %s", err)
	}

	// Set up logger
	logger := zap.NewExample()

	// Instantiate CategoriesStorage
	categoriesStorage := &CategoriesStorage{
		cache:  rdb,
		logger: logger,
	}

	// Define a category to cache
	category := &models.FeedCategory{
		ID:      6,
		IconURL: "https://example.com/icon.png",
		Translations: []models.FeedCategoryTranslation{
			{Lang: "en", Name: "Test Category"},
		},
	}

	// Test caching the feed category
	err = categoriesStorage.cacheFeedCategory(context.Background(), category)
	assert.Nil(t, err)

	// Verify the cache content
	cacheKey := "feed_category:6"
	cachedCategory, err := rdb.Get(context.Background(), cacheKey).Result()
	assert.Nil(t, err)
	assert.NotEmpty(t, cachedCategory)
}
