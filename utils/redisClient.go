package utils

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	rdbClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
	_, err := rdbClient.Ping(context.TODO()).Result()
	if err != nil {
		log.Fatal(err)
	}
	return rdbClient
}
