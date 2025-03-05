package database

import (
	"context"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

func InitRedis() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URI"),
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping(context.TODO()).Result()

	if err != nil {
		log.Fatalf("Erro ao conectar no Redis: %v", err)
	}

	log.Println("Conectado ao Redis!")
}

func GetRedisConnection() *redis.Client {
	return redisClient
}
