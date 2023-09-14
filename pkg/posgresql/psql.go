package posgresql

import (
	"L0/interal/db"
	"L0/pkg/inmemory"
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
)

const IntoOrder = "INSERT INTO orders (order_uid, data) VALUES ($1, $2)"

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

func InsertOrder(pool *pgxpool.Pool, uid string, jsn []byte) error {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get connections: %v", err)
	}
	defer conn.Release()
	if _, err = conn.Exec(
		context.Background(), IntoOrder, uid, jsn,
	); err != nil {
		return fmt.Errorf("failed to transfer order to db :%v", err)
	}
	return nil
}

func GetOrder(pool *pgxpool.Pool, cash *inmemory.InMemory, order *db.Order) {
	ordUid, _ := pool.Query(context.Background(), fmt.Sprintf("select data from orders"))
	for ordUid.Next() {
		var jsn string
		ordUid.Scan(&jsn)
		order.ReadFile([]byte(jsn))
		log.Printf("wrote to cash uid: %s", order.OrderUID)
	}

}
