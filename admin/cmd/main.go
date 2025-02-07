/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-10-09 00:07:19
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-30 01:17:31
 * @FilePath: /sphere_posts/cmd/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
	grpc "github.com/projects/pro-sphere-backend/admin/internal/items/gRPC"
	"github.com/projects/pro-sphere-backend/admin/internal/items/gRPC/handlers"
	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
	"github.com/projects/pro-sphere-backend/admin/internal/items/storage"
	"github.com/projects/pro-sphere-backend/admin/logger"
	"github.com/projects/pro-sphere-backend/admin/pkg/redisconn"
	"go.uber.org/zap"
)

func main() {
	logger.Initialize()
	defer logger.Sync()
	zapLogger := logger.GetLogger()

	cfg, err := config.New()
	// feedParser := gofeed.NewParser()

	if err != nil {
		logger.Logger.Panic("Failed to load configuration", zap.Error(err))
	}
	rConn := redisconn.ConnectDB(&cfg.Redis)

	postgresDB, err := storage.ConnectDB(cfg)
	if err != nil {
		logger.Logger.Panic("Failed to connect to database", zap.Error(err))
	}
	//storages
	feedsStorage := storage.NewFeedsStorage(postgresDB, zapLogger, cfg, rConn)
	feedItemsStorage := storage.NewFeedItemsStorage(postgresDB, zapLogger, cfg)
	feedCategoriesStorage := storage.NewFeedCategoriesStorage(postgresDB, zapLogger, rConn)

	//services
	feedsService := service.NewFeedsService(feedsStorage, zapLogger)
	// feedItemsService := service.NewFeedItemsService(feedsStorage, feedItemsStorage, zapLogger, feedParser)
	feedCategoriesService := service.NewCategoriesService(feedCategoriesStorage, zapLogger)

	//gRPC handlers
	feedshandler := handlers.NewFeedServiceServer(feedsService, zapLogger)
	feedItemsHandler := handlers.NewFeedItemsHandler(feedItemsStorage, zapLogger)
	feedCategoriesHandler := handlers.NewCategoryHandler(feedCategoriesService, zapLogger)

	gRPCServer := grpc.NewSearchgRPCServer(feedCategoriesHandler, feedshandler, feedItemsHandler)

	go func() {
		gRPCServer.Run(cfg.Server.Host, cfg.Server.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	logger.Logger.Info("Server stopped gracefully")
}

// package main

// import (
// 	"context"

// 	"github.com/projects/pro-sphere-backend/admin/internal/items/config"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/models"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/service"
// 	"github.com/projects/pro-sphere-backend/admin/internal/items/storage"
// 	"github.com/projects/pro-sphere-backend/admin/logger"
// 	"github.com/projects/pro-sphere-backend/admin/pkg/redisconn"
// 	"go.uber.org/zap"
// )

// func main() {
// 	logger.Initialize()
// 	defer logger.Sync()
// 	zapLogger := logger.GetLogger()

// 	cfg, err := config.New()
// 	// feedParser := gofeed.NewParser()

// 	if err != nil {
// 		logger.Logger.Panic("Failed to load configuration", zap.Error(err))
// 	}
// 	rConn := redisconn.ConnectDB(&cfg.Redis)

// 	postgresDB, err := storage.ConnectDB(cfg)
// 	if err != nil {
// 		logger.Logger.Panic("Failed to connect to database", zap.Error(err))
// 	}
// 	//storages
// 	feedCategoriesStorage := storage.NewFeedCategoriesStorage(postgresDB, zapLogger, rConn)

// 	//services
// 	feedCategoriesService := service.NewCategoriesService(feedCategoriesStorage, zapLogger)

// 	feedCategoriesService.CreateFeedCategory(context.TODO(), "https://cdn-icons-png.flaticon.com/512/981/981272.png", []models.FeedCategoryTranslation{{
// 		Lang: "uz",
// 		Name: "Siyosat",
// 	},
// 		{
// 			Lang: "ru",
// 			Name: "Полтика",
// 		},
// 		{
// 			Lang: "en",
// 			Name: "Politics",
// 		}})
// }
