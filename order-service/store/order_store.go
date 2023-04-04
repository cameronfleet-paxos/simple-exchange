package store

import (
	"github.com/approved-designs/simple-exchange/order-service/model"
	"github.com/google/uuid"
)

type OrderStore interface {
	GetAll() []model.Order
	Upsert(order model.Order)
}

type InMemoryOrderStore struct {
	orders []model.Order
}

func NewInMemoryOrderStore() OrderStore {
	orderStore := new(InMemoryOrderStore)
	orderStore.orders = []model.Order{
		{
			Id:       uuid.New(),
			Symbol:   model.Symbol{BidAsset: model.BTC, AskAsset: model.ETH},
			BidPrice: 2,
			AskPrice: 56516,
		},
		{
			Id:       uuid.New(),
			Symbol:   model.Symbol{BidAsset: model.BTC, AskAsset: model.ETH},
			BidPrice: 1,
			AskPrice: 28258,
		},
	}
	return orderStore
}

func (s *InMemoryOrderStore) GetAll() []model.Order {
	return s.orders
}

func (s *InMemoryOrderStore) Upsert(order model.Order) {
	s.orders = append(s.orders, order)
}
