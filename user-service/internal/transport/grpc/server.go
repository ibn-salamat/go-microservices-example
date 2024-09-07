package grpc

import (
	"log"
	"net"
	"users/config"
	"users/internal/service"

	"google.golang.org/grpc"
)

func Listen(conf config.Config, service service.Service) {
	lis, err := net.Listen("tcp", conf.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}

	s := grpc.NewServer()

	// logs.RegisterLogServiceServer(s, &LogServer{Models: app.Models})

	log.Printf("gRPC server listening on %s", conf.GrpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
}
