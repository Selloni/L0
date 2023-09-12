package main

import (
	"L0/interal/db"
	user "L0/interal/handlers"
	"L0/pkg/inmemory"
	"L0/pkg/posgresql"
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/nats-io/stan.go"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	path := flag.String("json", "order.json", "path to file json")
	flag.Parse()
	router := httprouter.New()
	log.Println("register user handler")
	handler := user.NewHandler()
	handler.Register(router)
	Run(path, router)
}

func Run(path *string, router *httprouter.Router) {
	order := db.Order{}
	cash := inmemory.NewCash()
	log.Println(*path)
	data, err := order.OpenFile(path)
	if err != nil {
		log.Fatal(err)
	}
	if err = order.ReadFile(data); err != nil {
		log.Fatal(err)
	}
	log.Println(order.OrderUID)
	cash.Add(&order)
	psql, err := posgresql.ConnectPsql()
	if err != nil {
		log.Fatal(err)
	}
	err = posgresql.InsertOrder(psql, &order)
	if err != nil {
		log.Fatal(err)
	}

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
