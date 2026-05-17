package main

import (
	"bath_to_text/converter"
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	inputs := []decimal.Decimal{
		decimal.NewFromFloat(1234),
		decimal.NewFromFloat(33333.75),
	}

	for _, input := range inputs {
		fmt.Println("Input:", input)
		fmt.Println("Output:", converter.ThaiBahtText(input))
		fmt.Println()
	}
}
