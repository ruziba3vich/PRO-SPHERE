/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-10-09 00:07:19
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-18 14:47:33
 * @FilePath: /sphere_posts/cmd/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/projects/pro-sphere-backend/auth/internal/items/config"
	grpc "github.com/projects/pro-sphere-backend/auth/internal/items/gRPC"
	"github.com/projects/pro-sphere-backend/auth/internal/items/gRPC/handlers"
	"github.com/projects/pro-sphere-backend/auth/internal/items/service"
	"github.com/projects/pro-sphere-backend/auth/internal/items/storage"
	"github.com/projects/pro-sphere-backend/auth/internal/items/storage/cache"
	"github.com/projects/pro-sphere-backend/auth/pkg/logger"
	"github.com/projects/pro-sphere-backend/auth/pkg/redis"
	"go.uber.org/zap"
)

func main() {
	logger.Initialize()
	defer logger.Sync()
	zapLogger := logger.GetLogger()

	cfg, err := config.New()

	if err != nil {
		logger.Logger.Panic("Failed to load configuration", zap.Error(err))
	}

	rClient := redis.ConnectDB(&cfg.Redis)

	postgresDB, err := storage.ConnectDB(cfg)
	if err != nil {
		logger.Logger.Panic("Failed to connect to database", zap.Error(err))
	}

	userCache := cache.NewUserCache(rClient, zapLogger)
	//storages
	userStorage := storage.NewUserRepository(postgresDB, rClient, zapLogger)

	authService := service.NewAuthService(userStorage, userCache, cfg, zapLogger)
	userService := service.NewUserService(userStorage, zapLogger)

	authHandler := handlers.NewAuthHandler(authService, userStorage, zapLogger, cfg)
	userHandler := handlers.NewUserHandler(userService, zapLogger)

	gRPCServer := grpc.NewAuthgRPCServer(authHandler, userHandler)

	go func() {
		gRPCServer.Run(cfg.Server.Host, cfg.Server.Port)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop

	logger.Logger.Info("Server stopped gracefully")
}
