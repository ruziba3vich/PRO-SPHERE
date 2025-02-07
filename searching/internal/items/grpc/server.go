package grpc

import (
	"log"
	"net"
	"search_service/genproto/searching"
	"search_service/internal/items/grpc/handlers"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type SearchgRPCServer struct {
	Handler *handlers.SearchingHandler
}

func NewSearchgRPCServer(handler *handlers.SearchingHandler) *SearchgRPCServer {
	return &SearchgRPCServer{Handler: handler}
}

func (s *SearchgRPCServer) Run(host, port string) error {
	listener, err := net.Listen("tcp", host+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", port, err)
		return err
	}

	server := grpc.NewServer()

	searching.RegisterSearchingServiceServer(server, s.Handler)

	reflection.Register(server)

	log.Printf("gRPC server is running on port %s", port)

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
		return err
	}

	return nil
}
