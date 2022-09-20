package bill

import (
	"github.com/stretchr/testify/assert"
	"oops-billing-system/calculator"
	"oops-billing-system/cart"
	"testing"
)

func TestCreateNewBillGenerator(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() BillGenerator
		expectedResult BillGenerator
	}{
		{
			name: "Should create a new Bill",
			actualResult: func() BillGenerator {
				return NewBillGenerator()
			},
			expectedResult: SimpleBillGenerator{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, res, testCase.expectedResult)
		})
	}
}

func TestGenerateBill(t *testing.T) {
	cartWithTwoItems, _ := cart.NewCartWithTwoItems()
	mcs := &calculator.MockCalculationSystem{}
	mcs.On("ComputeTotal", cartWithTwoItems).Return(68.0)
	testCases := []struct {
		name           string
		actualResult   func() (Bill, error)
		expectedResult func() Bill
		expectedError  error
	}{
		{
			name: "Should return the expected bill object",
			actualResult: func() (Bill, error) {
				billGenerator := NewBillGenerator()
				return billGenerator.Generate(cartWithTwoItems, 1)
			},
			expectedResult: func() Bill {
				newBill, _ := NewBill(cartWithTwoItems, 1)
				return newBill
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
