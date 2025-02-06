package app

// import (
// 	"log"

// 	handlers "github.com/projects/pro-sphere-backend/admin/internal/items/http/handler"

// 	"github.com/gin-gonic/gin"
// 	"github.com/mmcdole/gofeed"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/http/middleware"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/storage"

// 	files "github.com/swaggo/files"
// 	ginSwagger "github.com/swaggo/gin-swagger"
// 	"go.uber.org/zap"
// )

// // @title Pro-Sphere Posts API
// // @version 1.0
// // @description This is a posts API for Pro-Sphere
// // @BasePath /
// // @schemes http
// // @in header
// func Run(cfg *config.Config, logger *zap.Logger) error {
// 	router := gin.Default()

// 	postgersDb, err := storage.ConnectDB(cfg)
// 	if err != nil {
// 		logger.Error("Connection error")
// 		return err
// 	}
// 	feedStorage := storage.NewFeedsStorage(postgersDb, logger, cfg)

// 	feedItemsStorage := storage.NewFeedItemsStorage(postgersDb, logger, cfg)

// 	feedParser := gofeed.NewParser()

// 	feedItemsService := service.NewFeedItemsService(feedStorage, feedItemsStorage, logger, feedParser)

// 	feedsFetcherHandler := handlers.NewFeedsFetcherHandler(feedItemsService, logger)

// 	feedHandler := handlers.NewFeedsHandler(
// 		service.NewFeedsService(
// 			feedStorage,
// 			logger,
// 		), logger,
// 	)
// 	feedCategoriesHandler := handlers.NewCategoriesHandler(
// 		service.NewCategoriesService(
// 			storage.NewCategoriesStorage(postgersDb, logger, cfg),
// 			logger,
// 		), logger,
// 	)

// 	router.Use(middleware.CORS())

// 	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))

// 	feeds := router.Group("/feeds")
// 	{
// 		feeds.POST("", feedHandler.CreateFeed)
// 		feeds.GET("/:id", feedHandler.GetFeed)
// 		feeds.GET("", feedHandler.GetAllFeeds)
// 		feeds.PUT("", feedHandler.UpdateFeed)
// 		feeds.DELETE("/:id", feedHandler.DeleteFeed)
// 	}

// 	categories := router.Group("/feedcategories")
// 	{
// 		categories.POST("", feedCategoriesHandler.CreateCategory)
// 		categories.GET("/:id", feedCategoriesHandler.GetCategory)
// 		categories.GET("", feedCategoriesHandler.GetAllCategories)
// 		categories.PUT("", feedCategoriesHandler.UpdateCategory)
// 		categories.DELETE("/:id", feedCategoriesHandler.DeleteCategory)
// 	}

// 	feedFetchTest := router.Group("/feedTest")
// 	{
// 		feedFetchTest.GET("/:id", feedsFetcherHandler.FetchByID)
// 	}
// 	log.Printf("Server started on: %s", cfg.Server.Port)

// 	if err := router.Run(cfg.Server.Host + ":" + cfg.Server.Port); err != nil {
// 		logger.Error("Failed to start server", zap.Error(err))
// 		return err
// 	}
// 	return nil
// }
