package main

import (
	"fmt"
	"sync"
)

// InMemoryDatabase is a mock database that stores orders.
type InMemoryDatabase struct {
	orders map[string]*Order
	audits []*AuditEvent
	mutex  sync.Mutex
}

// NewInMemoryDatabase creates a new instance of InMemoryDatabase.
func NewInMemoryDatabase() *InMemoryDatabase {
	return &InMemoryDatabase{
		orders: make(map[string]*Order),
	}
}

// AddOrder adds an order to the database.
func (db *InMemoryDatabase) AddOrder(order *Order) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.orders[order.ID]; ok {
		return fmt.Errorf("order with ID '%s' already exists", order.ID)
	}

	db.orders[order.ID] = order

	return nil
}

// GetOrder retrieves an order from the database by ID.
func (db *InMemoryDatabase) GetOrder(id string) (*Order, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	order, ok := db.orders[id]
	if !ok {
		return nil, fmt.Errorf("order with ID '%s' not found", id)
	}

	return order, nil
}

// GetAllOrders retrieves all orders from the database.
func (db *InMemoryDatabase) GetAllOrders() []*Order {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	orders := make([]*Order, 0, len(db.orders))
	for _, order := range db.orders {
		orders = append(orders, order)
	}

	return orders
}

// RemoveOrder removes an order from the database by ID.
func (db *InMemoryDatabase) RemoveOrder(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.orders[id]; !ok {
		return fmt.Errorf("order with ID '%s' not found", id)
	}

	delete(db.orders, id)

	return nil
}

// AddAuditEvent adds a new audit event to the in-memory database.
func (db *InMemoryDatabase) AddAuditEvent(event *AuditEvent) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.audits = append(db.audits, event)
}