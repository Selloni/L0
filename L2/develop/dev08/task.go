package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	for {
		shell()
	}
}

func shell() {
	pwd()
	fmt.Print(": ")
	reader := bufio.NewScanner(os.Stdin)
	if reader.Scan() {
		//multCommand := strings.Split(reader.Text(), "|")
		//for i := range multCommand {
		//	commandExecution(multCommand[i])
		//}
		commandExecution(reader.Text())
	}
}

func commandExecution(str string) {
	//str = strings.TrimSpace(str)
	command := strings.Split(str, " ")
	if command[0] == "pwd" {
		pwd()
		fmt.Print("\n")
	} else if command[0] == "cd" {

	} else if command[0] == "echo" {
		echo(command)
		fmt.Print("\n")
	} else if command[0] == "kill" {

	} else if command[0] == "ps" {

	} else {
		fmt.Printf("Команда не найдена: %s\n", command[0])
	}
}

func cd() {

}

func pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Print(pwd)
}

func echo(command []string) {
	if len(command) < 2 {
		fmt.Println(" ")
	}
	for i := 1; i < len(command); i++ {
		fmt.Printf("%s ", command[i])
	}
}
