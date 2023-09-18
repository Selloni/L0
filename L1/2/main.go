package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	//ExampleOne()
	fmt.Println("-----")
	ExampleTwo()
}

// Сделано с помощью мьютекса,
// обычный счетчик, который блокирует мьютекс
func ExampleOne() {
	var wg sync.WaitGroup
	nums := []int{2, 4, 6, 8, 10}
	for _, i := range nums {
		wg.Add(1)
		go func(i int) {
			res := math.Pow(float64(i), 2)
			fmt.Println(res)
			defer wg.Done()
		}(i)
		wg.Wait()
	}
} //
// - не встраивайте мьютекс в структуру
// - не храните ссылку на мьютекс в поле структуры
// - методы с мьютекс, должны иметь ссылочный ресивер

// todo chan
func ExampleTwo() {
	nums := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	go func() {
		for _, i := range nums {
			println(i)
			ch <- i
		}
		close(ch)
	}()
	squaring(ch)
}

func squaring(ch chan int) {
	res := math.Pow(float64(<-ch), 2)
	fmt.Println(res)
}
