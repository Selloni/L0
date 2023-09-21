// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные
// из второго канала должны выводиться в stdout
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 5, 7}
	ch1 := make(chan int)
	ch2 := make(chan int)
	go workerIn(nums, ch1)
	go workerOut(ch1, ch2)
	for write := range ch2 {
		fmt.Println(write) // выводим пока канал не закрыт
	}
}

func workerIn(nums []int, ch chan int) {
	for _, num := range nums { // пробегаемся по массиву
		ch <- num // записываем в канал
	}
	close(ch) // не закрыв канал, получим дедлок
}

func workerOut(in chan int, out chan int) {
	for num := range in { // пока канал не закрыт считываем
		out <- num * 2 // записываем в канал вывода
	}
	close(out)
}
