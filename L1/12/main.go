package main

import "fmt"

func main() {
	word := []string{"cat", "cat", "dog", "cat", "tree"}
	plenty := make(map[string]int)
	for _, w := range word {
		plenty[w] += 1
	}
	fmt.Println(plenty)
}
