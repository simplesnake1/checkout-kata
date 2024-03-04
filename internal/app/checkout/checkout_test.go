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
			if c.getPrice == nil {
				t.Fatalf("NewCheckout should create a new Checkout struct that has been passed a GetPriceFunc.")
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
			name:     "Mocked GetPriceFunc - GetTotalPrice returns calculated value of price for 1 sku in basket",
			checkout: GetTestCheckout(map[string]int{"A": 1}),
			expected: 1,
			errorMsg: "when calculator has returned 1 for that sku",
		},
		{
			name:     "Mocked GetPriceFunc - GetTotalPrice returns combined calculated value of price for 2 skus in basket",
			checkout: GetTestCheckout(map[string]int{"A": 1, "B": 2}),
			expected: 5,
			errorMsg: "when calculator has returned 5 for those skus",
		},
		{
			name:     "Calculator GetPriceFunc - GetTotalPrice returns correct value as per readme table 1 A and 2 B",
			checkout: &checkout{basket: map[string]int{"A": 1, "B": 2}, getPrice: GetCalculatorPriceFunc()},
			expected: 95,
			errorMsg: "when actual calculator with pricing list from readme has been used",
		},
		{
			name:     "Calculator GetPriceFunc - GetTotalPrice returns correct value as per readme table 3 A and 1 B",
			checkout: &checkout{basket: map[string]int{"A": 3, "B": 1}, getPrice: GetCalculatorPriceFunc()},
			expected: 160,
			errorMsg: "when actual calculator with pricing list from readme has been used",
		},
		{
			name:     "Calculator GetPriceFunc - GetTotalPrice returns correct value as per readme table 3 A and 2 B",
			checkout: &checkout{basket: map[string]int{"A": 3, "B": 2}, getPrice: GetCalculatorPriceFunc()},
			expected: 175,
			errorMsg: "when actual calculator with pricing list from readme has been used",
		},
		{
			name:     "Calculator GetPriceFunc - GetTotalPrice returns correct value as per readme table 3 A, 2 B, 3 C and 3 D",
			checkout: &checkout{basket: map[string]int{"A": 3, "B": 2, "C": 3, "D": 3}, getPrice: GetCalculatorPriceFunc()},
			expected: 280,
			errorMsg: "when actual calculator with pricing list from readme has been used",
		},
		{
			name:     "Calculator GetPriceFunc - GetTotalPrice returns correct value as per readme table 3 A, 2 B, 3 C and 3 D and 1 E which does not exist",
			checkout: &checkout{basket: map[string]int{"A": 3, "B": 2, "C": 3, "D": 3, "E": 0}, getPrice: GetCalculatorPriceFunc()},
			expected: 280,
			errorMsg: "when actual calculator with pricing list from readme has been used",
		},
		{
			name:     "Calculator GetPriceFunc - GetTotalPrice returns correct value as per readme table 7 A, 2 B, 3 C and 3 D and 1 E which does not exist",
			checkout: &checkout{basket: map[string]int{"A": 7, "B": 2, "C": 3, "D": 3, "E": 0}, getPrice: GetCalculatorPriceFunc()},
			expected: 460,
			errorMsg: "when actual calculator with pricing list from readme has been used",
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

func GetCalculatorPriceFunc() GetPriceFunc {
	pl := map[string]pricing.Pricing{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialThreshold: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialThreshold: 2},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}

	return pricing.NewCalculator(pl).GetPrice
}
