package main

import (
	"github.com/approved-designs/simple-exchange/order-service/handler"
	"github.com/approved-designs/simple-exchange/order-service/service"
	"github.com/approved-designs/simple-exchange/order-service/store"
)

func main() {
	orderStore := store.NewInMemoryOrderStore()
	orderHandler := &handler.OrderHandler{OrderStore: orderStore}
	orderServer := service.OrderService{OrderHandler: orderHandler}
	orderServer.Start()
}
