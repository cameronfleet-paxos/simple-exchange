package main

import (
	"fmt"

	"github.com/approved-designs/simple-exchange/monolith-decomposed/audit"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/matcher"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/order"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/venue"
)

func main() {
	auditor := audit.NewAuditor()
	orderService := order.NewOrderService()
	venues := []venue.Venue{}
	matcher := matcher.NewMatcher(*orderService, *auditor, venues)

	order := &order.Order{
		ID:        "ABC",
		Symbol:    "BTC/USD",
		Quantity:  10.0,
		Price:     367_291.0,
		Timestamp: 1681719620,
	}
	matcher.MatchOrder(order)
	matched := matcher.GetMatchedOrders()
	fmt.Println(matched)
}
