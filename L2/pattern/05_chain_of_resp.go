package main

import (
	"fmt"
	"time"
)

/*
	ЦЕПОЧКА ВЫЗОВОВ "Chain of Command"
	Цепочка обязанностей — это поведенческий паттерн проектирования,
	который позволяет передавать запросы последовательно по цепочке обработчиков.
	Каждый последующий обработчик решает, может ли он обработать
	запрос сам истоит ли передавать запрос дальше по цепи.

	=========================================================

	Преимущества
	 Уменьшает зависимость между клиентом и обработчиками.
	 Реализует принцип единственной обязанности.
	 Реализует принцип открытости/закрытости.
	Недостатки
	 Запрос может остаться никем не обработанным.
=======

*/

// интерфейс с переходу к следующему методу
// установка значения
type HandlerI interface {
	SetNext(handler HandlerI)
	HandlerI(data *Data)
}

// Обработка данных
type Data struct {
	GetSource    bool // были ли полученны данне
	UpdateSource bool // обработаны ли данные
}

type Device struct {
	Name string
	Next HandlerI // интерфейс
}

// происходит обработка данных
func (d *Device) HandlerI(data *Data) {
	if data.GetSource {
		fmt.Println("Данные уже были обработаны")
		d.Next.HandlerI(data)
		return
	}
	data.GetSource = true
	d.Next.HandlerI(data)
	for i := 0; i < 2; i++ {
		fmt.Println("Данные обрабатываються ...")
		time.Sleep(1 * time.Second)
	}
}

func (d *Device) SetNext(data HandlerI) {
	d.Next = data
}

// обновение данных
type UpdateData struct {
	Name string
	Next HandlerI // интерфейс
}

// происходит обработка данных
func (d *UpdateData) HandlerI(data *Data) {
	if data.GetSource {
		fmt.Println("Данные уже были обновленны")
		d.Next.HandlerI(data)
		return
	}
	data.GetSource = true
	d.Next.HandlerI(data)
	for i := 0; i < 2; i++ {
		fmt.Println("Обновление данных ...")
		time.Sleep(1 * time.Second)
	}
}

// передаем данные дальше
func (d *UpdateData) SetNext(data HandlerI) {
	d.Next = data
}

// Сохранение данных
type SaveData struct {
	Next HandlerI
}

func (d *SaveData) HandlerI(data *Data) {
	if data.GetSource {
		fmt.Println("Данные уже сохранены")
		return
	}
	data.GetSource = true
	d.Next.HandlerI(data)
	for i := 0; i < 2; i++ {
		fmt.Println("Сохроняем данные ...")
		time.Sleep(1 * time.Second)
	}
}

func (d *SaveData) SetNext(data HandlerI) {
	d.Next = data
}

func main() {
	dt := Data{}
	dv := Device{Name: "mySpoon"}
	ud := UpdateData{Name: "updateSpoon"}
	sd := SaveData{}
	dv.SetNext(&ud)
	ud.SetNext(&sd)
	dv.HandlerI(&dt)
}
