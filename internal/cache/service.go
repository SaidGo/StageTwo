//go:build disable_extras

package cache

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strconv"
	"time"

	"example.com/local/Go2part/domain"
	rds "example.com/local/Go2part/pkg/redis"
)

const (
	keyLegalEntities = "cache:legal_entities"
	keyBankAccounts  = "cache:bank_accounts"
)

type Service struct {
	rds *rds.RDS
	ttl time.Duration
}

func NewService(r *rds.RDS) *Service {
	ttl := 5 * time.Minute
	if v := os.Getenv("CACHE_TTL_SECONDS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			ttl = time.Duration(n) * time.Second
		}
	}
	return &Service{rds: r, ttl: ttl}
}

// ===== Legal entities =====

func (s *Service) CacheLegalEntities(ctx context.Context, entities []domain.LegalEntity) {
	if s == nil || s.rds == nil || s.rds.Client == nil {
		return
	}
	b, _ := json.Marshal(entities)
	if err := s.rds.Client.Set(ctx, keyLegalEntities, b, s.ttl).Err(); err != nil {
		log.Printf("[cache] CacheLegalEntities set err: %v", err)
	}
}

func (s *Service) GetLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	if s == nil || s.rds == nil || s.rds.Client == nil {
		return nil, context.Canceled
	}
	b, err := s.rds.Client.Get(ctx, keyLegalEntities).Bytes()
	if err != nil {
		log.Printf("[cache] GetLegalEntities miss: %v", err)
		return nil, err
	}
	var out []domain.LegalEntity
	if err := json.Unmarshal(b, &out); err != nil {
		log.Printf("[cache] GetLegalEntities unmarshal err: %v", err)
		return nil, err
	}
	log.Printf("[cache] GetLegalEntities hit (%d items)", len(out))
	return out, nil
}

func (s *Service) ClearLegalEntities(ctx context.Context) {
	if s == nil || s.rds == nil || s.rds.Client == nil {
		return
	}
	if err := s.rds.Client.Del(ctx, keyLegalEntities).Err(); err != nil {
		log.Printf("[cache] ClearLegalEntities err: %v", err)
	}
}

// ===== Bank accounts =====

func (s *Service) CacheBankAccounts(ctx context.Context, accounts []domain.BankAccount) {
	if s == nil || s.rds == nil || s.rds.Client == nil {
		return
	}
	b, _ := json.Marshal(accounts)
	if err := s.rds.Client.Set(ctx, keyBankAccounts, b, s.ttl).Err(); err != nil {
		log.Printf("[cache] CacheBankAccounts set err: %v", err)
	}
}

func (s *Service) GetBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	if s == nil || s.rds == nil || s.rds.Client == nil {
		return nil, context.Canceled
	}
	b, err := s.rds.Client.Get(ctx, keyBankAccounts).Bytes()
	if err != nil {
		log.Printf("[cache] GetBankAccounts miss: %v", err)
		return nil, err
	}
	var out []domain.BankAccount
	if err := json.Unmarshal(b, &out); err != nil {
		log.Printf("[cache] GetBankAccounts unmarshal err: %v", err)
		return nil, err
	}
	log.Printf("[cache] GetBankAccounts hit (%d items)", len(out))
	return out, nil
}

func (s *Service) ClearBankAccounts(ctx context.Context) {
	if s == nil || s.rds == nil || s.rds.Client == nil {
		return
	}
	if err := s.rds.Client.Del(ctx, keyBankAccounts).Err(); err != nil {
		log.Printf("[cache] ClearBankAccounts err: %v", err)
	}
}
