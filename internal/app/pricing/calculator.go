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
	_, exists := c.pricingList[sku]

	if exists {
		return -1
	}

	return
}
