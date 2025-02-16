package repository_test

import (
	"context"
	"testing"
	"unit-test/cart/repository"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	// connect to real caching system - redis
	redisOptions := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	redisClient := redis.NewClient(redisOptions)

	// instantiate cache repository
	repo := repository.New(redisClient)

	// execute repository methods
	gotErr := repo.AddToCartData(context.TODO(), "userID", "productID")

	// assert the result
	assert.NoError(t, gotErr, "it should not return error, redis container is there")

	// assert value of the stored cache
	gotResult, err := redisClient.HGet(context.Background(), "cart-userID", "name").Result()
	assert.NoError(t, err, "it should not return error")

	assert.Equal(t, "Sepatu Branded", gotResult, "it should return the expected cache product item name")
}
