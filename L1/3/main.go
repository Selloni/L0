// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"math"
	"sync"
)

// Таже логика что и во втором задании
func main() {
	nums := []int{2, 4, 6, 8, 10}
	ExampleOne(nums)
	fmt.Println("-----")
	ExampleTwo(nums)
	fmt.Println("-----")
	ExampleThree(nums)
}

// Сделано с помощью мьютекса,
// обычный счетчик, который блокирует мьютекс
func ExampleOne(nums []int) {
	var wg sync.WaitGroup
	var res float64
	for _, i := range nums {
		wg.Add(1)
		go func(i int) {
			res += math.Pow(float64(i), 2)
			defer wg.Done()
		}(i)
		wg.Wait()
	}
	fmt.Println(res)
}

// С помощью каналов
func ExampleTwo(nums []int) {
	var res int
	ch := make(chan int)
	go func() {
		for _, i := range nums {
			ch <- i // запись в канал
		}
		close(ch)
	}() // этот блок кода в горутине, выполняеться в дргуом ядре
	for cc := range ch { // считываю с канала, пока он не будет закрыт
		res += cc * cc
	}
	fmt.Println(res)
}

// буферизированный канал
func ExampleThree(nums []int) {
	var res float64
	ch := make(chan int, len(nums)) // небольшое отличие,сразу указываю количество данных могу записать
	go func() {
		for _, i := range nums {
			ch <- i
		}
	}()
	for i := 0; i != len(nums); i++ { // считываю данные, с размером с буфер
		res += math.Pow(float64(<-ch), 2)
	}
	fmt.Println(res)
}
