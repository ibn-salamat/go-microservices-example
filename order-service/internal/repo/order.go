package repo

import (
	"context"
	"orders/internal/models"
	"orders/internal/transport/http/handler/payload"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo interface {
	GetByID(id string) (*models.Order, error)
	Create(payload.CreateOrderPayload) (*models.Order, error)
}

type or struct {
	db *mongo.Client
}

func newOrderRepo(db *mongo.Client) or {
	return or{
		db: db,
	}
}

func (r or) Create(payload payload.CreateOrderPayload) (*models.Order, error) {
	collection := r.db.Database("orders").Collection("orders")

	order := models.Order{
		UserId:    payload.UserId,
		Quantity:  payload.Quantity,
		ProductId: payload.ProductId,
	}

	res, err := collection.InsertOne(context.TODO(), order)
	if err != nil {
		return nil, err
	}

	order.OrderId = res.InsertedID.(primitive.ObjectID).Hex()

	return &order, nil
}

func (r or) GetByID(id string) (*models.Order, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	collection := r.db.Database("orders").Collection("orders")
	orderId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var order models.Order
	err = collection.FindOne(ctx, bson.M{"_id": orderId}).Decode(&order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
