package main

import "fmt"

func main() {
	n, m := 5, 7 // получаем разницу между числами и оперируем ею
	m = m - n
	n = n + m
	m = n - m
	fmt.Println(n, m)
}
