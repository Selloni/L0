package main

import "fmt"

/*
	Фабричный метод
todo
*/

// выбераем какое строение будем строить
type BuildingI interface {
	Build()
}

type BuildS struct {
	develop BuildingI
}

type Nora struct {
	Door   int
	Window int
}

func (n *Nora) Build() {
	n.Door = 0
	n.Window = 0
	fmt.Println("Нора")
}

type Sauna struct {
	Door   int
	Window int
}

func (s *Sauna) Build() {
	s.Door = 1
	s.Window = 0
	fmt.Println("Сауна")
}

type House struct {
	Door   int
	Window int
}

func (h *House) Build() {
	h.Door = 1
	h.Window = 2
	fmt.Println("Дом")
}

func main() {
	home := House{}
	bb := BuildS{&home}
	bb.develop.Build()
	fmt.Println(home.Window, home.Door)
}
