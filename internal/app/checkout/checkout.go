package checkout

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type checkout struct {
	basket map[string]int
}

func NewCheckout() *checkout {
	c := new(checkout)

	c.basket = make(map[string]int)

	return c
}

func (c *checkout) Scan(item string) {
	c.basket[item]++
}
