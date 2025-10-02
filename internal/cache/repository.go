//go:build disable_extras

package cache

import (
	"example.com/local/Go2part/pkg/redis"
)

type Repository struct {
	rds *redis.RDS
}

func NewRepository(rds *redis.RDS) *Repository {
	return &Repository{
		rds: rds,
	}
}
