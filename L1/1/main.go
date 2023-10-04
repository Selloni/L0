package main

import (
	"fmt"
)

/*
Реализовать встраивание методов в структуре Action
от родительской структуры Human (аналог наследования).
*/
type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human // встраиваемая структура
}

func (h Human) Say() {
	fmt.Printf("My name %s\n", h.Name)
}

func (a *Action) DoIt() {
	a.Age++
	fmt.Printf("I'll be older next year %d\n", a.Age)
}

//func (a *Action) Say() { \\ переопределение
//	fmt.Printf("test\n")
//}

func main() {
	a := Action{
		Human{
			Name: "Grandpat",
			Age:  24,
		},
	}
	a.Say() // вызывем метод родителя
	a.DoIt()

}
