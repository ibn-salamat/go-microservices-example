package service

import (
	"encoding/json"
	"log"
	"orders/internal/repo"

	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	Order() OrderService
	User() UserService
	ListenRMQ() error
}

type service struct {
	r      repo.Repo
	broker *amqp091.Connection
}

func New(r repo.Repo, broker *amqp091.Connection) Service {
	return service{
		r:      r,
		broker: broker,
	}
}

func (s service) Order() OrderService {
	o := os{
		or: s.r.Order(),
		ur: s.r.User(),
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
