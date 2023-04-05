package external

import (
	"github.com/approved-designs/simple-exchange/matching-service/httputil"
	"github.com/approved-designs/simple-exchange/matching-service/model"
)

const orderServiceUrl = "http://localhost:8080/"

type OrderService interface {
	GetAll() ([]model.Order, error)
}

type HttpOrderService struct{}

func (h *HttpOrderService) GetAll() ([]model.Order, error) {
	return httputil.GetAll[model.Order](orderServiceUrl, "orders")
}

