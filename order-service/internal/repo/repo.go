package repo

import (
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	Order() OrderRepo
	User() UserRepo
	ListenRMQ() error
}

type repo struct {
	db     *mongo.Client
	broker *amqp091.Connection
}

func New(db *mongo.Client, broker *amqp091.Connection) Repo {
	return repo{
		db:     db,
		broker: broker,
	}
}

func (r repo) Order() OrderRepo {
	u := newOrderRepo(r.db)
	return u
}

func (r repo) User() UserRepo {
	u := newUserRepo(r.db)
	return u
}

func (r repo) ListenRMQ() error {
	ch, err := r.broker.Channel()
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

		r.User().Create(userId)
	}

	return nil
}
