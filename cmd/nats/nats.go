package main

import (
	"L0/interal/db"
	"flag"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	nats, err := ConnectNATS()
	if err != nil {
		log.Fatal(err)
	}
	order := db.Order{}
	path := flag.String("json", "order.json", "path to file json")
	flag.Parse()
	log.Println(*path)
	data, err := order.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	if err = order.ReadFile(data); err != nil {
		log.Fatal(err)
	}
	nats.Publish("orders", data)
}

func ConnectNATS() (stan.Conn, error) {
	sc, err := stan.Connect("cluster", "client")
	if err != nil {
		return nil, fmt.Errorf("Expected to connect correctly, got err %v", err)
	}
	defer sc.Close()
	sub, err := sc.Subscribe("orders", func(msg *stan.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	}, stan.DurableName("i-will-remember"))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	return sc, nil
	//
	////// что то что бы запоминать переданное сообщение
	//handle := func(msg *stan.Msg) {
	//	// ...
	//}
	//
	//sub, err = sc.Subscribe(
	//	"stream-name",
	//	handle,
	//	stan.DurableName("i-will-remember"),
	//)
	//if err != nil {
	//	return nil, err
	//}
	/////////////////
}
