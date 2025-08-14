package legalentities

import (
	"context"
	"time"

	"example.com/local/Go2part/domain"
	"github.com/google/uuid"
)

// Сервисный интерфейс.
type ServiceInterface interface {
	GetAllLegalEntities(ctx context.Context) ([]*domain.LegalEntity, error)
	GetLegalEntity(ctx context.Context, uuid string) (*domain.LegalEntity, error) // NEW
	CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	UpdateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	DeleteLegalEntity(ctx context.Context, uuid string) error
}

// Сервисная реализация.
type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

var _ ServiceInterface = (*Service)(nil)

func (s *Service) GetAllLegalEntities(ctx context.Context) ([]*domain.LegalEntity, error) {
	return s.repo.List(ctx)
}

func (s *Service) GetLegalEntity(ctx context.Context, id string) (*domain.LegalEntity, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error {
	if entity.UUID == "" {
		entity.UUID = uuid.New().String()
	}
	// Таймстемпы на стороне доменного сервиса
	now := time.Now()
	if entity.CreatedAt.IsZero() {
		entity.CreatedAt = now
	}
	entity.UpdatedAt = now
	return s.repo.Create(ctx, entity)
}

func (s *Service) UpdateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error {
	// Обновляем UpdatedAt при любом изменении
	entity.UpdatedAt = time.Now()
	return s.repo.Update(ctx, entity)
}

func (s *Service) DeleteLegalEntity(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
