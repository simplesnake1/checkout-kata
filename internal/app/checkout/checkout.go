package checkout

type GetPriceFunc func(string, int) int

type ICheckout interface {
	Scan(item string)
	GetTotalPrice() int
}

type checkout struct {
	basket   map[string]int
	getPrice GetPriceFunc
}

func NewCheckout(getPrice GetPriceFunc) *checkout {
	c := new(checkout)

	c.basket = make(map[string]int)
	c.getPrice = getPrice

	return c
}

func (c *checkout) Scan(item string) {
	c.basket[item]++
}

func (c *checkout) GetTotalPrice() int {
	tp := 0
	for sku, count := range c.basket {
		tp += c.getPrice(sku, count)
	}

	return tp
}
