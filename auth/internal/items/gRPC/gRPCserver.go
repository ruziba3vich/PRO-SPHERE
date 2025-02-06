/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-11-22 23:36:57
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-18 14:35:26
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

	"github.com/projects/pro-sphere-backend/auth/genproto/genproto/auth"
	"github.com/projects/pro-sphere-backend/auth/internal/items/gRPC/handlers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type gRPCServer struct {
	auth *handlers.AuthHandler
	user *handlers.UserHandler
}

func NewAuthgRPCServer(authHandler *handlers.AuthHandler, userHandler *handlers.UserHandler) *gRPCServer {

	return &gRPCServer{
		auth: authHandler,
		user: userHandler,
	}
}

func (s *gRPCServer) Run(host, port string) error {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
		return err
	}

	server := grpc.NewServer()
	auth.RegisterAuthenticationServer(server, s.auth)
	auth.RegisterUserManagementServer(server, s.user)

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
