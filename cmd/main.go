package main

import (
	"fmt"

	"github.com/simplesnake1/checkout-kata/internal/app/checkout"
)

func main() {
	c := checkout.NewCheckout(func(s string, i int) int { return -1 })
	if c != nil {
		fmt.Println("Checkout created.")
	}

	c.Scan("B")
	c.Scan("A")
	c.Scan("B")
}
