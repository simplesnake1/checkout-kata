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
			name: "Constructs a new checkout instance",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := NewCheckout()
			if c == nil {
				t.Fatalf("NewCheckout should create a new Checkout struct and return a pointer to it.")
			}
		})
	}
}
