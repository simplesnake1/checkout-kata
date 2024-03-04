package checkout

import (
	"testing"

	"github.com/simplesnake1/checkout-kata/internal/app/pricing"
)

func TestCheckout_NewCheckout(t *testing.T) {
	type Test struct {
		name         string
		getPriceFunc GetPriceFunc
	}

	tests := []Test{
		{
			name:         "Constructs a new checkout instance with a basket ready to go",
			getPriceFunc: func(s string, i int) int { return -1 },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCheckout(test.getPriceFunc)
			if c == nil {
				t.Fatalf("NewCheckout should create a new Checkout struct and return a pointer to it.")
			}
			if c.basket == nil {
				t.Fatalf("NewCheckout should create a new Checkout struct with a basket ready to go.")
			}
		})
	}
}

func TestCheckout_Scan(t *testing.T) {
	type Test struct {
		name     string
		checkout *checkout
		scanned  []string
		expected map[string]int
	}

	tests := []Test{
		{
			name:     "Scan adds item to basket",
			checkout: &checkout{basket: make(map[string]int)},
			scanned:  []string{"A"},
			expected: map[string]int{"A": 1},
		},
		{
			name:     "Scan adds item to basket, and increments count on subsequent calls with the same item",
			checkout: &checkout{basket: make(map[string]int)},
			scanned:  []string{"A", "A"},
			expected: map[string]int{"A": 2},
		},
		{
			name:     "Scan adds item to basket, and creates new Key Value Pair on subsequent calls with different items",
			checkout: &checkout{basket: make(map[string]int)},
			scanned:  []string{"A", "B"},
			expected: map[string]int{"A": 1, "B": 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			for _, item := range test.scanned {
				test.checkout.Scan(item)
			}

			for item, count := range test.expected {
				if test.checkout.basket[item] != count {
					t.Fatalf("Basket should contain %d for %s", count, item)
				}
			}
		})

	}
}

func TestCheckout_GetTotalPrice(t *testing.T) {
	type Test struct {
		name        string
		checkout    *checkout
		pricingList map[string]pricing.Pricing
		expected    int
		errorMsg    string
	}

	tests := []Test{
		{
			name:        "GetTotalPrice returns 0 when basket does not contain any item with valid sku",
			checkout:    &checkout{basket: map[string]int{"E": 1}},
			pricingList: GetTestPricingList(),
			expected:    0,
			errorMsg:    "when basket does not contain any item with valid sku",
		},
		{
			name:        "GetTotalPrice gets Unit Price of 1 item when only 1 item is in basket",
			checkout:    &checkout{basket: map[string]int{"A": 1}},
			pricingList: GetTestPricingList(),
			expected:    50,
			errorMsg:    "when these items are in the basket",
		},
		{
			name:        "GetTotalPrice gets Unit Price of 1 item when only 1 item that exists in the pricing list is in basket",
			checkout:    &checkout{basket: map[string]int{"E": 1, "A": 1}},
			pricingList: GetTestPricingList(),
			expected:    50,
			errorMsg:    "when these items are in the basket and 1 does not exist in the pricing list",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := test.checkout.GetTotalPrice()
			if r != test.expected {
				t.Fatalf("GetTotalPrice should return %d not %d %s", test.expected, r, test.errorMsg)
			}
		})
	}
}

func GetTestPricingList() map[string]pricing.Pricing {
	return map[string]pricing.Pricing{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialThreshold: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialThreshold: 2},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}
}
