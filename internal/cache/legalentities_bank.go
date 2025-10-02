package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/pkg/redis"
)

const (
	keyLegalEntities = "cache:legal_entities"
	keyBankAccounts  = "cache:bank_accounts"
)

// Минимальный сервис кеша (если уже есть Service в этом пакете — дополним его полями).
type Service struct {
	rds                  *redis.RDS
	cacheLegalTTL        time.Duration
	cacheBankAccountsTTL time.Duration
}

func NewService(rds *redis.RDS, legalTTL, bankTTL time.Duration) *Service {
	return &Service{
		rds:                  rds,
		cacheLegalTTL:        legalTTL,
		cacheBankAccountsTTL: bankTTL,
	}
}

// ---- Юрлица ----

func (s *Service) CacheLegalEntities(ctx context.Context, entities []domain.LegalEntity) {
	if s == nil || s.rds == nil {
		return
	}
	b, err := json.Marshal(entities)
	if err != nil {
		return
	}
	_ = s.rds.SetStr(ctx, keyLegalEntities, string(b), s.cacheLegalTTL)
}

func (s *Service) GetLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	if s == nil || s.rds == nil {
		return nil, errors.New("cache not configured")
	}
	raw, err := s.rds.GetStr(ctx, keyLegalEntities)
	if err != nil || raw == "" {
		return nil, errors.New("cache miss")
	}
	var out []domain.LegalEntity
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, errors.New("cache miss")
	}
	return out, nil
}

func (s *Service) ClearLegalEntities(ctx context.Context) {
	if s == nil || s.rds == nil {
		return
	}
	_ = s.rds.Del(ctx, keyLegalEntities)
}

// ---- Банковские счета ----

func (s *Service) CacheBankAccounts(ctx context.Context, accounts []domain.BankAccount) {
	if s == nil || s.rds == nil {
		return
	}
	b, err := json.Marshal(accounts)
	if err != nil {
		return
	}
	_ = s.rds.SetStr(ctx, keyBankAccounts, string(b), s.cacheBankAccountsTTL)
}

func (s *Service) GetBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	if s == nil || s.rds == nil {
		return nil, errors.New("cache not configured")
	}
	raw, err := s.rds.GetStr(ctx, keyBankAccounts)
	if err != nil || raw == "" {
		return nil, errors.New("cache miss")
	}
	var out []domain.BankAccount
	if err := json.Unmarshal([]byte(raw), &out); err != nil {
		return nil, errors.New("cache miss")
	}
	return out, nil
}

func (s *Service) ClearBankAccounts(ctx context.Context) {
	if s == nil || s.rds == nil {
		return
	}
	_ = s.rds.Del(ctx, keyBankAccounts)
}
