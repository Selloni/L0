package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func main() {
	nats, err := ConnectNATS()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(nats)
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
