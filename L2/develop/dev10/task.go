package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
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

type ConnectOption struct {
	Host    string
	Port    string
	Timeout time.Duration
}

func main() {
	_, err := run()
	if err != nil {
		log.Fatal(err)
	}

}

func run() (net.Conn, error) {
	co := ConnectOption{}
	if err := co.parsFlag(); err != nil {
		log.Fatal(err)
	}
	dial := net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), co.Timeout)
	conn, err := dial.DialContext(ctx, "tcp", fmt.Sprintf("%s:%s", co.Host, co.Port))

	if err != nil {
		return nil, err
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(conn, text+"\n")
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message from server: " + message)
	}
	defer cancel()
	return conn, nil
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
