package order

import (
	"fmt"

	"github.com/approved-designs/simple-exchange/monolith-decomposed/model"
)

func VerifyOrder(order *model.Order) error {
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

func verifyKYCReqs(order *model.Order) error {
	return nil
}
