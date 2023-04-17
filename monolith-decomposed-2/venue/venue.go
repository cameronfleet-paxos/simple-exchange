package venue

import "github.com/approved-designs/simple-exchange/monolith-decomposed/order"

type Venue interface {
	GetMatchingOrders(symbol string, minPrice float64) []order.Order
}
