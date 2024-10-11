package redis

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	client        *redis.Client
	ctx           = context.Background()
	once          sync.Once
	redisPassword string
	redisHost     string
)

func initEnv() {
	redisPassword = os.Getenv("REDIS_PASSWORD")
	if redisPassword == "" {
		redisPassword = "long.nguyen"
	}

	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}
}

func initRedis() {
	once.Do(func() {
		initEnv()
		fmt.Printf("Connecting to Redis at %s %s", redisHost, redisPassword)
		client = redis.NewClient(&redis.Options{
			Addr:     redisHost,
			Password: redisPassword,
		})

		// Optionally, you can ping the Redis server to check the connection
		_, err := client.Ping(ctx).Result()
		if err != nil {
			panic("Failed to connect to Redis: " + err.Error())
		}
	})
}

func GetRedisClient() *redis.Client {
	initRedis()
	return client
}
