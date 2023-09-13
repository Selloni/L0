package main

import (
	"L0/interal/db"
	user "L0/interal/handlers"
	"L0/pkg/inmemory"
	"L0/pkg/posgresql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/stan.go"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {

	order := db.Order{}

	router := httprouter.New()
	log.Println("register user handler")

	//// nats
	sc, err := stan.Connect("test-cluster", "consumer",
		stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Printf("Expected to connect correctly, got err %v", err)
	}
	defer sc.Close()
	sub, err := sc.Subscribe("orders", func(msg *stan.Msg) {
		if err = order.ReadFile(msg.Data); err != nil {
			log.Fatal(err)
		}
		Run(order, router)
		log.Println(order.OrderUID)
	}, stan.DurableName("i-will-remember"))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()
	listener, ListenErr := net.Listen("tcp", fmt.Sprintf("127.0.0.1:8080"))
	if ListenErr != nil {
		log.Fatal(ListenErr)
	}
	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
	log.Println("server is listening port")
}

func Run(order db.Order, router *httprouter.Router) {

	cash := inmemory.NewCash()
	handler := user.NewHandler(&cash)
	handler.Register(router)

	cash.Add(&order)
	psql, err := posgresql.ConnectPsql()
	if err != nil {
		log.Fatal(err)
	}
	err = posgresql.InsertOrder(psql, &order)
	if err != nil {
		log.Fatal(err)
	}
}

func ConnectNATS() {
	sc, err := stan.Connect("test-cluster", "consumer",
		stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Printf("Expected to connect correctly, got err %v", err)
	}
	defer sc.Close()
	_, err = sc.Subscribe("orders", func(msg *stan.Msg) {
		log.Printf("Received message: %s", string(msg.Data))
	}, stan.DurableName("i-will-remember"))
	if err != nil {
		log.Fatal(err)
	}
}
