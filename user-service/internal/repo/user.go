package repo

import (
	"context"
	"encoding/json"
	"time"
	"users/internal/models"
	"users/internal/transport/http/handler/payload"

	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type UserRepo interface {
	GetByID(id int) (*models.User, error)
	Create(payload.CreateUserPayload) (*models.User, error)
	SendUserIDToBroker(id int) error
}

type ur struct {
	db     *gorm.DB
	broker *amqp091.Connection
}

func newUserRepo(db *gorm.DB, broker *amqp091.Connection) ur {
	return ur{
		db:     db,
		broker: broker,
	}
}

func (r ur) GetByID(id int) (*models.User, error) {
	var user models.User

	result := r.db.Model(&user).Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r ur) Create(payload payload.CreateUserPayload) (*models.User, error) {
	user := models.User{
		Username: payload.Username,
		Email:    payload.Email,
	}

	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r ur) SendUserIDToBroker(id int) error {
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
