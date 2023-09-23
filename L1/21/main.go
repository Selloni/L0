package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

// Патерн адаптер, суть, патерн позваоляет
// не совместимым пакетам рабоать в одином ключе
// скрывая всю работу от пользователя

// внешний сервис, для работы с xml, который работает только с xml
type Analytical interface {
	SendXml(bb []byte) interface{}
}

type XmlDock struct {
	Name string ` xml:"name"`
	Age  int    ` xml:"age"`
}

func (doc XmlDock) SendXml(bb []byte) {
	fmt.Println(string(bb))
}

////////// наш сервис //////////

// наш сервис работате только c json
type JsonDock struct {
	Name string `json:"name" xml:"name"`
	Age  int    `json:"age" xml:"age"`
}

func (doc JsonDock) ConvertToXml(path string) { // конвертируем в xml
	//var doX XmlDock
	file, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(file, &doc)
	fmt.Println(doc)
	bytes, err := xml.Marshal(doc)
	if err != nil {
		fmt.Println(err)
	}
	var xx XmlDock
	xx.SendXml(bytes)
}

type JsonDockAdapter struct { // адаптер
	jsD *JsonDock
}

// сервисы между собой не связанны
func (adapter *JsonDockAdapter) SendXml(file string) {
	fmt.Println("Отправка xml")
	adapter.jsD.ConvertToXml(file)
}

func main() {
	var adapter JsonDockAdapter
	adapter.SendXml("OD.json")
}
