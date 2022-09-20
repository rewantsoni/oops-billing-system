package cart

import (
	"errors"
	"oops-billing-system/item"
)

type Quantity int

type Cart struct {
	itemWithQuantity map[item.Item]Quantity
}

func (c Cart) GetItemWithQuantity() map[item.Item]Quantity {
	return c.itemWithQuantity
}

func NewCart(items map[item.Item]Quantity) (Cart, error) {
	if err := validateItems(items); err != nil {
		return Cart{}, err
	}

	return Cart{itemWithQuantity: items}, nil
}

func validateItems(items map[item.Item]Quantity) error {
	if len(items) == 0 {
		return errors.New("item cannot be empty")
	}

	for _, quantity := range items {
		if quantity <= 0 {
			return errors.New("quantity cannot be negative or zero")
		}
	}
	return nil
}
