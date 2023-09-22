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
		return arr
	}
	var (
		less []int
		more []int
	)
	pivot := len(arr) - 1
	more = append(more, arr[pivot])
	for c := 0; c < pivot; c++ {
		if arr[c] <= arr[pivot] {
			less = append(less, arr[c])
		} else {
			more = append(more, arr[c])
		}
	}
	return append(quickSort(less), quickSort(more)...)
}
