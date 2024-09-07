package service

import (
	"users/internal/repo"

	"github.com/rabbitmq/amqp091-go"
)

type Service interface {
	User() UserService
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

func (s service) User() UserService {
	u := us{
		r:      s.r.User(),
		broker: s.broker,
	}
	return u
}
