package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

var RDB *redis.Client

func Connect(options *redis.Options) {
	client := redis.NewClient(options)
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis连接失败")
		return
	}
	RDB = client
}
