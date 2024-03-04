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
	p, exists := c.pricingList[sku]

	if exists {
		if p.SpecialThreshold > 0 && count >= p.SpecialThreshold {
			price += (count / p.SpecialThreshold) * p.SpecialPrice
			count %= p.SpecialThreshold
		}

		price += count * p.UnitPrice
	}

	return
}
