package store

import (
	"github.com/approved-designs/simple-exchange/order-service/model"
	"github.com/google/uuid"
)

type OrderStore interface {
	GetAll() []model.Order
	Upsert(order model.Order)
}

type OrderStoreImpl struct {
	orders []model.Order
}

func NewOrderStore() OrderStore {
	orderStore := new(OrderStoreImpl)
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

func (s *OrderStoreImpl) GetAll() []model.Order {
	return s.orders
}

func (s *OrderStoreImpl) Upsert(order model.Order) {
	s.orders = append(s.orders, order)
}
