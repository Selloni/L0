package main

import "fmt"

func main() {
	arr := []int{-3, 1, 2, 3, 4, 15, 17, 26, 102}
	fmt.Println(Search(3, arr))
	fmt.Println(Search(10, arr))
}

// ищем элемент в отсортированном массиве,
// двигаемся относительно середены, постоянно делим массив на две части
// до момента пока не найдем нужное число
func Search(num int, arr []int) bool {
	if num == arr[len(arr)/2] {
		return true
	} else if len(arr) < 2 {
		return false
	} else if num < arr[len(arr)/2] {
		return Search(num, arr[:len(arr)/2]) // рекурсивно ищем в левой части если значение меньше
	} else {
		return Search(num, arr[len(arr)/2:]) // если значение больше
	}
}
