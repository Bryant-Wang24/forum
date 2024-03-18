package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"example.com/gin_forum/config"
	"example.com/gin_forum/models"
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

const (
	POPULAR_TAGS_KEY = "popular_tags"
	USER_PROFILE_KEY = "user_profile_"
)

func GetPopularTags(ctx context.Context) ([]string, error) {
	cmd := rdb.ZRange(ctx, POPULAR_TAGS_KEY, 0, 10)
	return cmd.Result()
}

func GetUserProfile(ctx context.Context, userName string) (*models.User, error) {
	js, err := rdb.Get(ctx, USER_PROFILE_KEY+userName).Result()
	if err != nil {
		return nil, err
	}
	user := &models.User{}
	err = json.Unmarshal([]byte(js), user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func SetUserProfile(ctx context.Context, userName string, user *models.User, ttl int64) error {
	js, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := rdb.Set(ctx, USER_PROFILE_KEY+userName, string(js), time.Duration(ttl)*time.Second).Err(); err != nil {
		return err
	}
	return nil
}

func DeleteUserProfile(ctx context.Context, userName string) error {
	return rdb.Del(ctx, USER_PROFILE_KEY+userName).Err()
}
