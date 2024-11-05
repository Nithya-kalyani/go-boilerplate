package cache

import (
	"context"
	"log"

	"github.com/Nithya-kalyani/go-boilerplate/pkg/config"
	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var ctx = context.Background()

func ConnectRedis(cfg *config.Config) {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.RedisHost + ":" + cfg.RedisPort,
		Password: cfg.RedisPassword,
		DB:       0, // default DB
	})

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis successfully.")

}

func SetCache(key string, value interface{}) error {
	err := RedisClient.Set(ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetCache(key string) (string, error) {
	val, err := RedisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
