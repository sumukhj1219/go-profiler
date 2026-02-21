package cache

import (
	"github.com/redis/go-redis/v9"
	"github.com/sumukhj1219/go-profiler/internals/config"
)

func NewRedisClient(cfg config.RedisConfig) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: "",
		DB:       0,
	})
}
