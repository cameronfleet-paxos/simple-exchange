package matcher

import (
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/approved-designs/simple-exchange/monolith-decomposed/audit"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/model"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/order"
	"github.com/approved-designs/simple-exchange/monolith-decomposed/store"
)

// Matcher is responsible for matching buy and sell orders and creating matched orders.
type Matcher struct {
	orderStore  store.OrderStore
	auditor     audit.Auditor
	matched     []*model.MatchedOrder
	venues      []model.Venue
	matchedLock sync.Mutex
}

// NewMatcher creates a new instance of Matcher.
func NewMatcher(orderStore store.OrderStore, auditor audit.Auditor, venues []model.Venue) *Matcher {
	return &Matcher{
		orderStore:  orderStore,
		auditor:     auditor,
		matched:     make([]*model.MatchedOrder, 0),
		venues:      venues,
		matchedLock: sync.Mutex{},
	}
}

// MatchOrder matches a new order against existing orders in the database.
func (m *Matcher) MatchOrder(incomingOrder *model.Order) error {
	if err := order.VerifyOrder(incomingOrder); err != nil {
		return err
	}

	// Find all matching orders in all venues
	matchingOrders := make([]*model.Order, 0)
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
			m.orderStore.UpsertOrder(matchingOrder)
			m.orderStore.UpsertOrder(incomingOrder)
			matchedOrder := m.createMatchedOrder(incomingOrder, matchingOrder)
			m.createAuditEvent(incomingOrder, matchingOrder)
			m.createTradeCaptureReport(&matchedOrder)
			return nil
		} else {
			incomingOrder.Quantity -= matchingOrder.Quantity
			matchingOrder.Quantity = 0
			m.orderStore.UpsertOrder(matchingOrder)
			matchedOrder := m.createMatchedOrder(incomingOrder, matchingOrder)
			m.createAuditEvent(incomingOrder, matchingOrder)
			m.createTradeCaptureReport(&matchedOrder)
		}
	}

	// If the order was not completely filled, add it to the database.
	if incomingOrder.Quantity > 0 {
		m.orderStore.UpsertOrder(incomingOrder)
	}
	return nil
}

// sortOrdersByTimestamp sorts a slice of orders by timestamp in ascending order.
func sortOrdersByTimestamp(orders []*model.Order) {
	orderByTimestamp := func(i, j int) bool {
		return orders[i].Timestamp < orders[j].Timestamp
	}
	sort.SliceStable(orders, orderByTimestamp)
}

// createMatchedOrder creates a new matched order and adds it to the list of matched orders.
func (m *Matcher) createMatchedOrder(buyOrder *model.Order, sellOrder *model.Order) model.MatchedOrder {
	match := &model.MatchedOrder{
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
func (m *Matcher) GetMatchedOrders() []*model.MatchedOrder {
	m.matchedLock.Lock()
	defer m.matchedLock.Unlock()
	matched := make([]*model.MatchedOrder, len(m.matched))
	copy(matched, m.matched)

	return matched
}

// createAuditEvent creates a new audit event for a matched order.
func (m *Matcher) createAuditEvent(buyOrder *model.Order, sellOrder *model.Order) {
	event := &model.AuditEvent{
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
func (m *Matcher) createTradeCaptureReport(match *model.MatchedOrder) error {
	report := &model.TradeCaptureReport{
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
func sendTradeCaptureReport(report *model.TradeCaptureReport) error {
	// In a real implementation, this function would send the report to the post-trade system via REST or another protocol.
	// Here, we simply print the report to the console.
	fmt.Println("Sent trade captured report:", report)
	return nil
}
