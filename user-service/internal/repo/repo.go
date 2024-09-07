package repo

import (
	"github.com/rabbitmq/amqp091-go"
	"gorm.io/gorm"
)

type Repo interface {
	User() UserRepo
}

type repo struct {
	db     *gorm.DB
	broker *amqp091.Connection
}

func New(db *gorm.DB, broker *amqp091.Connection) Repo {
	return repo{
		db:     db,
		broker: broker,
	}
}

func (r repo) User() UserRepo {
	u := newUserRepo(r.db, r.broker)
	return u
}
