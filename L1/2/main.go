package main

import (
	"fmt"
	"math"
	"sync"
)

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
	for _, i := range nums {
		wg.Add(1) // счетчик увеличиваме на один
		go func(i int) {
			res := math.Pow(float64(i), 2)
			fmt.Println(res)
			defer wg.Done() // уменьшаем на один
		}(i)
		wg.Wait()
	}
} //
// - не встраивайте мьютекс в структуру
// - не храните ссылку на мьютекс в поле структуры
// - методы с мьютекс, должны иметь ссылочный ресивер
/////////////////////////////////////////////////////

// С помощью каналов
func ExampleTwo(nums []int) {
	ch := make(chan int)
	go func() {
		for _, i := range nums {
			ch <- i // запись в канал
		}
		close(ch)
	}() // этот блок кода в горутине, выполняеться в дргуом ядре
	for cc := range ch { // считываю с канала, пока он не будет закрыт
		res := cc * cc
		fmt.Println(res)
	}
} //
//Запись в неициализированный канал блокирует поток навсегда;
//Чтение из неинициализированного канала блокирует поток навсегда;
//Запись в закрытый канал вызывает панику;
//Чтение из закрытого канала даёт нулевое значение мгновенно.

// буферизированный канал
func ExampleThree(nums []int) {
	ch := make(chan int, len(nums)) // небольшое отличие,сразу указываю количество данных могу записать
	go func() {
		for _, i := range nums {
			ch <- i
		}
	}()
	for i := 0; i != len(nums); i++ { // считываю данные, с размером с буфер
		res := math.Pow(float64(<-ch), 2)
		fmt.Println(res)
	}
}
