package inmemory

import (
	"L0/interal/db"
	"fmt"
)

type inMemory struct {
	store map[string]db.Order
}

func NewCash() inMemory {
	return inMemory{
		store: make(map[string]db.Order),
	}
}

func (m *inMemory) GetStore() *map[string]db.Order {
	return &m.store
}

func (m *inMemory) Add(order *db.Order) error {
	if _, ok := m.store[order.OrderUID]; !ok {
		m.store[order.OrderUID] = *order
	} else {
		return fmt.Errorf("such a key already exists : %s", order.OrderUID)
	}
	return nil
}

func (m *inMemory) Get(orderUID string) (db.Order, error) {
	if value, ok := m.store[orderUID]; ok {
		return value, nil
	}
	return db.Order{}, fmt.Errorf("there is no such key : %s", orderUID)
}

func (m *inMemory) Delete(orderUID string) error {
	if _, ok := m.store[orderUID]; !ok {
		return fmt.Errorf("there is no such key : %s", orderUID)
	}
	delete(m.store, orderUID)
	return nil
}
