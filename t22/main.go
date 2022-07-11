package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := "13196492172837087230472324"
	b := "834620348720847208472084720"

	first := &big.Int{}
	first.SetString(a, 10)
	second := &big.Int{}
	second.SetString(b, 10)

	fmt.Println(new(big.Int).Add(first, second))
	fmt.Println(new(big.Int).Sub(first, second))
	fmt.Println(new(big.Int).Mul(first, second))
	fmt.Println(new(big.Int).Div(first, second))
}
