/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:36:57
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-29 14:42:01
 * @FilePath: /admin/internal/items/gRPC/grpcserver.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package grpc

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/projects/pro-sphere-backend/admin/genproto/genproto/feeds"
	"github.com/projects/pro-sphere-backend/admin/internal/items/gRPC/handlers"
	"google.golang.org/grpc"

	"google.golang.org/grpc/reflection"
)

type gRPCServer struct {
	feedCatHandler   *handlers.FeedsCategoryHandler
	feedsHandler     *handlers.FeedServiceServer
	feedItemsHandler *handlers.FeedsItemsHandler
}

func NewSearchgRPCServer(feedCatHandler *handlers.FeedsCategoryHandler,
	feedsHandler *handlers.FeedServiceServer,
	feedItemsHandler *handlers.FeedsItemsHandler) *gRPCServer {

	return &gRPCServer{
		feedCatHandler:   feedCatHandler,
		feedsHandler:     feedsHandler,
		feedItemsHandler: feedItemsHandler,
	}
}

func (s *gRPCServer) Run(host, port string) error {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
		return err
	}

	server := grpc.NewServer()

	feeds.RegisterCategoriesServiceServer(server, s.feedCatHandler)
	feeds.RegisterFeedItemsServiceServer(server, s.feedItemsHandler)
	feeds.RegisterFeedsServiceServer(server, s.feedsHandler)

	reflection.Register(server)

	log.Printf("gRPC server is running on port %s", port)

	// Graceful shutdown
	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Shutting down gRPC server gracefully...")
	server.GracefulStop()
	return nil
}
