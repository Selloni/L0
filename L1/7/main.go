// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

func main() {
	mp := make(map[int]int)
	// будем использовать мютекс для конкурентной записи
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func(i int) { // каждая горутина пытаеться записать в мапу
			defer wg.Done()
			mutex.Lock()   // создаем очередь, блокируем запись, если уже кто то записывает
			mp[i] = i * i  // запись
			mutex.Unlock() // разблокируем, дальше работатем конкурентно
		}(i)
	}
	wg.Wait()
	fmt.Println(mp)
}
