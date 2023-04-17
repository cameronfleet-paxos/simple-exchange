package model

type Order struct {
	ID        string  `json:"id"`
	Symbol    string  `json:"symbol"`
	Quantity  float64 `json:"quantity"`
	Price     float64 `json:"price"`
	Timestamp int64   `json:"timestamp"`
}
