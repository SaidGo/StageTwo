package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

type RDS struct {
	Client *goredis.Client
}

func New(addr string, db int, password string) (*RDS, error) {
	rdb := goredis.NewClient(&goredis.Options{
		Addr:     addr,
		DB:       db,
		Password: password,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return &RDS{Client: rdb}, nil
}

// DSN формат: redis:<password>@<host>:<port>/<dbIndex>
// Пример: "redis:@127.0.0.1:6300/0" или "redis:pass@localhost:6379/1"
func FromDSN(dsn string) (*RDS, error) {
	if !strings.HasPrefix(dsn, "redis:") {
		return nil, fmt.Errorf("invalid redis dsn: %s", dsn)
	}
	raw := strings.TrimPrefix(dsn, "redis:")
	parts := strings.SplitN(raw, "@", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid redis dsn (no @): %s", dsn)
	}
	password := parts[0]
	hostdb := parts[1]

	addr := hostdb
	dbIdx := 0
	if i := strings.LastIndex(hostdb, "/"); i != -1 {
		addr = hostdb[:i]
		_ = parseInt(hostdb[i+1:], &dbIdx)
	}
	return New(addr, dbIdx, password)
}

func parseInt(s string, out *int) error {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		return err
	}
	*out = n
	return nil
}

// ---- strings ----

func (r *RDS) GetStr(ctx context.Context, key string) (string, error) {
	if r == nil || r.Client == nil {
		return "", errors.New("nil redis client")
	}
	return r.Client.Get(ctx, key).Result()
}

func (r *RDS) SetStr(ctx context.Context, key, val string, ttl time.Duration) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	return r.Client.Set(ctx, key, val, ttl).Err()
}

// ---- pubsub ----

func (r *RDS) Publish(ctx context.Context, channel, msg string) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	return r.Client.Publish(ctx, channel, msg).Err()
}

// ---- keys ----

func (r *RDS) Del(ctx context.Context, keys ...string) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	return r.Client.Del(ctx, keys...).Err()
}

func (r *RDS) Expire(ctx context.Context, key string, ttl time.Duration) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	return r.Client.Expire(ctx, key, ttl).Err()
}

func (r *RDS) Ping(ctx context.Context) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	return r.Client.Ping(ctx).Err()
}

// ---- hashes ----

// HSET(ctx, key, field1, val1, field2, val2, ...) — принимает любые типы, значения приводятся к строке.
func (r *RDS) HSET(ctx context.Context, key string, pairs ...interface{}) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	if len(pairs)%2 != 0 {
		return errors.New("HSET requires even number of pairs: field,value,...")
	}
	args := make([]any, 0, len(pairs))
	for _, v := range pairs {
		switch t := v.(type) {
		case string:
			args = append(args, t)
		case []byte:
			args = append(args, string(t))
		default:
			args = append(args, fmt.Sprint(t))
		}
	}
	return r.Client.HSet(ctx, key, args...).Err()
}

func (r *RDS) HGetAll(ctx context.Context, key string) (map[string]string, error) {
	if r == nil || r.Client == nil {
		return nil, errors.New("nil redis client")
	}
	return r.Client.HGetAll(ctx, key).Result()
}

// ---- sets ----

// SAddString(ctx, key, vals...) — принимает любые типы; элементы приводятся к строке.
func (r *RDS) SAddString(ctx context.Context, key string, vals ...interface{}) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	if len(vals) == 0 {
		return nil
	}
	args := make([]any, 0, len(vals))
	for _, v := range vals {
		switch t := v.(type) {
		case string:
			args = append(args, t)
		case []byte:
			args = append(args, string(t))
		default:
			args = append(args, fmt.Sprint(t))
		}
	}
	return r.Client.SAdd(ctx, key, args...).Err()
}

// SRemString(ctx, key, vals...) — принимает любые типы; элементы приводятся к строке.
func (r *RDS) SRemString(ctx context.Context, key string, vals ...interface{}) error {
	if r == nil || r.Client == nil {
		return errors.New("nil redis client")
	}
	if len(vals) == 0 {
		return nil
	}
	args := make([]any, 0, len(vals))
	for _, v := range vals {
		switch t := v.(type) {
		case string:
			args = append(args, t)
		case []byte:
			args = append(args, string(t))
		default:
			args = append(args, fmt.Sprint(t))
		}
	}
	return r.Client.SRem(ctx, key, args...).Err()
}

// SGet(ctx, key) -> members
func (r *RDS) SGet(ctx context.Context, key string) ([]string, error) {
	if r == nil || r.Client == nil {
		return nil, errors.New("nil redis client")
	}
	return r.Client.SMembers(ctx, key).Result()
}

// SIsMember(ctx, key, member)
func (r *RDS) SIsMember(ctx context.Context, key, member string) (bool, error) {
	if r == nil || r.Client == nil {
		return false, errors.New("nil redis client")
	}
	return r.Client.SIsMember(ctx, key, member).Result()
}
