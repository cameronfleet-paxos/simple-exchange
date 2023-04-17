package order

import (
	"fmt"
)

func VerifyOrder(order *Order) error {
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

func verifyKYCReqs(order *Order) error {
	return nil
}
