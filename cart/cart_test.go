package cart

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"oops-billing-system/item"
	"testing"
)

func TestNewCart(t *testing.T) {

	testCases := []struct {
		name           string
		actualResult   func() (Cart, error)
		expectedResult func() Cart
		expectedError  error
	}{
		{
			name: "Should create a new cart",
			actualResult: func() (Cart, error) {
				return NewCartWithTwoItems()
			},
			expectedResult: func() Cart {
				apple, _ := item.NewItemApple()
				orange, _ := item.NewItemOrange()
				return Cart{itemWithQuantity: map[item.Item]Quantity{
					apple:  2,
					orange: 4,
				}}
			},
		},
		{
			name: "Should not create a new cart if the items are nil",
			actualResult: func() (Cart, error) {
				return NewCart(nil)
			},
			expectedError: errors.New("item cannot be empty"),
		},
		{
			name: "Should not create a new cart if the item is empty",
			actualResult: func() (Cart, error) {
				return NewCartWithNoItems()
			},
			expectedError: errors.New("item cannot be empty"),
		},
		{
			name: "Should not create a new cart if the item quantity is zero or negative",
			actualResult: func() (Cart, error) {
				return NewCartWithZeroOrNegativeQuantity()
			},
			expectedError: errors.New("quantity cannot be negative or zero"),
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

func TestGetItemWithQuantity(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() map[item.Item]Quantity
		expectedResult func() map[item.Item]Quantity
	}{
		{
			name: "Should return map of items with quantity",
			actualResult: func() map[item.Item]Quantity {
				cart, _ := NewCartWithTwoItems()
				return cart.GetItemWithQuantity()
			},
			expectedResult: func() map[item.Item]Quantity {
				apple, _ := item.NewItemApple()
				orange, _ := item.NewItemOrange()
				return map[item.Item]Quantity{
					apple: 2, orange: 4,
				}
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, res, testCase.expectedResult())
		})
	}
}
