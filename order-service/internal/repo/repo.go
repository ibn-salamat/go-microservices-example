package repo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo interface {
	Order() OrderRepo
	User() UserRepo
}

type repo struct {
	db *mongo.Client
}

func New(db *mongo.Client) Repo {
	return repo{
		db: db,
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
