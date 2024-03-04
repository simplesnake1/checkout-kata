package pricing

type calculator struct {
	pricingList map[string]Pricing
}

func NewCalculator(pl map[string]Pricing) *calculator {
	c := new(calculator)

	c.pricingList = pl

	return c
}

func (c *calculator) GetPrice(sku string, count int) (price int) {
	return -1
}
