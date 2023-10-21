package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.

//https://www.kelche.co/blog/go/socket-programming/
*/

// храним адрес
type ConnectOption struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() error {
	co := ConnectOption{}
	if err := co.parsFlag(); err != nil {
		log.Fatal(err)
	}
	// создаем конетекст для разрыва соединения
	ctx, cancel := context.WithCancel(context.Background())
	// соедение с сервером, возвращем интерфейс для чтения и записи данных
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%s", co.Host, co.Port), co.Timeout)
	if err != nil {
		return err
	}
	if err != nil {
		return fmt.Errorf("не удалось подключитсья: %w", err)
	}
	// оправляем сообщение, в ответ считываем что пришло
	output := make(chan string)
	input := make(chan string)
	go readServer(conn, output)
	go sendData(conn, input)
	for {
		select {
		case data := <-input:
			if _, err := conn.Write([]byte(data)); err != nil {
				fmt.Printf("Writing error: %s\n", err)
				return nil
			}
		case data := <-output:
			fmt.Println(data)
		case <-ctx.Done():
			return nil
		}
	}
	defer cancel()
	return nil
}

func readServer(conn net.Conn, ch chan<- string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ch <- scanner.Text()
	}
}

func sendData(conn net.Conn, ch chan<- string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			// Stop by EOF (Ctrl + D)
			if err != io.EOF {
				log.Fatalln("cannot scan stdin")
			}
			break
		}
		ch <- text
	}
}
func (co *ConnectOption) parsFlag() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("не достаточно аргументов")
	}
	fl := flag.Duration("timeout", time.Duration(10)*time.Second, "timeout duration")
	flag.Parse()
	co.Timeout = *fl
	co.Host = flag.Arg(0)
	co.Port = flag.Arg(1)
	return nil
}
