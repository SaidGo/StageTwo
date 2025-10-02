package app

import (
	"time"

	"example.com/local/Go2part/internal/cache"
	"example.com/local/Go2part/pkg/redis"
)

func newCacheService(rds *redis.RDS) *cache.Service {
	// TTL - 5 минут
	return cache.NewService(rds, 5*time.Minute, 5*time.Minute)
}
