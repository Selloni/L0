package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(1234567890123456789)
	b := big.NewInt(9876543210987654)
	fmt.Println(new(big.Int).Add(a, b))
	fmt.Println(new(big.Int).Div(a, b))
	fmt.Println(new(big.Int).Mul(a, b))
	fmt.Println(new(big.Int).Sub(a, b))

}
