package models

type Order struct {
	OrderId   string `bson:"_id,omitempty" json:"id,omitempty"`
	UserId    int    `bson:"user_id" json:"name"`
	ProductId int    `bson:"product_id" json:"name"`
	Quantity  int    `bson:"quantity" json:"name"`
}
