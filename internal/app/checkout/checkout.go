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

func NewCheckout(GetPrice GetPriceFunc) *checkout {
	c := new(checkout)

	c.basket = make(map[string]int)

	return c
}

func (c *checkout) Scan(item string) {
	c.basket[item]++
}

func (c *checkout) GetTotalPrice() int {
	return -1
}
