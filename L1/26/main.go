package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(checkOne("abcAdkqL"))
	fmt.Println(checkOne("abc213"))
	fmt.Println(checkOne("afsd232"))
	fmt.Println(checkTwo("abcAdkqL"))
	fmt.Println(checkTwo("abc213"))
	fmt.Println(checkTwo("afsd232"))
}

func checkOne(str string) bool {
	mm := make(map[int32]bool)
	for _, i := range str {
		if mm[i] == true { // если символ уже был
			return false
		}
		mm[i] = true    // встречаем символ
		mm[i-32] = true // топорный способ, не оптимизированно
	}
	return true
}

// логика как в приемер вышел
func checkTwo(str string) bool {
	str = strings.ToUpper(str) // все символы переводим в верхний регистр
	mm := make(map[rune]struct{})
	for _, char := range str {
		if _, ok := mm[char]; ok {
			return false
		}
		mm[char] = struct{}{} // используем структуру, потому что она ни чегоне весит
	}
	return true
}
