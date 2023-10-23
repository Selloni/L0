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

	Реальные примеры использования паттерна "Состояние":
	1. Автомат продажи билетов: состояниями могут быть "ожидание ввода денег",
	"выбор места", "оплата", "печать билета".
	2. Редактор текста: состояниями могут быть "режим вставки",
	"режим выделения", "режим редактирования".
	3. Игра с различными уровнями сложности: состояниями могут быть "легкий уровень",
	"средний уровень", "сложный уровень".

*/

type PlayerI interface {
	next(context *State)
	previous(context *State)
	play(context *State)
	lock(context *State)
}

type pauseState struct {
	s State
}

func (p *pauseState) next(context *State) {
	context.idTrack++
	fmt.Println("трек остановлен ++ ")
}
func (p *pauseState) previous(context *State) {
	context.idTrack--
	fmt.Println("трек остановлен -- ")
}
func (p *pauseState) play(context *State) {
	context.setState(&playState{})
	context.off = true
	fmt.Println("Play")
}
func (p *pauseState) lock(context *State) {
	context.setState(&lockState{})
	context.off = false
	fmt.Println("lock")
}

type playState struct {
	s State
}

func (p *playState) next(context *State) {
	context.idTrack++
	fmt.Println("туц туц ")
}
func (p *playState) previous(context *State) {
	context.idTrack--
	fmt.Println("цут цут ")
}

func (p *playState) play(context *State) {
	context.setState(&pauseState{})
	context.off = true
	fmt.Println("stop")
}
func (p *playState) lock(context *State) {
	context.setState(&lockState{})
	context.off = true
	fmt.Println("lock")
}

type lockState struct {
	s State
}

func (p *lockState) next(context *State) {
	fmt.Println("...")
}
func (p *lockState) previous(context *State) {
	fmt.Println("...")
}
func (p *lockState) play(context *State) {
	fmt.Println("...")
}
func (p *lockState) lock(context *State) {
	context.setState(&pauseState{})
	fmt.Println("unlock")
}

type State struct {
	player  PlayerI
	idTrack int
	off     bool
}

func (s *State) setState(p PlayerI) {
	s.player = p
}

func (s *State) next() {
	s.player.next(s)
}
func (s *State) previous() {
	s.player.previous(s)
}
func (s *State) play() {
	s.player.play(s)
}
func (s *State) lock() {
	s.player.lock(s)
}

func main() {
	pp := State{}
	pp.setState(&pauseState{})
	pp.play()
	pp.next()
	pp.next()
	pp.next()
	fmt.Println(pp.idTrack)
	pp.play()
	fmt.Println(pp.off)
	pp.lock()
	pp.lock()
	pp.play()
	pp.previous()
	fmt.Println(pp.idTrack)
}
