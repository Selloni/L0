// Реализовать собственную функцию sleep.
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	SleepOne(2)
	fmt.Println("sleep 2 sec")
	SleepTwo(2)
	fmt.Println("end 2")
}

// простая реализация на контекстах
func SleepOne(sec time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), sec*time.Second) // через sec* отправим в канал сигнал
	defer cancel()
	func() {
		for {
			select {
			case <-ctx.Done(): // программа в цикле ждет чтения из канал
				return
			}
		}
	}()
}

// работа с разницей по времени
func SleepTwo(sec time.Duration) {
	tt := time.Now()
	ll := tt.Add(sec * time.Second) // получаем время в которе выходим из цикла
	for time.Now().Before(ll) {     // сообщает, находится ли перед нами момент времени
	}
}
