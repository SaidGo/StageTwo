package legalentities

import (
	"context"

	"github.com/google/uuid"

	"example.com/local/Go2part/domain"
)

type ServiceInterface interface {
	// LegalEntity
	GetAllLegalEntities(ctx context.Context) ([]*domain.LegalEntity, error)
	GetLegalEntity(ctx context.Context, id string) (*domain.LegalEntity, error)
	CreateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	UpdateLegalEntity(ctx context.Context, entity *domain.LegalEntity) error
	DeleteLegalEntity(ctx context.Context, id string) error

	// BankAccount
	GetAllBankAccounts(ctx context.Context, legalEntityUUID string) ([]*domain.BankAccount, error)
	CreateBankAccount(ctx context.Context, legalEntityUUID string, acc *domain.BankAccount) error
	UpdateBankAccount(ctx context.Context, legalEntityUUID string, acc *domain.BankAccount) error
	DeleteBankAccount(ctx context.Context, legalEntityUUID string, accountUUID string) error
	ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error)
}

type Service struct{ repo Repository }

func NewService(repo Repository) ServiceInterface { return &Service{repo: repo} }

// LegalEntity
func (s *Service) GetAllLegalEntities(ctx context.Context) ([]*domain.LegalEntity, error) {
	return s.repo.GetAll(ctx)
}
func (s *Service) GetLegalEntity(ctx context.Context, id string) (*domain.LegalEntity, error) {
	return s.repo.GetByID(ctx, id)
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
func (s *Service) DeleteLegalEntity(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

// BankAccount
func (s *Service) GetAllBankAccounts(ctx context.Context, legalEntityUUID string) ([]*domain.BankAccount, error) {
	return s.repo.GetAllBankAccounts(ctx, legalEntityUUID)
}
func (s *Service) CreateBankAccount(ctx context.Context, legalEntityUUID string, acc *domain.BankAccount) error {
	if acc.UUID == "" {
		acc.UUID = uuid.New().String()
	}
	acc.LegalEntityUUID = legalEntityUUID
	return s.repo.CreateBankAccount(ctx, acc)
}
func (s *Service) UpdateBankAccount(ctx context.Context, legalEntityUUID string, acc *domain.BankAccount) error {
	acc.LegalEntityUUID = legalEntityUUID
	return s.repo.UpdateBankAccount(ctx, acc)
}
func (s *Service) DeleteBankAccount(ctx context.Context, legalEntityUUID string, accountUUID string) error {
	return s.repo.DeleteBankAccount(ctx, legalEntityUUID, accountUUID)
}

// ListBankAccounts — возвращает список счетов через репозиторий.
func (s *Service) ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error) {
	return s.repo.ListBankAccounts(ctx, leUUID)
}
