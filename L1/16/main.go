// Реализовать быструю сортировку массива (quicksort)
// встроенными методами языка.
package main

import (
	"fmt"
)

func main() {
	arr := []int{2, 17, 27, 11, -9, 14}
	fmt.Println(quickSort(arr))
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr // завершение рекурсии
	}
	var (
		less []int
		more []int
	)
	pivot := len(arr) - 1           // сравнение идет по последниму элементу
	more = append(more, arr[pivot]) // он запсиываеться в начало массива где значения относительно него больше
	for c := 0; c < pivot; c++ {
		if arr[c] <= arr[pivot] {
			less = append(less, arr[c]) // массив с меньшими значениями относительно pivot
		} else {
			more = append(more, arr[c]) // массив с большими значениями
		}
	}
	return append(quickSort(less), quickSort(more)...)
	// рекурсивный вызов функции
	// в завершении рекурсии будет аппендиться отсортированная строка
	// возможные пробелмы, постоянное создание массивов
}
