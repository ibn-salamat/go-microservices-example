package service

import (
	"context"
	"encoding/json"
	"log"
	pb "orders/internal/proto"
	"orders/internal/repo"
	"orders/internal/transport/grpc/response"

	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Order() OrderService
	User() UserService
	ListenRMQ() error
	GetByIDFromGRPC(userId int) (*response.GetUserByIdResponse, error)
}

type service struct {
	r          repo.Repo
	broker     *amqp091.Connection
	grpcClient pb.UserServiceClient
}

func New(r repo.Repo, broker *amqp091.Connection, grpcClient pb.UserServiceClient) Service {
	return service{
		r:          r,
		broker:     broker,
		grpcClient: grpcClient,
	}
}

func (s service) Order() OrderService {
	o := os{
		or:      s.r.Order(),
		ur:      s.r.User(),
		service: s,
	}
	return o
}

func (s service) User() UserService {
	u := us{
		ur: s.r.User(),
	}
	return u
}

func (s service) ListenRMQ() error {
	ch, err := s.broker.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"user-action", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	log.Println("Listening for user-action queue")

	for d := range msgs {
		var result map[string]int

		json.Unmarshal(d.Body, &result)
		userId := result["user_id"]

		log.Printf("Received user_id: %d", userId)

		s.r.User().Create(userId)
	}

	return nil
}

func (s service) GetByIDFromGRPC(userId int) (*response.GetUserByIdResponse, error) {
	request := &pb.GetUserByIdRequest{UserId: int32(userId)}
	resp, err := s.grpcClient.GetUserById(context.Background(), request)
	if err != nil {
		return nil, err
	}

	user := response.GetUserByIdResponse{
		Email:    resp.Email,
		ID:       int(resp.UserId),
		Username: resp.Username,
	}

	return &user, nil
}
