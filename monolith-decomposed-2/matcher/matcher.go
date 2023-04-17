package matcher

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/approved-designs/simple-exchange/monolith-decomposed/audit"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/order"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/venue"
)

type MatchedOrder struct {
	BuyOrderID  string  `json:"buy_order_id"`
	SellOrderID string  `json:"sell_order_id"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Timestamp   int64   `json:"timestamp"`
}

type tradeCaptureReport struct {
	ID          string  `json:"id"`
	BuyOrderID  string  `json:"buy_order_id"`
	SellOrderID string  `json:"sell_order_id"`
	Quantity    float64 `json:"quantity"`
	Price       float64 `json:"price"`
	Timestamp   int64   `json:"timestamp"`
}

// Matcher is responsible for matching buy and sell orders and creating matched orders.
type Matcher struct {
	orderService order.OrderService
	auditor      audit.Auditor
	matched      []*MatchedOrder
	venues       []venue.Venue
	matchedLock  sync.Mutex
}

// NewMatcher creates a new instance of Matcher.
func NewMatcher(orderService order.OrderService, auditor audit.Auditor, venues []venue.Venue) *Matcher {
	return &Matcher{
		orderService: orderService,
		auditor:      auditor,
		matched:      make([]*MatchedOrder, 0),
		venues:       venues,
		matchedLock:  sync.Mutex{},
	}
}

// MatchOrder matches a new order against existing orders in the database.
func (m *Matcher) MatchOrder(incomingOrder *order.Order) error {
	if err := order.VerifyOrder(incomingOrder); err != nil {
		return err
	}

	// Find all matching orders in all venues
	matchingOrders := make([]*order.Order, 0)
	for _, venue := range m.venues {
		for _, existingOrder := range venue.GetMatchingOrders(incomingOrder.Symbol, incomingOrder.Quantity) {
			matchingOrders = append(matchingOrders, &existingOrder)
		}
	}

	// Sort matching orders by timestamp.
	sortOrdersByTimestamp(matchingOrders)

	// Fill the order with the first matching order.
	for _, matchingOrder := range matchingOrders {
		if matchingOrder.Quantity >= incomingOrder.Quantity {
			matchingOrder.Quantity -= incomingOrder.Quantity
			m.orderService.UpsertOrder(matchingOrder)
			m.orderService.UpsertOrder(incomingOrder)
			matchedOrder := m.createMatchedOrder(incomingOrder, matchingOrder)
			m.createAuditEvent(incomingOrder, matchingOrder)
			m.createTradeCaptureReport(&matchedOrder)
			return nil
		} else {
			incomingOrder.Quantity -= matchingOrder.Quantity
			matchingOrder.Quantity = 0
			m.orderService.UpsertOrder(matchingOrder)
			matchedOrder := m.createMatchedOrder(incomingOrder, matchingOrder)
			m.createAuditEvent(incomingOrder, matchingOrder)
			m.createTradeCaptureReport(&matchedOrder)
		}
	}

	// If the order was not completely filled, add it to the database.
	if incomingOrder.Quantity > 0 {
		m.orderService.UpsertOrder(incomingOrder)
	}
	return nil
}

// createMatchedOrder creates a new matched order and adds it to the list of matched orders.
func (m *Matcher) createMatchedOrder(buyOrder *order.Order, sellOrder *order.Order) MatchedOrder {
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
func (m *Matcher) createAuditEvent(buyOrder *order.Order, sellOrder *order.Order) {
	event := &audit.AuditEvent{
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
	m.auditor.Audit(event)
}

// createTradeCaptureReport creates a new trade captured report and sends it to the post-trade system.
func (m *Matcher) createTradeCaptureReport(match *MatchedOrder) error {
	report := &tradeCaptureReport{
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

// sortOrdersByTimestamp sorts a slice of orders by timestamp in ascending order.
func sortOrdersByTimestamp(orders []*order.Order) {
	orderByTimestamp := func(i, j int) bool {
		return orders[i].Timestamp < orders[j].Timestamp
	}
	sort.SliceStable(orders, orderByTimestamp)
}

// SendTradeCaptureReport sends a trade captured report to the post-trade system.
func sendTradeCaptureReport(report *tradeCaptureReport) error {
	// In a real implementation, this function would send the report to the post-trade system via REST or another protocol.
	// Here, we simply print the report to the console.
	fmt.Println("Sent trade captured report:", report)
	return nil
}
