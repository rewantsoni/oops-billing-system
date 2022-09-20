package item

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewItem(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (Item, error)
		expectedResult Item
		expectedError  error
	}{
		{
			name: "Should create a new Item successfully",
			actualResult: func() (Item, error) {
				return NewItemApple()
			},
			expectedResult: Item{
				name:  "apple",
				price: 12.0,
			},
		},
		{
			name: "Should not create a new item if the item name is empty",
			actualResult: func() (Item, error) {
				return NewItemWithoutName()
			},
			expectedError: errors.New("item name cannot be empty"),
		},
		{
			name: "Should not create a new item if the price is zero or negative",
			actualResult: func() (Item, error) {
				return NewItemWithZeroOrNegativePrice()
			},
			expectedError: errors.New("item price cannot be less than or equal to zero"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res, err := testCase.actualResult()
			assert.Equal(t, err, testCase.expectedError)
			if err == nil {
				assert.Equal(t, res, testCase.expectedResult)
			}
		})
	}
}

func TestGetName(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() string
		expectedResult string
	}{
		{
			name: "Should return the expected name for item",
			actualResult: func() string {
				apple, _ := NewItemApple()
				return apple.GetName()
			},
			expectedResult: "apple",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, res, testCase.expectedResult)
		})
	}
}

func TestGetPrice(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() float64
		expectedResult float64
	}{
		{
			name: "Should return the expected price for the item",
			actualResult: func() float64 {
				apple, _ := NewItemApple()
				return apple.GetPrice()
			},
			expectedResult: 12.0,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := testCase.actualResult()
			assert.Equal(t, res, testCase.expectedResult)
		})
	}
}
