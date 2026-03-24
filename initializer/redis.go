package initializer

import (
	"fmt"
	"url_shortener/logger"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	L *logger.Logger
}

func(r *RedisClient)ConnectTORedis() {
  r.L.LogMessage("Entering into connecting redis")
  rdb:=redis.NewClient(&redis.Options{
	    Addr:     "localhost:6379",
		Password: "", 
		DB:       0,  
		Protocol: 2,
  })
  fmt.Println("rdb is ",rdb)
  return
}