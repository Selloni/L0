// Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	exOne()
}

// Делаем по схожему принцепу
// По истечению промежутка времени (context.WithTimeout)
func exOne() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel() // очищает ресурсы связанные с контекстом
	input := make(chan int)

	wg.Add(1)
	go worker(ctx, &wg, input)

	go write(ctx, input)
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, input <-chan int) {
	defer wg.Done()
	for {
		select { // слушаем не пришел ли сигнал
		case <-ctx.Done(): // если считываем сигнал, завершаем работу
			return // если не завершить канал будет ждать ввода
		default:
			if val, ok := <-input; ok { // считываем пока в канал поступают данные
				time.Sleep(1 * time.Second)
				fmt.Printf("%d seconds out of 10\n", val)
			}
		}
	}
}

func write(ctx context.Context, input chan<- int) {
	defer close(input) // выполняеться при выходе из функции
	i := 0
	for {
		select {
		case <-ctx.Done(): // завершение работы
			return
		default:
			i++
			input <- i // записываеи данные в канал
		}
	}
}
