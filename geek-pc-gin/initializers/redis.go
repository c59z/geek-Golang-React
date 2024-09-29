package initializers

import (
	"geek-pc-gin/config"
	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{
		Addr: config.AppConfig.RedisAddr,
	})
}
