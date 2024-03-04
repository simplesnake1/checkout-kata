package checkout

import (
	"testing"
)

func TestCheckout_NewCheckout(t *testing.T) {
	type Test struct {
		name string
	}

	tests := []Test{
		{
			name: "Constructs a new checkout instance with a basket ready to go",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCheckout()
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
		scanned  []string
		expected map[string]int
	}

	tests := []Test{
		{
			name:     "Scan adds item to basket",
			scanned:  []string{"A"},
			expected: map[string]int{"A": 1},
		},
		{
			name:     "Scan adds item to basket, and increments count on subsequent calls with the same item",
			scanned:  []string{"A", "A"},
			expected: map[string]int{"A": 2},
		},
		{
			name:     "Scan adds item to basket, and creates new Key Value Pair on subsequent calls with different items",
			scanned:  []string{"A", "B"},
			expected: map[string]int{"A": 1, "B": 1},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCheckout()

			for _, item := range test.scanned {
				c.Scan(item)
			}

			for item, count := range test.expected {
				if c.basket[item] != count {
					t.Fatalf("Basket should contain %d for %s", count, item)
				}
			}
		})

	}
}
