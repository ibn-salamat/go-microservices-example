package repo

import (
	"context"
	"orders/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo interface {
	Create(userId int) error
	GetByID(id int) (*models.User, error)
}

type ur struct {
	db *mongo.Client
}

func newUserRepo(db *mongo.Client) ur {
	return ur{
		db: db,
	}
}

func (r ur) Create(userId int) error {
	collection := r.db.Database("orders").Collection("users")

	user := models.User{
		UserId: userId,
	}

	_, err := collection.InsertOne(context.TODO(), user)
	return err
}

func (r ur) GetByID(id int) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := r.db.Database("orders").Collection("users")

	var user models.User
	err := collection.FindOne(ctx, bson.M{"user_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
