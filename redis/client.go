package redis

import (
	"context"
	rd "github.com/go-redis/redis/v8"
	"github.com/leaq-ru/parser/config"
	"github.com/leaq-ru/parser/logger"
	"time"
)

var Client *rd.Client

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	Client = rd.NewClient(&rd.Options{
		Addr: config.Env.Redis.URL,
	})

	err := Client.Ping(ctx).Err()
	logger.Must(err)
}
