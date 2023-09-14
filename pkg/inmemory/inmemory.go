package inmemory

import (
	"L0/interal/db"
	"fmt"
)

type InMemory struct {
	store map[string]db.Order
}

func NewCash() InMemory {
	return InMemory{
		store: make(map[string]db.Order),
	}
}

func (m *InMemory) Add(order *db.Order) error {
	if _, ok := m.store[order.OrderUID]; !ok {
		m.store[order.OrderUID] = *order
	} else {
		return fmt.Errorf("such a key already exists : %s", order.OrderUID)
	}
	return nil
}

func (m *InMemory) Get(orderUID string) (db.Order, error) {
	if value, ok := m.store[orderUID]; ok {
		return value, nil
	}
	return db.Order{}, fmt.Errorf("there is no such key : %s", orderUID)
}
