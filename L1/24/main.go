package main

import "fmt"

type point struct { // до этого метода не добраться из вне
	x int
	y int
}

func NewPoint(a int, b int) *point {
	return &point{
		x: a, // пердаем значеие в структуру
		y: b,
	}
}

func main() {
	fmt.Printf("Дистанция %d", distans())
}

func distans() int { // функция расчета
	pp := NewPoint(8, 2)
	return pp.x - pp.y
}
