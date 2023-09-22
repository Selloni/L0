package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, 11, -21.0, 32.5}
	sort.Float64s(arr)                // сортируем для удобства
	mm := make(map[float64][]float64) // хранилище
	join(arr, mm)
	fmt.Println(mm)
}

func join(arr []float64, mm map[float64][]float64) {
	var tmp []float64
	ten := 0
	for i := 1; i < len(arr); i++ {
		if arr[ten]+10 < arr[i] { // провряем входит ли в диапозон
			mm[arr[ten]] = tmp
			ten = i   // шаг с которым сравниваемся
			tmp = nil // зануление для повторного заполнения
		} else {
			tmp = append(tmp, arr[i]) // заполняме слаиз если входим в диапозон чисел
		}
	}
}
