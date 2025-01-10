package cart_test

import (
	"context"
	"errors"
	"testing"
	cart "unit-test/cart/service"
	"unit-test/cart/service/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestShoppingCart_AddItemToCart_Success(t *testing.T){
	//TODO: test AddItemToCart
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockRepositoryManager(ctrl)
	repositoryMock.EXPECT().AddToCart(context.Background(), "user-1", "product-a").Return(nil)
	
	shoppingCartService := cart.New(repositoryMock)
	err := shoppingCartService.AddItemToCart(context.Background(), "user-1", "product-a")
	assert.NoError(t, err, "it should not return any error")

}

func TestShoppingCart_AddItemToCart_ErrorOnRedist(t *testing.T){
	//TODO: test AddItemToCart
	ctrl := gomock.NewController(t)
	repositoryMock := mock.NewMockRepositoryManager(ctrl)
	repositoryMock.EXPECT().AddToCart(context.Background(), "user-2", "product-b").Return(errors.New("something failing on cache system"))
	
	shoppingCartService := cart.New(repositoryMock)
	err := shoppingCartService.AddItemToCart(context.Background(), "user-2", "product-b")
	assert.Error(t, err, "it should return any error")

}