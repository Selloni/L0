package main

import "fmt"

type point struct {
	x int
	y int
}

func NewPoint() *point {
	return &point{
		x: 2,
		y: 8,
	}
}

func main() {
	fmt.Printf("Дистанция %d", distans())
}

func distans() int {
	pp := NewPoint()
	return pp.x - pp.y
}
