package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/InsomniaCoder/go-redis-load/config"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

// Init Initialize the Redis instance
func Init() error {
	 config := config.GetConfig()
	redisCon := fmt.Sprintf("%v:%v", config.Redis.Host, config.Redis.Port)
	log.Printf("redis connection to: %v", redisCon)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisCon,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return nil
}

// Set a key/value
func Set(ctx context.Context, key string, data interface{}) error {
	err := redisClient.Set(ctx, key, data, 0).Err()
	return err
}