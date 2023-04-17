package main

// Order represents a single order.
type Order struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}

// MatchedOrder represents a pair of matching orders.
type MatchedOrder struct {
	BuyOrderID  string  `json:"buy_order_id"`
	SellOrderID string  `json:"sell_order_id"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Timestamp   int64   `json:"timestamp"`
}

// AuditEvent represents a single audit event.
type AuditEvent struct {
	ID         string      `json:"id"`
	Action     string      `json:"action"`
	EntityType string      `json:"entity_type"`
	EntityID   string      `json:"entity_id"`
	Details    interface{} `json:"details"`
	Timestamp  int64       `json:"timestamp"`
}

// TradeCaptureReport represents a single trade capture report.
type TradeCaptureReport struct {
	ID          string  `json:"id"`
	BuyOrderID  string  `json:"buy_order_id"`
	SellOrderID string  `json:"sell_order_id"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Timestamp   int64   `json:"timestamp"`
}

type Venue interface {
	GetMatchingOrders(symbol string, minPrice float64) []Order
}