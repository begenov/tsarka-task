package redis

import (
	"fmt"

	"github.com/begenov/tsarka-task/internal/config"

	"github.com/go-redis/redis"
)

func NewRedisClient(cfg config.RedisConfig) (*redis.Client, error) {
	fmt.Println(cfg)
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
