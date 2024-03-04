package checkout

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {
	basket map[string]int
}

func NewCheckout() *Checkout {
	c := new(Checkout)
	return c
}
