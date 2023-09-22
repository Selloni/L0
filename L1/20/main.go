package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "что угодно, I won't give up on you"
	words := strings.Fields(str) // сплитем строку на слова
	var tmp []string
	for i := len(words) - 1; i >= 0; i-- {
		tmp = append(tmp, words[i]) // перезаписываем сова реверстном порядке
	}
	fmt.Println(tmp)
}
