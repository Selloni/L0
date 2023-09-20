// Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

func main() {
	exOne()
	exTwo()

}

func exOne() {
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

// sync.Map — это примерно как map[any]any, но она готова к конкуретному доступу,
// т.е. её нет необходимости обкладывать блокировками.
func exTwo() {
	var mm sync.Map
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			mm.Store(k, k*2) // заполнени
		}(i)
	}
	wg.Wait()
	for k := 1; k <= 3; k++ {
		fmt.Println(mm.Load(k)) // вывод содержимого по ключу
	}
}
