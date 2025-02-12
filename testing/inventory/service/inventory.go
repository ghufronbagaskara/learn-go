package service

import "context"

// inventory service

// data store interface

type StoreManager interface {
	AddItem(ctx context.Context, itemId string, quantity int) error
	GetStock(ctx context.Context, itemId string) (numStock int, err error)
}

type Inventory struct {
	Store StoreManager
}

func NewInventoryService(store StoreManager) Inventory {
	return Inventory{
		Store: store,
	}
}

func (i Inventory) AddStock(ctx context.Context, itemId string, quantity int) error {
	// another complex bussiness logic reside here
	return i.Store.AddItem(ctx, itemId, quantity)
}

func (i Inventory) GetStock(ctx context.Context, itemId string, ) (int, error) {
	return i.Store.GetStock(ctx, itemId)
}