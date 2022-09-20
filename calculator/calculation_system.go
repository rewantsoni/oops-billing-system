package calculator

import "oops-billing-system/cart"

type CalculationSystem interface {
	ComputeTotal(cart.Cart) (float64, error)
}

type SimpleCalculationSystem struct{}

func NewCalculationSystem() (CalculationSystem, error) {
	return SimpleCalculationSystem{}, nil
}

func (SimpleCalculationSystem) ComputeTotal(cart cart.Cart) (float64, error) {
	total := 0.0
	for item, quantity := range cart.GetItemWithQuantity() {
		total += item.GetPrice() * float64(quantity)
	}
	return total, nil
}
