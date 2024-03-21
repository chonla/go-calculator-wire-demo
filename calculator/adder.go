package calculator

type Adder struct{}

func NewAdder() Adder {
	return Adder{}
}

func (a Adder) Add(x, y int) int {
	return x + y
}
