package bill

import (
	"errors"
	"fmt"
	"oops-billing-system/cart"
)

type Bill struct {
	cart  cart.Cart
	total float64
}

func NewBill(cart cart.Cart, total float64) (Bill, error) {
	if !validateTotal(total) {
		return Bill{}, errors.New("total cannot be zero or negative")
	}
	return Bill{cart: cart, total: total}, nil
}

func validateTotal(total float64) bool {
	return total > 0
}

func (b Bill) String() string {
	builder := fmt.Sprintf("Item\tPrice\tQuantity\n")
	for item, quantity := range b.cart.GetItemWithQuantity() {
		builder += fmt.Sprintf("%s\t%f\t%d\n", item.GetName(), item.GetPrice(), quantity)
	}
	builder += fmt.Sprintf("\nTotal: %f", b.total)
	return builder
}
