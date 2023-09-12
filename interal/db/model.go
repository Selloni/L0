package db

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Items   `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

type Delivery struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Payment struct {
	Id           int    `json:"id"`
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       int    `json:"amount"`
	PaymentDt    int    `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost int    `json:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total"`
	CustomFee    int    `json:"custom_fee"`
}

type Items struct {
	Id          int    `json:"id"`
	ChrtID      int    `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       int    `json:"price"`
	Rid         string `json:"rid"`
	Name        string `json:"name"`
	Sale        int    `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  int    `json:"total_price"`
	NmID        int    `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      int    `json:"status"`
}

func (o *Order) OpenFile(path *string) ([]byte, error) {
	data, err := os.ReadFile(*path)
	if err != nil {
		return nil, fmt.Errorf("could not open the file, err: %v", err)
	}
	return data, err
}

func (o *Order) ReadFile(data []byte) error {
	if err := json.Unmarshal(data, &o); err != nil {
		return err
	}
	return nil
}

//
//func (o *Order) GetAllOrders(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
//	for
//
//	if _, err := w.Write([]byte("Get all list")); err != nil {
//		log.Fatal(err)
//	}
//}
//
//func (o *Order) GetOrder(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
//	w.Write([]byte(o.DeliveryService))
//}
