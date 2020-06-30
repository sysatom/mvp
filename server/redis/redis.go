package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var RDB *redis.Client

func Connect() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6318",
		Password: "123456",
		DB:       0,
	})
	pong, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("redis", pong)
}
