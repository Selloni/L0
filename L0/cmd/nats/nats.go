package main

import (
	"L0/interal/db"
	"github.com/nats-io/stan.go"
	"log"
	"os"
)

func main() {
	myfile := os.Args[1:]
	order := db.Order{}
	for i := 0; i < len(myfile); i++ {
		log.Printf(myfile[i])
		data, err := order.OpenFile(&myfile[i])
		if err != nil {
			log.Fatal(err)
		}
		ConnectNATS(data)
	}
}

func ConnectNATS(data []byte) {
	sc, err := stan.Connect("test-cluster", "publisher", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Println(err)
	}
	defer sc.Close()
	sc.Publish("orders", data)

}
