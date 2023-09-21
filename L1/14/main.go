//Разработать программу, которая в рантайме способна определить
//тип переменной: int, string, bool, channel
//из переменной типа interface{}.

package main

import "fmt"

func main() {
	var x interface{} = "1"
	// в операторе switch есть специальный синтаксис ,
	//который позволяет  определить тип
	switch expr := x.(type) {
	case int:
		fmt.Printf("int :%d", expr*2)
	case string:
		fmt.Printf("string: %s", expr)
	case bool:
		fmt.Printf("bool: %t", expr)
	case chan int:
		fmt.Println("chan int")
	case chan string:
		fmt.Println("chan string")
	case chan bool:
		fmt.Println("chan bool")
	default:
		fmt.Println("I dont know")
	}

}
