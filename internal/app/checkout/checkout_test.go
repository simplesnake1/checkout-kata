package checkout

import (
	"testing"
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
		name     string
		checkout *checkout
		expected int
		errorMsg string
	}

	tests := []Test{
		{
			name:     "GetTotalPrice returns calculated value of price for 1 sku in basket",
			checkout: GetTestCheckout(map[string]int{"A": 1}),
			expected: 1,
			errorMsg: "when calculator has returned 1 for that sku",
		},
		{
			name:     "GetTotalPrice returns combined calculated value of price for 2 skus in basket",
			checkout: GetTestCheckout(map[string]int{"A": 1, "B": 2}),
			expected: 5,
			errorMsg: "when calculator has returned 5 for those skus",
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

func GetTestCalculator(sku string, count int) int {
	if sku == "A" {
		return 1 * count
	} else if sku == "B" {
		return 2 * count
	} else {
		return 0
	}
}

func GetTestCheckout(b map[string]int) *checkout {
	return &checkout{basket: b, getPrice: GetTestCalculator}
}
