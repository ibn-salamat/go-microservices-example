package repo

import (
	"orders/internal/models"
	"orders/internal/transport/http/handler/payload"

	"go.mongodb.org/mongo-driver/mongo"
)

type OrderRepo interface {
	GetByID(id int) (*models.Order, error)
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

func (r or) GetByID(id int) (*models.Order, error) {
	var order models.Order

	return &order, nil
}

func (r or) Create(payload payload.CreateOrderPayload) (*models.Order, error) {
	order := models.Order{
		UserId:    payload.UserId,
		Quantity:  payload.Quantity,
		ProductId: payload.ProductId,
	}

	return &order, nil
}
