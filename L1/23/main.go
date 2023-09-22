package main

import "fmt"

func main() {
	alf := []string{"a", "b", "c", "d", "f"}
	fmt.Println(delete(alf, 2))

}

func delete(alf []string, i int) (tmp []string) {
	tmp = alf[:i]
	tmp = append(tmp, alf[i+1:]...)
	return
}
