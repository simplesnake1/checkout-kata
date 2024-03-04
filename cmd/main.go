package main

import (
	"fmt"

	"github.com/simplesnake1/checkout-kata/internal/app/checkout"
	"github.com/simplesnake1/checkout-kata/internal/app/pricing"
)

func main() {
	pl := map[string]pricing.Pricing{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialThreshold: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialThreshold: 2},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}
	calc := pricing.NewCalculator(pl)
	c := checkout.NewCheckout(calc.GetPrice)
	if c != nil {
		fmt.Println("Checkout created.")
	}

	c.Scan("B")
	c.Scan("A")
	c.Scan("B")

	fmt.Println(c.GetTotalPrice())
}
