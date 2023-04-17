package model

type Venue interface {
	GetMatchingOrders(symbol string, minPrice float64) []Order
}
