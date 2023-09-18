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

	cash := inmemory.NewCash()
	psql, err := posgresql.ConnectPsql()
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
			log.Println(err)
		} else {
			cash.Add(&order)
			if err != nil {
				log.Println(err)
			}
			err = posgresql.InsertOrder(psql, order.OrderUID, msg.Data)
			if err != nil {
				log.Print(err)
			}
			log.Printf("received via channel %s", order.OrderUID)
		}
	}, stan.DurableName("i-will-remember"))
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	posgresql.GetOrder(psql, &cash, &order)
	UpServer(&cash)
}

func UpServer(cash *inmemory.InMemory) {
	router := httprouter.New()
	handler := user.NewHandler(cash)
	handler.Register(router)

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
	log.Println("server is listening port 8080")
}
