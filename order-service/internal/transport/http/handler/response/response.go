package response

type CreateOrderResponse struct {
	OrderId   string `json:"order_id"`
	UserId    int    `json:"user_id"`
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}
