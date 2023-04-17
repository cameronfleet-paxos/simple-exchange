package order

import (
	"fmt"
	"sync"
)

type orderStore interface {
	upsertOrder(order *Order)
	getOrder(id string) (*Order, error)
	getAllOrders() []*Order
	removeOrder(id string) error
}

// inMemoryOrderStore is a mock database that stores orders.
type inMemoryOrderStore struct {
	orders map[string]*Order
	mutex  sync.Mutex
}

func newInMemoryOrderStore() *inMemoryOrderStore {
	return &inMemoryOrderStore{
		orders: make(map[string]*Order),
		mutex:  sync.Mutex{},
	}
}

// UpsertOrder upserts an order to db
func (db *inMemoryOrderStore) upsertOrder(order *Order) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.orders[order.ID] = order
}

// GetOrder retrieves an order from the database by ID.
func (db *inMemoryOrderStore) getOrder(id string) (*Order, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	order, ok := db.orders[id]
	if !ok {
		return nil, fmt.Errorf("order with ID '%s' not found", id)
	}

	return order, nil
}

// GetAllOrders retrieves all orders from the database.
func (db *inMemoryOrderStore) getAllOrders() []*Order {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	orders := make([]*Order, 0, len(db.orders))
	for _, order := range db.orders {
		orders = append(orders, order)
	}

	return orders
}

// RemoveOrder removes an order from the database by ID.
func (db *inMemoryOrderStore) removeOrder(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.orders[id]; !ok {
		return fmt.Errorf("order with ID '%s' not found", id)
	}

	delete(db.orders, id)

	return nil
}
