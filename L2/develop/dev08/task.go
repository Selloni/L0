package main

import (
	"bufio"
	"fmt"
	"github.com/mitchellh/go-ps"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
Допп зажание : конвеер на пайпах

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
	// считывание с терминала
	if reader.Scan() {
		commandExecution(reader.Text())
	}
}

func commandExecution(str string) {
	command := strings.Split(str, " ")
	if command[0] == "pwd" {
		pwd()
		fmt.Print("\n")
	} else if command[0] == "cd" {
		cd(command)
	} else if command[0] == "echo" {
		echo(command)
		fmt.Print("\n")
	} else if command[0] == "kill" {
		if len(command) < 2 {
			fmt.Printf("kill: not enough arguments")
			return
		}
		kill(command[1:])
	} else if command[0] == "ps" {
		Fps()
	} else {
		fork(command)
	}
}

// PID процессов
func Fps() {
	fmt.Printf("%5s %-7s\n", "PID", "TTY")
	processes, err := ps.Processes()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, pp := range processes {
		fmt.Printf("%5d %-7s\n", pp.Pid(), pp.Executable())
	}
}

// поддержка стандартных команд
func fork(command []string) {
	cmd := exec.Command(command[0], command[1:]...)
	out, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

// удаление процесса
func kill(command []string) {
	for _, s := range command {
		i, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
		}
		proc, err := os.FindProcess(i)
		if err != nil {
			fmt.Println(err)
		}
		proc.Kill()
	}
}

// переход попапкам
func cd(command []string) {
	// пеерход в домашнюю деректорию
	if len(command) == 1 {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
		}
		err = os.Chdir(home)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// на деректорию ниже
	} else if command[1] == ".." {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		parent := filepath.Dir(wd)
		err = os.Chdir(parent)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		// переход в подпапку
		err := os.Chdir(command[1])
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

// где сейчас находимся
func pwd() {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Print(pwd)
}

// выводим сообщение
func echo(command []string) {
	if len(command) < 2 {
		fmt.Println(" ")
	}
	for i := 1; i < len(command); i++ {
		fmt.Printf("%s ", command[i])
	}
}
