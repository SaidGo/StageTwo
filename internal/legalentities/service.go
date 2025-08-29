package legalentities

import (
	"context"

	"github.com/google/uuid"

	"example.com/local/Go2part/dto"
)

type ServiceInterface interface {
	ListLegalEntities(ctx context.Context) ([]dto.LegalEntity, error)
	CreateLegalEntity(ctx context.Context, in dto.LegalEntityCreate) (dto.LegalEntity, error)
	GetLegalEntity(ctx context.Context, id uuid.UUID) (dto.LegalEntity, error)
	UpdateLegalEntity(ctx context.Context, id uuid.UUID, in dto.LegalEntityUpdate) (dto.LegalEntity, error)
	DeleteLegalEntity(ctx context.Context, id uuid.UUID) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) ListLegalEntities(ctx context.Context) ([]dto.LegalEntity, error) {
	return s.repo.List(ctx)
}

func (s *Service) CreateLegalEntity(ctx context.Context, in dto.LegalEntityCreate) (dto.LegalEntity, error) {
	return s.repo.Create(ctx, in)
}

func (s *Service) GetLegalEntity(ctx context.Context, id uuid.UUID) (dto.LegalEntity, error) {
	return s.repo.Get(ctx, id)
}

func (s *Service) UpdateLegalEntity(ctx context.Context, id uuid.UUID, in dto.LegalEntityUpdate) (dto.LegalEntity, error) {
	return s.repo.Update(ctx, id, in)
}

func (s *Service) DeleteLegalEntity(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
