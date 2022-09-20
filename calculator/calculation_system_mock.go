package calculator

import (
	"github.com/stretchr/testify/mock"
	"oops-billing-system/cart"
)

type MockCalculationSystem struct {
	mock.Mock
}

func (mcs *MockCalculationSystem) ComputeTotal(cart cart.Cart) (float64, error) {
	args := mcs.Called(cart)
	return args.Get(0).(float64), args.Error(1)
}
