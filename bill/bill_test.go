package bill

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"oops-billing-system/cart"
	"testing"
)

func TestNewBill(t *testing.T) {

	testCases := []struct {
		name           string
		actualResult   func() (Bill, error)
		expectedResult func() Bill
		expectedError  error
	}{
		{
			name: "Should create a new Bill",
			actualResult: func() (Bill, error) {
				cartWithTwoItems, _ := cart.NewCartWithTwoItems()
				total := 1.0
				return NewBill(cartWithTwoItems, total)
			},
			expectedResult: func() Bill {
				cartWithTwoItems, _ := cart.NewCartWithTwoItems()
				total := 1.0
				return Bill{cart: cartWithTwoItems, total: total}
			},
		},
		{
			name: "Should not create a new Bill if the total is zero or negative",
			actualResult: func() (Bill, error) {
				cartWithTwoItems, _ := cart.NewCartWithTwoItems()
				total := 0.0
				return NewBill(cartWithTwoItems, total)
			},
			expectedError: errors.New("total cannot be zero or negative"),
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

func TestString(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() string
		expectedResult string
	}{
		{
			name: "Should create a new Bill",
			actualResult: func() string {
				cartWithTwoItems, _ := cart.NewCartWithTwoItems()
				total := 1.0
				bill, _ := NewBill(cartWithTwoItems, total)
				return bill.String()
			},
			expectedResult: "Item\tPrice\tQuantity\napple\t12.000000\t2\norange\t10.000000\t4\n\nTotal: 1.000000",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, res, testCase.expectedResult)

		})
	}
}
