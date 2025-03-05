package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	redisClient *redis.Client
}

func NewRedisCache(redisClient *redis.Client) *RedisCache {
	return &RedisCache{redisClient: redisClient}
}

func (cache *RedisCache) SetCache(key string, value string, expiration time.Duration) error {
	return cache.redisClient.Set(context.TODO(), key, value, expiration).Err()
}

func (cache *RedisCache) GetCache(key string) (string, error) {
	return cache.redisClient.Get(context.TODO(), key).Result()
}
