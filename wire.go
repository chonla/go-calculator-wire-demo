package main

import (
	"calculator/calculator"

	"github.com/google/wire"
)

func InitializeCalculator() calculator.Calculator {
	wire.Build(calculator.NewCalculator, calculator.NewAdder)

	return calculator.Calculator{}
}
