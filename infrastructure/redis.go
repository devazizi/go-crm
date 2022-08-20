package infrastructure

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type RedisConnection struct {
	Connection        *redis.Client
	ContextBackGround context.Context
}

func NewRedis(dsn map[string]any) RedisConnection {
	ctx := context.Background()
	connection := redis.NewClient(&redis.Options{
		Addr:     dsn["addr"].(string),
		Password: dsn["password"].(string),
		DB:       dsn["db"].(int),
	})

	return RedisConnection{Connection: connection, ContextBackGround: ctx}
}
