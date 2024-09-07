package service

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"users/internal/models"
	"users/internal/repo"
	"users/internal/transport/http/handler/payload"

	"github.com/rabbitmq/amqp091-go"
)

type UserService interface {
	GetByID(id int) (*models.User, error)
	Create(payload.CreateUserPayload) (*models.User, error)
	SendUserIDToBroker(id int) error
}

type us struct {
	r      repo.UserRepo
	broker *amqp091.Connection
}

func (s us) GetByID(id int) (*models.User, error) {
	return s.r.GetByID(id)
}

func (s us) Create(payload payload.CreateUserPayload) (*models.User, error) {
	user, err := s.r.Create(payload)
	if err != nil {
		return nil, err
	}

	err = s.SendUserIDToBroker(user.ID)
	if err != nil {
		log.Printf("Failed to send user_id to broker: %s", err)
	}

	return user, nil
}

func (s us) SendUserIDToBroker(id int) error {
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	m := make(map[string]int)
	m["user_id"] = id
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        []byte(j),
		})
	if err != nil {
		return err
	}

	return nil
}
