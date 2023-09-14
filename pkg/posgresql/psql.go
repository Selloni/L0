package posgresql

import (
	"L0/interal/db"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

const QuerOrder = "INSERT INTO orders " +
	"(order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) " +
	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
const QuerDelivery = "INSERT INTO delivery " +
	"(order_uid,name, phone, zip, city, address, region, email) " +
	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)"
const QuerPayment = "INSERT INTO payment " +
	"(order_uid,transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) " +
	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
const QuerItems = "INSERT INTO items " +
	"(order_uid, chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) " +
	"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"

type ConfigPsql struct {
	Name, Pass, Host, Port, Database string
}

func ConnectPsql() (*pgxpool.Pool, error) {
	cnf := ConfigPsql{
		Name:     "grandpat",
		Pass:     "grandpat",
		Host:     "localhost",
		Port:     "5432",
		Database: "postgres",
	}
	psqlClient, err := NewClient(context.Background(), cnf)
	if err != nil {
		return nil, err
	}
	return psqlClient, nil
}

func NewClient(ctx context.Context, con ConfigPsql) (*pgxpool.Pool, error) {
	var pool *pgxpool.Pool
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", con.Name, con.Pass, con.Host,
		con.Port, con.Database)
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}
	pool, err = pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	return pool, nil
}

func InsertOrder(pool *pgxpool.Pool, data *db.Order) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get connections: %v", err)
	}
	defer conn.Release()
	if _, err = conn.Exec(context.Background(), QuerOrder,
		data.OrderUID, data.TrackNumber, data.Entry, data.Locale,
		data.InternalSignature, data.CustomerID, data.DeliveryService,
		data.Shardkey, data.SmID, data.DateCreated, data.OofShard,
	); err != nil {
		return fmt.Errorf("failed to transfer order to db :%v", err)
	}
	delivery := data.Delivery
	if _, err = conn.Exec(context.Background(), QuerDelivery, data.OrderUID,
		delivery.Name, delivery.Phone, delivery.Zip, delivery.City,
		delivery.Address, delivery.Region, delivery.Email,
	); err != nil {
		return fmt.Errorf("delivery: %v", err)
	}
	payment := data.Payment
	if _, err = conn.Exec(context.Background(), QuerPayment, data.OrderUID,
		payment.Transaction, payment.RequestID, payment.Currency,
		payment.Provider, payment.Amount, payment.PaymentDt, payment.Bank,
		payment.DeliveryCost, payment.GoodsTotal, payment.CustomFee,
	); err != nil {
		return fmt.Errorf("payment: %v", err)
	}
	item := data.Items
	for i := 0; i < len(item); i++ {
		if _, err = conn.Exec(context.Background(), QuerItems, data.OrderUID,
			item[i].ChrtID, item[i].TrackNumber, item[i].Price,
			item[i].Rid, item[i].Name, item[i].Sale, item[i].Size,
			item[i].TotalPrice, item[i].NmID, item[i].Brand, item[i].Status,
		); err != nil {
			return fmt.Errorf("item[%d]: %v", i, err)
		}
	}
	return nil
}

func GetOrder(pool *pgxpool.Pool, cash *db.Order, uid string) {
	order, err := pool.Query(context.Background(), fmt.Sprintf("SELECT * FROM orders WHERE order_uid = %s", uid))
	if err != nil {
		log.Printf("fatal get data out BD :", err)
	}
	defer order.Close()
	for order.Next() {
		err := order.Scan(&cash.OrderUID, &cash.TrackNumber, &cash.Entry, &cash.Locale,
			&cash.InternalSignature, &cash.CustomerID, &cash.DeliveryService,
			&cash.Shardkey, &cash.SmID, &cash.DateCreated, &cash.OofShard)
		if err != nil {
			log.Print(err)
		}
	}
}
