package main

import "fmt"

func main() {
	var ss int64 = 5
	fmt.Println(setBit(ss, 1, 1))
}

func setBit(ss int64, i int, bit int) int64 {
	var tmp int64 = 1 // маска
	tmp = tmp << i    // сдвиг по маске
	if bit == 0 {     //  устанавливаемое значение
		return tmp ^ ss
	}
	return tmp | ss //
}

//todo - чательнее проверить , напсиать тесты
