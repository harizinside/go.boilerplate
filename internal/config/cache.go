package config

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
)

type RedisConfig struct {
	Address  string
	Password string
	DB       int
}

func ConnectRedis(REDIS_URI string, REDIS_PASS string) error {
	cfg := RedisConfig{
		Address:  REDIS_URI,
		Password: REDIS_PASS,
		DB:       0,
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := redisClient.Ping(ctx).Err(); err != nil {
		return err
	}

	log.Println("Connected to Redis")
	return nil
}

func DisconnectRedis() {
	if redisClient == nil {
		return
	}

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := redisClient.Close(); err != nil {
		log.Printf("Failed to disconnect Redis: %v", err)
	} else {
		log.Println("Disconnected from Redis")
	}
}

func GetRedisClient() *redis.Client {
	return redisClient
}
