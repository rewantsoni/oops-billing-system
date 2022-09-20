package calculator

import (
	"github.com/stretchr/testify/assert"
	"oops-billing-system/cart"
	"testing"
)

func TestCreateNewCalculationSystem(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (CalculationSystem, error)
		expectedResult func() CalculationSystem
		expectedError  error
	}{
		{
			name: "Should create a new calculation system successfully",
			actualResult: func() (CalculationSystem, error) {
				return NewCalculationSystem()
			},
			expectedResult: func() CalculationSystem {
				return SimpleCalculationSystem{}
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

func TestComputeTotal(t *testing.T) {
	testCases := []struct {
		name           string
		actualResult   func() (float64, error)
		expectedResult float64
		expectedError  error
	}{
		{
			name: "Should return the expected value for ComputeTotal",
			actualResult: func() (float64, error) {
				cart, _ := cart.NewCartWithTwoItems()
				calculationSystem, _ := NewCalculationSystem()
				return calculationSystem.ComputeTotal(cart)
			},
			expectedResult: 64.0,
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
