package configs

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	// БД
	DBDriver string // "postgres" | "sqlite"
	DBDSN    string

	// Redis
	RedisDSN string // формат: redis:<password>@<host>:<port>/<dbIndex>
	RedisTTL time.Duration

	// Kafka (по необходимости)
	KafkaBrokers string // "localhost:29092"
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func parseDurationEnv(key string, def time.Duration) time.Duration {
	if v := os.Getenv(key); v != "" {
		if d, err := time.ParseDuration(v); err == nil {
			return d
		}
		// поддержка числовых секунд
		if n, err := strconv.Atoi(v); err == nil {
			return time.Duration(n) * time.Second
		}
	}
	return def
}

func Load() Config {
	return Config{
		DBDriver:     getenv("DB_DRIVER", "sqlite"),
		DBDSN:        getenv("DB_DSN", "file:go2part.db?_fk=1"),
		RedisDSN:     getenv("REDIS_DSN", "redis:@127.0.0.1:6300/0"),
		RedisTTL:     parseDurationEnv("REDIS_TTL", 300*time.Second),
		KafkaBrokers: getenv("KAFKA_BROKERS", "localhost:29092"),
	}
}
