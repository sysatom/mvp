package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
)

var ctx = context.Background()

var RDB *redis.Client

func Connect() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	pong, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("redis", pong)
}
