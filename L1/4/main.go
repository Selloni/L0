package main

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	fmt.Println("Напиши количество воркеров")
	var N int
	fmt.Fscan(os.Stdin, &N)
	exOne(N)
}

// По явному сигналу отмены (context.WithCancel)
// По истечению промежутка времени (context.WithTimeout)
// По наступлению временной отметки или дедлайна (context.WithDeadline)
func exOne(N int) {
	var wg sync.WaitGroup
	ctx := context.Background() // родительский контест
	input := make(chan int)
	for i := 1; i <= N; i++ {
		wg.Add(1)
		go worker(ctx, &wg, input)
	}
	go write(ctx, input)
	wg.Wait()
}

func worker(ctx context.Context, wg *sync.WaitGroup, input <-chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			if val, ok := <-input; ok {
				time.Sleep(1 * time.Second)
				fmt.Println(val)
			}
		}
	}
}

func write(ctx context.Context, input chan<- int) {
	defer close(input)
	i := 1
	for {
		select {
		case <-ctx.Done():
			return
		default:
			i *= 2
			input <- i
			time.Sleep(1 * time.Second)
		}

	}
}
