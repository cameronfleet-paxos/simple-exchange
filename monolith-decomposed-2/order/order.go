package order

type Order struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

type OrderService struct {
	orderStore orderStore
}

func NewOrderService() *OrderService {
	return &OrderService{
		orderStore: newInMemoryOrderStore(),
	}
}

func (o *OrderService) UpsertOrder(order *Order) {
	o.orderStore.upsertOrder(order)
}
