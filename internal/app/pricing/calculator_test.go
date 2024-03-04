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
			name:       "Error Handling - Ignores price of item when it is not in the pricing list.",
			calculator: GetTestCalculator(),
			sku:        "E",
			count:      1,
			expected:   0,
			errorMsg:   "as E is not an existing SKU in the pricing list",
		},
		{
			name:       "Unit Price - Calculates the price for 1 item of A",
			calculator: GetTestCalculator(),
			sku:        "A",
			count:      1,
			expected:   50,
			errorMsg:   "as the Unit Price of A is 50, and there is a count of 1",
		},
		{
			name:       "Unit Price - Calculates the price for 2 items of A",
			calculator: GetTestCalculator(),
			sku:        "A",
			count:      2,
			expected:   100,
			errorMsg:   "as the Unit Price of A is 50, and there is a count of 2",
		},
		{
			name:       "Unit Price - Calculates the price for 1 item of B",
			calculator: GetTestCalculator(),
			sku:        "B",
			count:      1,
			expected:   30,
			errorMsg:   "as the Unit Price of A is 30, and there is a count of 1",
		},
		{
			name:       "Unit Price - Calculates the price for 2 items of B",
			calculator: GetTestCalculator(),
			sku:        "B",
			count:      2,
			expected:   60,
			errorMsg:   "as the Unit Price of A is 30, and there is a count of 2",
		},
		{
			name:       "Unit Price - Calculates the price for 2 items of A",
			calculator: GetTestCalculator(),
			sku:        "A",
			count:      2,
			expected:   100,
			errorMsg:   "as the Unit Price of A is 50, and there is a count of 2",
		},
		{
			name:       "Special Price - Calculates the price for 3 items of A",
			calculator: GetTestCalculator(),
			sku:        "A",
			count:      3,
			expected:   130,
			errorMsg:   "as the Special Price of A is 130, and there is a count of 3 which matches the Special Threshold",
		},
		{
			name:       "Special Price - Calculates the price for 6 items of A",
			calculator: GetTestCalculator(),
			sku:        "A",
			count:      6,
			expected:   260,
			errorMsg:   "as the Special Price of A is 130, and there is a count of 6 which matches the Special Threshold exactly twice",
		},
		{
			name:       "Special Price With Remainder - Calculates the price for 7 items of A",
			calculator: GetTestCalculator(),
			sku:        "A",
			count:      7,
			expected:   310,
			errorMsg:   "as the Unit Price of A is 50, the Special Price of A is 130, and there is a count of 7 which matches the Special Threshold twice, with 1 remainder",
		},
		{
			name:       "Special Price - Calculates the price for 2 items of B",
			calculator: GetTestCalculator(),
			sku:        "B",
			count:      2,
			expected:   45,
			errorMsg:   "as the Special Price of B is 45, and there is a count of 2 which matches the Special Threshold",
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
