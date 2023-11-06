package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

// простой сервер прослушивающий порт 8080
func main() {
	// возвращяем интерфейс для приема подключений
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on port 8080")
	// метод для прослушивания входящих сообщений
	conn, err := ln.Accept()
	if err != nil {
		log.Fatal(err)
	}
	for {
		// считываем сообщение которое пришло
		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Message Received:", string(message))
		// переводим верхний регистр
		newmessage := strings.ToUpper(message)
		// отправляем сообщение
		conn.Write([]byte(newmessage + "\n"))
	}
}
