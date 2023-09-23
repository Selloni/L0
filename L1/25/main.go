// Реализовать собственную функцию sleep.
// todo: добавить еще реализаций
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	SleepOne(3)
	fmt.Println("end")
}

// простая реализация на контекстах
func SleepOne(sec time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), sec*time.Second) // через sec* отправим в канал сигнал
	defer cancel()
	work(ctx)
}

func work(ctx context.Context) {
	for {
		select {
		case <-ctx.Done(): // программа в цикле ждет чтения из канал
			return
		}
	}
}
