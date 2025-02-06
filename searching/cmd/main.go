/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-20 22:02:17
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-11-20 23:43:18
 * @FilePath: /searching/cmd/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"os"
	"os/signal"
	"search_service/internal/items/config"
	"search_service/internal/items/grpc"
	"search_service/internal/items/grpc/handlers"
	"search_service/internal/items/service"
	"search_service/internal/items/storage/postgres/integrations"
	"search_service/pkg/logger"
	"syscall"

	"go.uber.org/zap"
)

func main() {
	logger.Initialize()
	defer logger.Sync()
	log := logger.GetLogger()
	cfg, err := config.New()
	if err != nil {
		logger.Logger.Panic("Failed to load configuration", zap.Error(err))
	}

	intgSearch := integrations.NewSearchInegreation(cfg, log)

	searchService := service.NewSearchService(intgSearch, log)
	searchHandler := handlers.NewSerachHandler(searchService, log)
	grpcService := grpc.NewSearchgRPCServer(searchHandler)

	go func() {
		if err := grpcService.Run(cfg.Server.Host, cfg.Server.Port); err != nil {
			logger.Logger.Fatal("Application exited with error", zap.Error(err))
		}
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	logger.Logger.Info("Server stopped gracefully")
}
