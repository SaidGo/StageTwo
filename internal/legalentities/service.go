package legalentities

import (
	"context"

	"example.com/local/Go2part/domain"
	"github.com/google/uuid"
)

type ServiceInterface interface {
	GetAllLegalEntities(ctx context.Context) ([]*domain.LegalEntity, error)
	CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	UpdateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	DeleteLegalEntity(ctx context.Context, uuid string) error
}

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

func (s *Service) CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error {
	if entity.UUID == "" {
		entity.UUID = uuid.New().String()
	}
	return s.repo.Create(ctx, entity)
}

func (s *Service) UpdateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error {
	return s.repo.Update(ctx, entity)
}

func (s *Service) DeleteLegalEntity(ctx context.Context, uuid string) error {
	return s.repo.Delete(ctx, uuid)
}
