package calculator

type Calculator struct {
	adder Adder
}

func NewCalculator(adder Adder) Calculator {
	return Calculator{
		adder,
	}
}

func (c Calculator) Add(x, y int) int {
	return c.adder.Add(x, y)
}
