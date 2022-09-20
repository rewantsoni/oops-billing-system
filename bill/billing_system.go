package bill

import (
	"oops-billing-system/calculator"
	"oops-billing-system/cart"
)

type BillingSystem struct {
	calculationSystem calculator.CalculationSystem
	billGenerator     BillGenerator
}

func NewBillingSystem(calculationSystem calculator.CalculationSystem, billGenerator BillGenerator) BillingSystem {
	return BillingSystem{
		calculationSystem: calculationSystem,
		billGenerator:     billGenerator,
	}
}

func (b BillingSystem) GenerateBill(cart cart.Cart) (Bill, error) {
	total, err := b.calculationSystem.ComputeTotal(cart)
	if err != nil {
		return Bill{}, err
	}
	return b.billGenerator.Generate(cart, total)
}
