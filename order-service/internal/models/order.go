package models

type Order struct {
	OrderId   string `bson:"_id,omitempty"`
	UserId    int    `bson:"user_id"`
	ProductId int    `bson:"product_id" `
	Quantity  int    `bson:"quantity"`
}
