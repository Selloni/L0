package main

import (
	"L0/interal/db"
	"L0/pkg/posgresql"
	"flag"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
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
	//fmt.Println(order)
	psql, err := posgresql.ConnectPsql()
	if err != nil {
		log.Fatal(err)
	}
	err = posgresql.InsertOrder(psql, &order)
	if err != nil {
		log.Fatal(err)
	}
	//nats, err := ConnectNATS()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println(nats)
}

func ConnectNATS() (stan.Conn, error) {
	sc, err := stan.Connect("test-cluster", "test-client",
		stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		return nil, fmt.Errorf("Expected to connect correctly, got err %v", err)
	}
	defer sc.Close()
	sub, err := sc.Subscribe("orders", func(msg *stan.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	}, stan.DurableName("i-will-remember"))
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()
	return sc, nil
}
