package redis

import (
	"context"
	"log"
	"os"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Используем уже существующий type RDS из redis.go
// Здесь только создаём клиент из переменных окружения.
func NewFromEnv() *RDS {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "127.0.0.1:6300" // порт проброшен из Docker Desktop
	}
	pass := os.Getenv("REDIS_PASSWORD")

	c := goredis.NewClient(&goredis.Options{
		Addr:        addr,
		Password:    pass,
		DB:          0,
		DialTimeout: 3 * time.Second,
		ReadTimeout: 2 * time.Second,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := c.Ping(ctx).Err(); err != nil {
		log.Printf("[redis] ping failed: %v", err)
	} else {
		log.Printf("[redis] connected to %s", addr)
	}

	return &RDS{Client: c}
}
