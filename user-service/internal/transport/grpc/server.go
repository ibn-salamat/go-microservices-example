package grpc

import (
	"context"
	"log"
	"net"
	"users/config"
	pb "users/internal/proto"
	"users/internal/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UserServiceServer
	service service.Service
}

func Listen(conf config.Config, service service.Service) {
	lis, err := net.Listen("tcp", conf.GrpcPort)
	if err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterUserServiceServer(s, &server{
		service: service,
	})

	log.Printf("gRPC server listening on %s", conf.GrpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to listen gRPC: %v", err)
	}
}

func (s *server) GetUserById(ctx context.Context, req *pb.GetUserByIdRequest) (*pb.GetUserByIdResponse, error) {
	user, err := s.service.User().GetByID(int(req.UserId))
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found by id: %d", req.UserId)
	}

	return &pb.GetUserByIdResponse{
		UserId:   int32(user.ID),
		Email:    user.Email,
		Username: user.Username,
	}, nil
}
