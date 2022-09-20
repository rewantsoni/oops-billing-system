package cart

import (
	"oops-billing-system/item"
)

func NewCartWithTwoItems() (Cart, error) {
	apple, _ := item.NewItemApple()
	orange, _ := item.NewItemOrange()
	return NewCart(map[item.Item]Quantity{
		apple: 2, orange: 4,
	})
}

func NewCartWithNoItems() (Cart, error) {
	return NewCart(map[item.Item]Quantity{})
}

func NewCartWithZeroOrNegativeQuantity() (Cart, error) {
	apple, _ := item.NewItemApple()
	orange, _ := item.NewItemOrange()
	return NewCart(map[item.Item]Quantity{
		apple: 0, orange: 0,
	})
}
