package model

// MatchedOrder represents a pair of matching orders.
type MatchedOrder struct {
	BuyOrderID  string  `json:"buy_order_id"`
	SellOrderID string  `json:"sell_order_id"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Timestamp   int64   `json:"timestamp"`
}
