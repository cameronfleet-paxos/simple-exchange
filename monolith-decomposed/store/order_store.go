package store

import (
	"fmt"
	"sync"

	"github.com/approved-designs/simple-exchange/monolith-decomposed/model"
)

type OrderStore interface {
	UpsertOrder(order *model.Order)
	GetOrder(id string) (model.Order, error)
	GetAllOrders() []*model.Order
	RemoveOrder(id string) error
}

// InMemoryOrderStore is a mock database that stores orders.
type InMemoryOrderStore struct {
	orders map[string]*model.Order
	mutex  sync.Mutex
}

// UpsertOrder upserts an order to db
func (db *InMemoryOrderStore) UpsertOrder(order *model.Order) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.orders[order.ID] = order
}

// GetOrder retrieves an order from the database by ID.
func (db *InMemoryOrderStore) GetOrder(id string) (*model.Order, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	order, ok := db.orders[id]
	if !ok {
		return nil, fmt.Errorf("order with ID '%s' not found", id)
	}

	return order, nil
}

// GetAllOrders retrieves all orders from the database.
func (db *InMemoryOrderStore) GetAllOrders() []*model.Order {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	orders := make([]*model.Order, 0, len(db.orders))
	for _, order := range db.orders {
		orders = append(orders, order)
	}

	return orders
}

// RemoveOrder removes an order from the database by ID.
func (db *InMemoryOrderStore) RemoveOrder(id string) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	if _, ok := db.orders[id]; !ok {
		return fmt.Errorf("order with ID '%s' not found", id)
	}

	delete(db.orders, id)

	return nil
}
