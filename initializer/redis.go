package initializer

import (
	"context"
	"url_shortener/logger"
    "os"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	L *logger.Logger
}

func ConnectTORedis(r *RedisClient) *redis.Client {
    r.L.LogMessage("Entering into connecting redis")
    addr := os.Getenv("REDIS_HOST")
    if addr == "" {
        addr = "localhost"
    }
    password := os.Getenv("REDIS_PASSWORD")
    rdb := redis.NewClient(&redis.Options{
        Addr:     addr + ":6379",
        Password: password,
        DB:       0,
        Protocol: 2,
    })
    _, err := rdb.Ping(context.Background()).Result()
    if err != nil {
        r.L.LogFatalMessage("Failed to connect to Redis")
        return nil
    }
    r.L.LogMessage("Redis connected successfully")
    return rdb  
}