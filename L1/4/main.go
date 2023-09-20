// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят
// в stdout. Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C.
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Напиши количество воркеров")
	var N int
	fmt.Fscan(os.Stdin, &N) // считываем с консоли
	exOne(N)
}

// По явному сигналу отмены (context.WithCancel)
// По истечению промежутка времени (context.WithTimeout)
// По наступлению временной отметки или дедлайна (context.WithDeadline)
func exOne(N int) {
	var wg sync.WaitGroup
	// syscall.SIGINT - прирывания ctr c
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM) // контекст для изящного завершения
	defer stop()                                                                             // очищает ресурсы связанные с контекстом
	input := make(chan int)
	for i := 1; i <= N; i++ {
		wg.Add(1)
		go worker(ctx, &wg, input) // паралельно будет работать N функций
	}
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
				time.Sleep(2 * time.Second) // для наглядности усыпляем вывод
				fmt.Println(val)
			}
		}
	}
}

func write(ctx context.Context, input chan<- int) {
	defer close(input) // выполняеться при выходе из функции
	i := 1
	for {
		select {
		case <-ctx.Done(): // завершение работы
			return
		default:
			i *= 2
			input <- i // записываеи данные в канал
		}
	}
}
