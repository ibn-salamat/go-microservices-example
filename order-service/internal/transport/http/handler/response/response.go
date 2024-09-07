package response

type CreateOrderResponse struct {
	OrderId   string `json:"order_id"`
	UserId    int    `json:"user_id"`
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type GetOrderByIdResponse struct {
	OrderId   string `json:"order_id"`
	UserId    int    `json:"user_id"`
	ProductId int    `json:"product_id"`
	Quantity  int    `json:"quantity"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}
