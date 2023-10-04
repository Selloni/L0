package pattern

import "fmt"

/*
ФАСАД

Тип: Структурный

Фасад предоставляет простой интерфейс к сложной системе -
клиенту досточно пердоставить один метод,
для взаимодействия с подсистемами
/////
Преимущества
 Изолирует клиентов от компонентов системы.
 Уменьшает зависимость между подсистемой и клиентами.
Недостатки
 Фасад рискует стать божественным объектом,
привязанным ко всем классам программы.
//Вместо того, чтобы общаться друг с другом непосредственно,
//другие объекты полагаются на божественный объект.
//Так как на божественный объект ссылается так много кода,
//его обслуживание (внесение изменений) становится сложным:
//велик риск сломать существующую функциональность.
//////
Примеры применения
-Бибилиотеки и фраймворки
*/

type Read struct{}

func (r *Read) readFile(str string) {
	fmt.Printf("Открываем файл по пути %s\n считывем данные", str)
}

type Send struct{}

func (s *Send) pars() {
	fmt.Println("Парсим данные")
}

func (s *Send) sendData() {
	fmt.Println("Отправляем данные")
}

type Facade struct {
	read *Read
	send *Send
}

func (f Facade) SendData(str string) {
	facade := f.newFacade()
	facade.read.readFile(str)
	facade.send.pars()
	facade.send.sendData()
}

func (f Facade) newFacade() *Facade {
	return &Facade{
		read: &Read{},
		send: &Send{},
	}
}

func main() {
	f := Facade{}
	f.SendData("Home/dir/my_file")
}
