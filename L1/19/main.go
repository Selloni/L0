package main

import "fmt"

func main() {
	str := "что угодно, i dont refuse to you"
	rune := []rune(str)
	fmt.Println(string(reverse(rune))) // руны переводи в троку
}

func reverse(str []rune) []rune {
	flip := make([]rune, len(str)) // буффер дял заполнения
	lenght := len(str)
	for i, j := 0, lenght-1; i < lenght; i, j = i+1, j-1 {
		flip[i] = str[j] // простейщая перезапись в обратно порядке
	}
	return flip
}
