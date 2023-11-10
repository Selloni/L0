package inmemory

import (
	"L0/interal/db"
)

type Cash interface {
	Add(order *db.Order) error
	Get(orderUID string) (db.Order, error)
	GetStore() *map[string]db.Order
}
