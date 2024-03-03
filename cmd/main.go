package main

import (
	"fmt"

	"github.com/simplesnake1/checkout-kata/internal/app/checkout"
)

func main() {
	c := checkout.NewCheckout()
	if c != nil {
		fmt.Println("Checkout created.")
	}
}
