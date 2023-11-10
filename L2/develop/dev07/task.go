package main

import (
	"fmt"
	"sync"
	"time"
)

/*
Реализуем паттерн Fan-In

=== Or channel ===

Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}

start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))
*/

func main() {

	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		//sig(2*time.Hour),
		//sig(5*time.Minute),
		sig(1*time.Second),
		//sig(1*time.Hour),
		sig(3*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))

}

// обьединяем каналы в один канал
func or(channels ...<-chan interface{}) <-chan interface{} {
	if len(channels) == 0 {
		return nil
	}
	done := make(chan interface{})
	wg := sync.WaitGroup{}
	wg.Add(len(channels)) // ожидаем выполненяи всех горутин
	for _, ch := range channels {
		go func(ch <-chan interface{}) {
			defer wg.Done()
			for c := range ch {
				done <- c // все пришедшие данные записываем в один канал
			}
		}(ch)
	}
	// как все горутины выполнятся закрываем канал
	go func() {
		wg.Wait()
		close(done)
	}()
	return done
}
