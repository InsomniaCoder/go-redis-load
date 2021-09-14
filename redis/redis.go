package redis

import (
	"context"
	"fmt"

	"github.com/InsomniaCoder/go-redis-load/config"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// Init Initialize the Redis instance
func Init() error {
	 config := config.GetConfig()

	redisCon := fmt.Sprintf("%v:%v", config.Redis.Server, config.Redis.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisCon,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return nil
}

// Set a key/value
func Set(key string, data interface{}, ctx context.Context) error {
	err := redisClient.Set(ctx, key, data, 0).Err()
	return err
}