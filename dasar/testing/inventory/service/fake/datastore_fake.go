package fake

import (
	"context"
	"errors"
)


type FakeStore struct {
	inventory map[string]int
}

func NewFakeStore() *FakeStore {
	return &FakeStore{
		inventory: make(map[string]int),
	}
}


func (f *FakeStore) AddItem(ctx context.Context, itemId string, quantity int) (err error) {
	if quantity < 0 {
		return errors.New("quantity must be greater than 0")
	}
	f.inventory[itemId] += quantity 
	return nil
}

func (f *FakeStore) GetStock(ctx context.Context, itemId string) (numStock int, err error) {
	numStock, exist := f.inventory[itemId]
	if !exist {
		return 0, errors.New("item not found")
	}
	return numStock, nil
}