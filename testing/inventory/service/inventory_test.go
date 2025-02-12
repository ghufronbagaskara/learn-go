package service_test

import (
	"context"
	"testing"
	"unit-test/inventory/service"
	"unit-test/inventory/service/fake"

	"github.com/stretchr/testify/assert"
)


func TestInventory_AddStock(t *testing.T)  {
	ctx := context.Background()
	
	fakeStore := fake.NewFakeStore()

	
	service := service.NewInventoryService(fakeStore)
	itemId := "item-1"
	initialStock := 10

	err := service.AddStock(ctx, itemId, initialStock)
	assert.NoError(t, err, "it should not return any error on the initial stage")

	currentStock, err := service.GetStock(ctx, itemId)
	assert.NoError(t, err, "it should not return any error")

	assert.Equal(t, 10, currentStock, "it should return 5, after stock added on the previous step")
}

