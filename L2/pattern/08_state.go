package main

import (
	"fmt"
)

/*
	Состояние
	тип: поведенческий

	 Шаги реализации
1. Определитесь с классом, который будет играть роль контекста. Это может быть как существующий класс,
	в котором уже есть зависимость от состояния, так и новый класс,
	если код состояний размазан по нескольким классам.
2. Создайте интерфейс состояний. Он должен описывать методы, общие для всех состояний,
	обнаруженных в контексте. Заметьте, что не всё поведение контекста нужно переносить в состояние,
	а только то, которое зависит от состояний.
3. Для каждого фактического состояния, создайте класс, реализующий интерфейс состояния.
	Переместите весь код, связанный с конкретным состоянием в нужный класс. В конце концов,
	все методы интерфейса состояния должны быть реализованы.
	При переносе поведения из контекста, вы можете столкнуться с тем, что это поведение зависит
	от приватных полей или методов контекста, к кторым нет доступа из
	состояния. Есть парочка способов обойти эту проблему.
	Самый простой — оставить поведение внутри контекста, вызывая его из объекта состояния.
	С другой стороны, вы может сделать классы состояний вложенными в класс контекста, и тогда они получат
	доступ ко всем приватным частям контекста. Но последний способ
	доступен только в некоторых языках программирования (например, Java, C#).
4. Создайте в контексте поле для хранения объектов- состояний, а также публичный метод для изменения значения этого поля.
5. Старые методы контекста, в которых находился зависимый от состояния код,
	замените на вызовы соответствующих методов объекта-состояния.
6. В зависимости от бизнес-логики, разместите код, который переключает состояние контекста либо внутри контекста,
	либо внутри классов конкретных состояний.


	Преимущества
	 Избавляет от множества больших условных операторов
	машины состояний.
	 Концентрирует в одном месте код, связанный с определённым состоянием.
	 Упрощает код контекста.
	 Недостатки
	 Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/

type PlayerI interface {
	next()
	previous()
	play()
	lock()
}

type pauseState struct {
	s State
}

func (p *pauseState) next() {
	p.s.idTrack++
	fmt.Println("трек остановлен ++ ")
}
func (p *pauseState) previous() {
	p.s.idTrack--
	fmt.Println("трек остановлен -- ")
}
func (p *pauseState) play() {
	p.s.off = false
	p.s.player = &playState{}
	fmt.Println("Play")
}
func (p *pauseState) lock() {
	p.s.player = &lockState{}
	p.s.off = false
	fmt.Println("lock")
}

type playState struct {
	s State
}

func (p *playState) next() {
	p.s.idTrack++
	fmt.Println("туц туц ")
}
func (p *playState) previous() {
	p.s.idTrack++
	fmt.Println("цут цут ")
}

func (p *playState) play() {
	p.s.off = true
	fmt.Println("stop")
	p.s.player = &pauseState{}
}
func (p *playState) lock() {
	p.s.off = true
	fmt.Println("lock")
	p.s.player = &lockState{}

}

type lockState struct {
	s State
}

func (p *lockState) next() {
	fmt.Println("...")
}
func (p *lockState) previous() {
	fmt.Println("...")
}
func (p *lockState) play() {
	fmt.Println("...")
}
func (p *lockState) lock() {
	p.s.player = &pauseState{}
	fmt.Println("unlock")
}

type State struct {
	player  PlayerI
	idTrack int
	off     bool
}

func main() {
	pp := State{}
	pp.player = &pauseState{}

	pp.player.play()
	fmt.Println(pp.off)
	pp.player.next()
	pp.player.lock()
	pp.player.next()
	pp.player.lock()
	fmt.Println(pp.idTrack)
	pp.player.play()
	pp.player.previous()
}
