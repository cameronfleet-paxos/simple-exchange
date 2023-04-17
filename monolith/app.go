package main

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	db := NewInMemoryDatabase()
	venues := []Venue{}
	matcher := NewMatcher(db, venues)

	order := &Order{
		ID:        "ABC",
		Symbol:    "BTC/USD",
		Quantity:  10.0,
		Price:     367_291.0,
		Timestamp: 1681719620,
	}
	matcher.MatchOrder(order)
}

// Matcher is responsible for matching buy and sell orders and creating matched orders.
type Matcher struct {
	db          *InMemoryDatabase
	matched     []*MatchedOrder
	venues      []Venue
	matchedLock sync.Mutex
}

// NewMatcher creates a new instance of Matcher.
func NewMatcher(db *InMemoryDatabase, venues []Venue) *Matcher {
	return &Matcher{
		db:          db,
		matched:     make([]*MatchedOrder, 0),
		venues:      venues,
		matchedLock: sync.Mutex{},
	}
}

// MatchOrder matches a new order against existing orders in the database.
func (m *Matcher) MatchOrder(order *Order) error {
	m.db.mutex.Lock()
	defer m.db.mutex.Unlock()

	if err := verifyOrder(order); err != nil {
		return err
	}

	// Find all matching orders in all venues
	matchingOrders := make([]*Order, 0)
	for _, venue := range m.venues {
		for _, existingOrder := range venue.GetMatchingOrders(order.Symbol, order.Quantity) {
			matchingOrders = append(matchingOrders, &existingOrder)
		}
	}

	// Sort matching orders by timestamp.
	SortOrdersByTimestamp(matchingOrders)

	// Fill the order with the first matching order.
	for _, matchingOrder := range matchingOrders {
		if matchingOrder.Quantity >= order.Quantity {
			matchingOrder.Quantity -= order.Quantity
			m.db.orders[matchingOrder.ID] = matchingOrder
			m.db.orders[order.ID] = order
			matchedOrder := m.createMatchedOrder(order, matchingOrder)
			m.createAuditEvent(order, matchingOrder)
			m.createTradeCaptureReport(&matchedOrder)
			return nil
		} else {
			order.Quantity -= matchingOrder.Quantity
			m.db.orders[matchingOrder.ID].Quantity = 0
			matchedOrder := m.createMatchedOrder(order, matchingOrder)
			m.createAuditEvent(order, matchingOrder)
			m.createTradeCaptureReport(&matchedOrder)
		}
	}

	// If the order was not completely filled, add it to the database.
	if order.Quantity > 0 {
		err := m.db.AddOrder(order)
		if err != nil {
			return err
		}
	}

	return nil
}

// SortOrdersByTimestamp sorts a slice of orders by timestamp in ascending order.
func SortOrdersByTimestamp(orders []*Order) {
	orderByTimestamp := func(i, j int) bool {
		return orders[i].Timestamp < orders[j].Timestamp
	}
	sort.SliceStable(orders, orderByTimestamp)
}

// createMatchedOrder creates a new matched order and adds it to the list of matched orders.
func (m *Matcher) createMatchedOrder(buyOrder *Order, sellOrder *Order) MatchedOrder {
	match := &MatchedOrder{
		BuyOrderID:  buyOrder.ID,
		SellOrderID: sellOrder.ID,
		Quantity:    buyOrder.Quantity,
		Price:       sellOrder.Price,
		Timestamp:   time.Now().Unix(),
	}
	m.matchedLock.Lock()
	defer m.matchedLock.Unlock()

	m.matched = append(m.matched, match)
	return *match
}

// GetMatchedOrders retrieves all matched orders.
func (m *Matcher) GetMatchedOrders() []*MatchedOrder {
	m.matchedLock.Lock()
	defer m.matchedLock.Unlock()
	matched := make([]*MatchedOrder, len(m.matched))
	copy(matched, m.matched)

	return matched
}

// createAuditEvent creates a new audit event for a matched order.
func (m *Matcher) createAuditEvent(buyOrder *Order, sellOrder *Order) {
	event := &AuditEvent{
		ID:         fmt.Sprintf("%d", time.Now().UnixNano()),
		Action:     "Matched",
		EntityType: "Order",
		EntityID:   "",
		Details: map[string]interface{}{
			"buy_order_id":  buyOrder.ID,
			"sell_order_id": sellOrder.ID,
			"quantity":      buyOrder.Quantity,
			"price":         sellOrder.Price,
			"timestamp":     time.Now().Unix(),
		},
	}
	m.db.AddAuditEvent(event)
}

// createTradeCaptureReport creates a new trade captured report and sends it to the post-trade system.
func (m *Matcher) createTradeCaptureReport(match *MatchedOrder) error {
	report := &TradeCaptureReport{
		ID:          fmt.Sprintf("%d", time.Now().UnixNano()),
		BuyOrderID:  match.BuyOrderID,
		SellOrderID: match.SellOrderID,
		Quantity:    match.Quantity,
		Price:       match.Price,
		Timestamp:   match.Timestamp,
	}
	err := sendTradeCaptureReport(report)
	if err != nil {
		return err
	}

	return nil
}

// SendTradeCaptureReport sends a trade captured report to the post-trade system.
func sendTradeCaptureReport(report *TradeCaptureReport) error {
	// In a real implementation, this function would send the report to the post-trade system via REST or another protocol.
	// Here, we simply print the report to the console.
	fmt.Println("Sent trade captured report:", report)
	return nil
}

func verifyOrder(order *Order) error {
	if order.Quantity == 0 {
		return fmt.Errorf("order quantity must be greater than zero")
	}

	if order.Price == 0 {
		return fmt.Errorf("order price must be greater than zero")
	}

	if err := verifyKYCReqs(order); err != nil {
		return err
	}

	return nil
}

func verifyKYCReqs(order *Order) error {
	return nil
}
