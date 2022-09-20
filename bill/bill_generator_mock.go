package bill

import (
	"github.com/stretchr/testify/mock"
	"oops-billing-system/cart"
)

type MockBillGenerator struct {
	mock.Mock
}

func (mbg *MockBillGenerator) Generate(cart cart.Cart, total float64) (Bill, error) {
	args := mbg.Called(cart, total)
	return args.Get(0).(Bill), args.Error(1)
}
