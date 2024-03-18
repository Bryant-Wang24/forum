package cache

import (
	"context"
	"fmt"

	"example.com/gin_forum/config"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: config.GetRedisAddr(),
		//Password: "", // no password set
		//DB:       0,  // use default DB
	})

	ping := rdb.Ping(context.Background())
	if err := ping.Err(); err != nil {
		panic(err)
	}
	fmt.Println("Redis connected successfully")
}
