package main

import "fmt"

func main() {
	word := []string{"cat", "cat", "dog", "cat", "tree"}
	plenty := make(map[string]int)
	// ключем будет само слово, значение будет увеличиваться
	// каждый раз когда повторно встречаеться слово
	for _, w := range word {
		plenty[w] += 1
	}
	fmt.Println(plenty)
}
