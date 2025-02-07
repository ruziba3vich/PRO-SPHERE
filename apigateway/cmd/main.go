/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-21 00:48:05
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-09 00:19:43
 * @FilePath: /apigateway/cmd/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/projects/pro-sphere-backend/apigateway/genproto/auth"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/feeds"
	"github.com/projects/pro-sphere-backend/apigateway/genproto/searching"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/config"
	"github.com/projects/pro-sphere-backend/apigateway/internal/items/http/app"
	"github.com/projects/pro-sphere-backend/apigateway/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	logger.Initialize()
	log := logger.GetLogger()
	cfg := config.Config{}
	if err := cfg.Load(); err != nil {
		logger.Logger.Panic("Failed to load configuration", zap.Error(err))
	}

	searchConn, err := grpc.NewClient(cfg.SearchServices.Host+":"+cfg.SearchServices.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Fatal("Failed to create search connection", zap.Error(err))
	}

	adminConn, err := grpc.NewClient(cfg.AdminServices.Host+":"+cfg.AdminServices.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Fatal("Failed to create admin connection", zap.Error(err))
	}

	authConn, err := grpc.NewClient(cfg.AuthService.Host+":"+cfg.AuthService.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.Logger.Fatal("Failed to create auth connection", zap.Error(err))
	}

	userClient := auth.NewUserManagementClient(authConn)
	authClient := auth.NewAuthenticationClient(authConn)
	searchingClient := searching.NewSearchingServiceClient(searchConn)
	feedCatClient := feeds.NewCategoriesServiceClient(adminConn)
	feedsClient := feeds.NewFeedsServiceClient(adminConn)
	feedItemsClient := feeds.NewFeedItemsServiceClient(adminConn)

	httpServer := app.NewHttpService(searchingClient, feedCatClient, feedsClient, feedItemsClient, authClient, userClient, log, &cfg)

	go func() {
		if err := httpServer.Run(); err != nil {
			log.Fatal("Application exited with error", zap.Error(err))

		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	logger.Logger.Info("Server stopped gracefully")
}
