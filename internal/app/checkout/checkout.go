package checkout

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type Checkout struct {
}

func NewCheckout() *Checkout {
	c := new(Checkout)
	return c
}
