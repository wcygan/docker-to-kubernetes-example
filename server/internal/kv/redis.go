package kv

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	Client *redis.Client
}

func NewRedisClient() *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-deployment:6379", // replace with your Redis server address
		Password: "",                      // replace with your password if any
		DB:       0,                       // use default DB
	})

	return &RedisClient{Client: rdb}
}

func (r *RedisClient) Put(key string, value string) error {
	ctx := context.Background()
	err := r.Client.Set(ctx, key, value, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to set key: %v", err)
	}
	return nil
}

func (r *RedisClient) Get(key string) (string, error) {
	ctx := context.Background()
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get key: %v", err)
	}
	return val, nil
}
