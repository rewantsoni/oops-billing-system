package bill

import (
	"oops-billing-system/cart"
)

type BillGenerator interface {
	Generate(cart.Cart, float64) (Bill, error)
}

type SimpleBillGenerator struct {
}

func NewBillGenerator() BillGenerator {
	return SimpleBillGenerator{}
}

func (SimpleBillGenerator) Generate(cart cart.Cart, total float64) (Bill, error) {
	return NewBill(cart, total)
}
