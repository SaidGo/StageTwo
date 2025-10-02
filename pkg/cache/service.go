package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"

	"example.com/local/Go2part/domain"
)

const (
	keyLegalEntities           = "cache:legal_entities"
	keyBankAccountsGlobal      = "cache:bank_accounts"
	keyBankAccountsByEntityFmt = "cache:bank_accounts:%s"
)

type Service struct {
	rdb        *redis.Client
	ttl        time.Duration
	logMiss    bool
	logEnabled bool
}

var (
	defaultOnce sync.Once
	defaultSvc  *Service
)

func NewService() *Service {
	defaultOnce.Do(func() {
		addr := getenv("REDIS_ADDR", "127.0.0.1:6300")
		pass := getenv("REDIS_PASSWORD", "")
		db := getint("REDIS_DB", 0)
		ttlSec := getint("CACHE_TTL_SECONDS", 300)

		opts := &redis.Options{Addr: addr, DB: db}
		if pass != "" {
			opts.Password = pass
		}
		rdb := redis.NewClient(opts)
		defaultSvc = &Service{
			rdb:        rdb,
			ttl:        time.Duration(ttlSec) * time.Second,
			logMiss:    getbool("CACHE_LOG_MISS", false),
			logEnabled: true,
		}
	})
	return defaultSvc
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
func getint(k string, def int) int {
	if v := os.Getenv(k); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}
func getbool(k string, def bool) bool {
	if v := os.Getenv(k); v != "" {
		switch v {
		case "1", "true", "TRUE", "True", "yes", "YES":
			return true
		case "0", "false", "FALSE", "False", "no", "NO":
			return false
		}
	}
	return def
}

// -------- Legal Entities (global list) --------

func (s *Service) CacheLegalEntities(ctx context.Context, entities []domain.LegalEntity) {
	if s == nil || s.rdb == nil {
		return
	}
	b, err := json.Marshal(entities)
	if err != nil {
		log.Printf("cache: marshal legal_entities error: %v", err)
		return
	}
	if err := s.rdb.Set(ctx, keyLegalEntities, b, s.ttl).Err(); err != nil {
		log.Printf("cache: set legal_entities error: %v", err)
		return
	}
	if s.logEnabled {
		log.Printf("cache: SET %s ttl=%ds", keyLegalEntities, int(s.ttl.Seconds()))
	}
}

func (s *Service) GetLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	if s == nil || s.rdb == nil {
		return nil, redis.Nil
	}
	val, err := s.rdb.Get(ctx, keyLegalEntities).Bytes()
	if err != nil {
		if err != redis.Nil {
			log.Printf("cache: get legal_entities error: %v", err)
		} else if s.logMiss {
			log.Printf("cache: MISS %s -> DB", keyLegalEntities)
		}
		return nil, err
	}
	var out []domain.LegalEntity
	if err := json.Unmarshal(val, &out); err != nil {
		log.Printf("cache: unmarshal legal_entities error: %v", err)
		return nil, err
	}
	if s.logEnabled {
		log.Printf("cache: HIT %s", keyLegalEntities)
	}
	return out, nil
}

func (s *Service) ClearLegalEntities(ctx context.Context) {
	if s == nil || s.rdb == nil {
		return
	}
	if err := s.rdb.Del(ctx, keyLegalEntities).Err(); err != nil {
		log.Printf("cache: clear legal_entities error: %v", err)
		return
	}
	if s.logEnabled {
		log.Printf("cache: DEL %s", keyLegalEntities)
	}
}

// -------- Bank Accounts (global list) --------

func (s *Service) CacheBankAccounts(ctx context.Context, accounts []domain.BankAccount) {
	if s == nil || s.rdb == nil {
		return
	}
	b, err := json.Marshal(accounts)
	if err != nil {
		log.Printf("cache: marshal bank_accounts error: %v", err)
		return
	}
	if err := s.rdb.Set(ctx, keyBankAccountsGlobal, b, s.ttl).Err(); err != nil {
		log.Printf("cache: set bank_accounts error: %v", err)
		return
	}
	if s.logEnabled {
		log.Printf("cache: SET %s ttl=%ds", keyBankAccountsGlobal, int(s.ttl.Seconds()))
	}
}

func (s *Service) GetBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	if s == nil || s.rdb == nil {
		return nil, redis.Nil
	}
	val, err := s.rdb.Get(ctx, keyBankAccountsGlobal).Bytes()
	if err != nil {
		if err != redis.Nil {
			log.Printf("cache: get bank_accounts error: %v", err)
		} else if s.logMiss {
			log.Printf("cache: MISS %s -> DB", keyBankAccountsGlobal)
		}
		return nil, err
	}
	var out []domain.BankAccount
	if err := json.Unmarshal(val, &out); err != nil {
		log.Printf("cache: unmarshal bank_accounts error: %v", err)
		return nil, err
	}
	if s.logEnabled {
		log.Printf("cache: HIT %s", keyBankAccountsGlobal)
	}
	return out, nil
}

func (s *Service) ClearBankAccounts(ctx context.Context) {
	if s == nil || s.rdb == nil {
		return
	}
	if err := s.rdb.Del(ctx, keyBankAccountsGlobal).Err(); err != nil {
		log.Printf("cache: clear bank_accounts error: %v", err)
		return
	}
	if s.logEnabled {
		log.Printf("cache: DEL %s", keyBankAccountsGlobal)
	}
}

// -------- Bank Accounts by Legal Entity --------

func (s *Service) CacheBankAccountsByLegalEntity(ctx context.Context, leUUID string, accounts []domain.BankAccount) {
	if s == nil || s.rdb == nil {
		return
	}
	b, err := json.Marshal(accounts)
	if err != nil {
		log.Printf("cache: marshal bank_accounts[%s] error: %v", leUUID, err)
		return
	}
	key := fmt.Sprintf(keyBankAccountsByEntityFmt, leUUID)
	if err := s.rdb.Set(ctx, key, b, s.ttl).Err(); err != nil {
		log.Printf("cache: set %s error: %v", key, err)
		return
	}
	if s.logEnabled {
		log.Printf("cache: SET %s ttl=%ds", key, int(s.ttl.Seconds()))
	}
}

func (s *Service) GetBankAccountsByLegalEntity(ctx context.Context, leUUID string) ([]domain.BankAccount, error) {
	if s == nil || s.rdb == nil {
		return nil, redis.Nil
	}
	key := fmt.Sprintf(keyBankAccountsByEntityFmt, leUUID)
	val, err := s.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if err != redis.Nil {
			log.Printf("cache: get %s error: %v", key, err)
		} else if s.logMiss {
			log.Printf("cache: MISS %s -> DB", key)
		}
		return nil, err
	}
	var out []domain.BankAccount
	if err := json.Unmarshal(val, &out); err != nil {
		log.Printf("cache: unmarshal %s error: %v", key, err)
		return nil, err
	}
	if s.logEnabled {
		log.Printf("cache: HIT %s", key)
	}
	return out, nil
}

func (s *Service) ClearBankAccountsByLegalEntity(ctx context.Context, leUUID string) {
	if s == nil || s.rdb == nil {
		return
	}
	key := fmt.Sprintf(keyBankAccountsByEntityFmt, leUUID)
	if err := s.rdb.Del(ctx, key).Err(); err != nil {
		log.Printf("cache: clear %s error: %v", key, err)
		return
	}
	if s.logEnabled {
		log.Printf("cache: DEL %s", key)
	}
}
