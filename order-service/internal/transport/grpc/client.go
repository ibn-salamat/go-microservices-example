package grpc

import (
	"log"
	pb "orders/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Connect(url string) pb.UserServiceClient {
	opts := grpc.WithTransportCredentials(insecure.NewCredentials())
	cc, err := grpc.NewClient(url, opts)
	if err != nil {
		log.Fatal(err)
	}

	client := pb.NewUserServiceClient(cc)

	return client
}
