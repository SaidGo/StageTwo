//go:build disable_extras



package legalentities

import (
    "context"
    "log"
    "example.com/local/Go2part/domain"
    cachepkg "example.com/local/Go2part/internal/cache"
	"context"
	"errors"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/internal/cache"
	"github.com/google/uuid"
)

// Минимальная in-memory реализация + опциональный Redis-кеш.
// Когда будет готово реальное хранилище/репозиторий — его можно
// подставить в методы вместо in-memory с сохранением кеш-логики.

type Service struct {
	items []domain.LegalEntity
	bank  []domain.BankAccount
	cache *cache.Service // может быть nil — тогда кеш просто не используется

	cache *cachepkg.Service
}

func NewService(c *cachepkg.Service) *Service { return &Service{cache: c} } }
func (s *Service) SetCache(c *cache.Service) { s.cache = c }

// --------------- ЮРЛИЦА ---------------

func (s *Service) GetAllLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
    if s.cache != nil {
        if items, err := s.cache.GetLegalEntities(ctx); err == nil {
            log.Printf("[legalentities] from cache (%d)", len(items))
            return items, nil
        }
    }
    // fallback: твоя "БД" (in-memory)
    items := s.items
    if s.cache != nil {
        s.cache.CacheLegalEntities(ctx, items)
    }
    return items, nil
}
	}
	// in-memory fallback
	out := make([]domain.LegalEntity, len(s.items))
	copy(out, s.items)

	if s.cache != nil {
		s.cache.CacheLegalEntities(ctx, out)
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
	if s.cache != nil {
		s.cache.ClearLegalEntities(ctx)
	}
	return in, nil
}

func (s *Service) UpdateLegalEntity(ctx context.Context, in domain.LegalEntity) (domain.LegalEntity, error) {
	if in.UUID == uuid.Nil {
		return domain.LegalEntity{}, errors.New("uuid is required")
	}
	for i := range s.items {
		if s.items[i].UUID == in.UUID {
			s.items[i] = in
			if s.cache != nil {
				s.cache.ClearLegalEntities(ctx)
			}
			return in, nil
		}
	}
	return domain.LegalEntity{}, errors.New("not found")
}

func (s *Service) DeleteLegalEntity(ctx context.Context, id domain.UUID) error {
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
	if s.cache != nil {
		s.cache.ClearLegalEntities(ctx)
	}
	return nil
}

// --------------- БАНКОВСКИЕ СЧЕТА ---------------

func (s *Service) GetAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	if s.cache != nil {
		if accs, err := s.cache.GetBankAccounts(ctx); err == nil {
			return accs, nil
		}
	}
	out := make([]domain.BankAccount, len(s.bank))
	copy(out, s.bank)

	if s.cache != nil {
		s.cache.CacheBankAccounts(ctx, out)
	}
	return out, nil
}

func (s *Service) CreateBankAccount(ctx context.Context, in domain.BankAccount) (domain.BankAccount, error) {
	if in.UUID == uuid.Nil {
		in.UUID = uuid.New()
	}
	s.bank = append(s.bank, in)
	if s.cache != nil {
		s.cache.ClearBankAccounts(ctx)
	}
	return in, nil
}

func (s *Service) UpdateBankAccount(ctx context.Context, in domain.BankAccount) (domain.BankAccount, error) {
	if in.UUID == uuid.Nil {
		return domain.BankAccount{}, errors.New("uuid is required")
	}
	for i := range s.bank {
		if s.bank[i].UUID == in.UUID {
			s.bank[i] = in
			if s.cache != nil {
				s.cache.ClearBankAccounts(ctx)
			}
			return in, nil
		}
	}
	return domain.BankAccount{}, errors.New("not found")
}

func (s *Service) DeleteBankAccount(ctx context.Context, id string) error {
	dst := s.bank[:0]
	var removed bool
	for _, a := range s.bank {
		if a.UUID.String() == id {
			removed = true
			continue
		}
		dst = append(dst, a)
	}
	if !removed {
		return errors.New("not found")
	}
	s.bank = dst
	if s.cache != nil {
		s.cache.ClearBankAccounts(ctx)
	}
	return nil
}
