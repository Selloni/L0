package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	c  int
	mx sync.Mutex
}

func main() {
	var count Counter
	var wg sync.WaitGroup
	arr := []int{2, 16, 4, 5, 8, 22, -3}
	for _, i := range arr {
		wg.Add(1)
		go func(k int) {
			fmt.Println(k * k) // иммитация действия
			count.Increment()  // внутри стурктуры счетчик увеличивается
			wg.Done()
		}(i)
	}
	wg.Wait()
	defer fmt.Printf("итоговое значение счетчика %d", count.c)
}

func (cnt *Counter) Increment() {
	cnt.mx.Lock() // блокируем запись для безопасной инкриментации
	cnt.c++
	cnt.mx.Unlock()
}
