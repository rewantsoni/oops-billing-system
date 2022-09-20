package bill

import (
	"github.com/stretchr/testify/assert"
	"oops-billing-system/calculator"
	"oops-billing-system/cart"
	"testing"
)

func TestNewBillingSystem(t *testing.T) {

	mcs := &calculator.MockCalculationSystem{}
	mbg := &MockBillGenerator{}
	testCases := []struct {
		name           string
		actualResult   func() BillingSystem
		expectedResult BillingSystem
	}{
		{
			name: "Should create a new billing system",
			actualResult: func() BillingSystem {
				return NewBillingSystem(mcs, mbg)
			},
			expectedResult: BillingSystem{
				mcs,
				mbg,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, res, testCase.expectedResult)
		})
	}
}

func TestGenerateBillFromBillingSystem(t *testing.T) {

	mcs := &calculator.MockCalculationSystem{}
	mbg := &MockBillGenerator{}
	cart, _ := cart.NewCartWithTwoItems()
	mcs.On("ComputeTotal", cart).Return(float64(64), nil)
	b, _ := NewBill(cart, float64(64))
	mbg.On("Generate", cart, float64(64)).Return(b, nil)
	billingSystem := NewBillingSystem(mcs, mbg)

	testCases := []struct {
		name           string
		actualResult   func() (Bill, error)
		expectedResult func() Bill
		expectedError  error
	}{
		{
			name: "Should generate a new bill",
			actualResult: func() (Bill, error) {
				return billingSystem.GenerateBill(cart)
			},
			expectedResult: func() Bill {
				bill, _ := NewBill(cart, 64)
				return bill
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult())
			}
		})
	}
}
