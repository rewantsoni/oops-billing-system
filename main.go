package main

import (
	"fmt"
	"oops-billing-system/bill"
	"oops-billing-system/calculator"
	"oops-billing-system/cart"
	"oops-billing-system/item"
)

func main() {
	apple, _ := item.NewItem("apple", 12.0)
	orange, _ := item.NewItem("orange", 15.0)
	pineapple, _ := item.NewItem("pineapple", 18.0)

	c, _ := cart.NewCart(map[item.Item]cart.Quantity{
		apple: 2, orange: 5, pineapple: 10,
	})

	cartTwo, _ := cart.NewCart(map[item.Item]cart.Quantity{
		apple: 10, orange: 28, pineapple: 50,
	})

	cal, _ := calculator.NewCalculationSystem()
	myBillGenerator := bill.NewBillGenerator()

	myBillingSystem := bill.NewBillingSystem(cal, myBillGenerator)

	billOne, _ := myBillingSystem.GenerateBill(c)
	billTwo, _ := myBillingSystem.GenerateBill(cartTwo)
	fmt.Println(billOne)
	fmt.Println(billTwo)

}
