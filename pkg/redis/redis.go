package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"github.com/djwhocodes/go-grpc-microservices/pkg/config"
)

var Ctx = context.Background()

func NewRedis(cfg *config.Config) (*redis.Client, error) {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	err := client.Ping(Ctx).Err()

	if err != nil {
		return nil, err
	}

	return client, nil
}
