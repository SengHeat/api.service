package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"api.drsb-purchase-service/config"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(cfg config.RedisConfig) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	return &Redis{client: client}
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := json.Marshal(value)

	if err != nil {
		return fmt.Errorf("failed to marshal value: %w", err)
	}

	return r.client.Set(ctx, key, data, expiration).Err()
}

func (r *Redis) Get(ctx context.Context, key string, dest interface{}) error {
	val, err := r.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return fmt.Errorf("key not found")
	}
	if err != nil {
		return fmt.Errorf("failed to get value: %w", err)
	}

	return json.Unmarshal([]byte(val), dest)
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

func (r *Redis) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.client.Exists(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return false, nil
	}

	return result > 0, err
}

func (r *Redis) Close() error {
	return r.client.Close()
}

func (r *Redis) Remember(ctx context.Context, key string, expiration time.Duration, dest interface{}, fetch func() (interface{}, error)) error {
	// 1. Try to get from cache
	err := r.Get(ctx, key, dest)
	if err == nil {
		return nil // Cache hit!
	}

	// 2. Cache miss: Fetch the data using the provided function
	data, err := fetch()
	if err != nil {
		return err
	}

	// 3. Store the newly fetched data in cache
	if err := r.Set(ctx, key, data, expiration); err != nil {
		return fmt.Errorf("failed to cache key %s: %w", key, err)
	}

	// 4. Marshal/Unmarshal to ensure 'dest' is populated with the fresh data
	// This is necessary because 'data' is the raw result from 'fetch'
	bytes, _ := json.Marshal(data)
	return json.Unmarshal(bytes, dest)
}
