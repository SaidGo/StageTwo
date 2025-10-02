package legalentities

import (
	"context"
	"errors"
	"log"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/internal/cache"
	"github.com/google/uuid"
)

type Service struct {
	items []domain.LegalEntity
	bank  []domain.BankAccount
	cache *cache.Service // может быть nil
}

func NewService(c *cache.Service) *Service { return &Service{cache: c} }

func (s *Service) GetAllLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	if s.cache != nil {
		if out, err := s.cache.GetLegalEntities(ctx); err == nil {
			log.Printf("[legalentities] from cache (%d)", len(out))
			return out, nil
		}
	}
	// fallback: in-memory
	out := make([]domain.LegalEntity, len(s.items))
	copy(out, s.items)

	if s.cache != nil {
		s.cache.CacheLegalEntities(ctx, out)
		log.Printf("[legalentities] cache miss -> stored (%d)", len(out))
	}
	return out, nil
}

func (s *Service) GetLegalEntity(ctx context.Context, id string) (domain.LegalEntity, error) {
	for _, e := range s.items {
		if e.UUID.String() == id {
			return e, nil
		}
	}
	return domain.LegalEntity{}, errors.New("not found")
}

func (s *Service) CreateLegalEntity(ctx context.Context, in domain.LegalEntity) (domain.LegalEntity, error) {
	if in.UUID == uuid.Nil {
		in.UUID = uuid.New()
	}
	s.items = append(s.items, in)
	s.cache.ClearLegalEntities(ctx)
	return in, nil
}

func (s *Service) UpdateLegalEntity(ctx context.Context, in domain.LegalEntity) (domain.LegalEntity, error) {
	if in.UUID == uuid.Nil {
		return domain.LegalEntity{}, errors.New("uuid is required")
	}
	for i := range s.items {
		if s.items[i].UUID == in.UUID {
			s.items[i] = in
			s.cache.ClearLegalEntities(ctx)
			return in, nil
		}
	}
	return domain.LegalEntity{}, errors.New("not found")
}

func (s *Service) DeleteLegalEntity(ctx context.Context, id string) error {
	dst := s.items[:0]
	var removed bool
	for _, e := range s.items {
		if e.UUID.String() == id {
			removed = true
			continue
		}
		dst = append(dst, e)
	}
	if !removed {
		return errors.New("not found")
	}
	s.items = dst
	s.cache.ClearLegalEntities(ctx)
	return nil
}

// -------- Банковские счета (задача 2.5) -----

func (s *Service) GetAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	if s.cache != nil {
		if out, err := s.cache.GetBankAccounts(ctx); err == nil {
			log.Printf("[bank-accounts] from cache (%d)", len(out))
			return out, nil
		}
	}
	out := make([]domain.BankAccount, len(s.bank))
	copy(out, s.bank)

	if s.cache != nil {
		s.cache.CacheBankAccounts(ctx, out)
		log.Printf("[bank-accounts] cache miss -> stored (%d)", len(out))
	}
	return out, nil
}
