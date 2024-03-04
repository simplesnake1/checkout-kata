package pricing

import "testing"

func TestCalculator_NewCalculator(t *testing.T) {
	type Test struct {
		name string
	}

	tests := []Test{
		{
			name: "Constructs a new Calculator instance with a pricing list ready to go",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCalculator(GetTestPricingList())
			if c == nil {
				t.Fatalf("NewCalculator should create a new Calculator struct and return a pointer to it.")
			}
			if c.pricingList == nil {
				t.Fatalf("NewCalculator should create a new Calculator struct with a pricing list ready to go.")
			}
		})
	}
}

func TestCalculator_GetPrice(t *testing.T) {
	type Test struct {
		name       string
		calculator *calculator
		sku        string
		count      int
		expected   int
		errorMsg   string
	}

	tests := []Test{
		{
			name:       "Ignores price of item when it is not in the pricing list.",
			calculator: GetTestCalculator(),
			sku:        "E",
			count:      1,
			expected:   0,
			errorMsg:   "as E is not an existing SKU in the pricing list",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := test.calculator.GetPrice(test.sku, test.count)
			if r != test.expected {
				t.Fatalf("GetTotalPrice should return %d not %d %s", test.expected, r, test.errorMsg)
			}
		})
	}
}

func GetTestPricingList() map[string]Pricing {
	return map[string]Pricing{
		"A": {UnitPrice: 50, SpecialPrice: 130, SpecialThreshold: 3},
		"B": {UnitPrice: 30, SpecialPrice: 45, SpecialThreshold: 2},
		"C": {UnitPrice: 20},
		"D": {UnitPrice: 15},
	}
}

func GetTestCalculator() *calculator {
	return &calculator{pricingList: GetTestPricingList()}
}
