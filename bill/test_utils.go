package bill

import cart "oops-billing-system/cart"

func NewBillWithTwoItems() Bill {
	cartWithTwoItems, _ := cart.NewCartWithTwoItems()
	return Bill{
		cart: cartWithTwoItems,
	}
}
