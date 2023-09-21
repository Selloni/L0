// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func main() {
	ex1 := []int{2, 4, 8, 16}
	ex2 := []int{1, 2, 3, 4, 5, 6, 7}
	mm := make(map[int]bool) // удобный инструмент
	fillMap(ex1, mm)
	fmt.Println(check(ex2, mm))

}

func fillMap(ex []int, mm map[int]bool) {
	for _, num := range ex {
		mm[num] = true // создаем мапу, со значениями из первого масива
	}
}
func check(ex2 []int, mm map[int]bool) []int {
	var ansver []int
	for _, num := range ex2 {
		if mm[num] { // если у нас есть совпадение во втором масиве
			ansver = append(ansver, num) // добавляем в слайс
		}
	}
	return ansver
}
