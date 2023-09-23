package main

import (
	"fmt"
	"math/big" // специальный пакет дял работы с большими числаи
)

func main() {
	a := big.NewInt(1234567890123456789)
	b := big.NewInt(9876543210987654)
	fmt.Println(new(big.Int).Add(a, b)) // метод сложения
	fmt.Println(new(big.Int).Div(a, b)) // деление
	fmt.Println(new(big.Int).Mul(a, b)) // умножение
	fmt.Println(new(big.Int).Sub(a, b)) // вычитание
}
