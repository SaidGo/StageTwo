package legalentities

import (
	"context"

	"example.com/local/Go2part/domain"
)

type ServiceInterface interface {
	CreateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error)
	ListAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error)
	GetBankAccount(ctx context.Context, uuid string) (domain.BankAccount, error)
	UpdateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error)
	DeleteBankAccount(ctx context.Context, uuid string) error

	// связь с юрлицом
	ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error)
}

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) CreateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error) {
	return s.repo.CreateBankAccount(ctx, ba)
}

func (s *Service) ListAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	return s.repo.ListAllBankAccounts(ctx)
}

func (s *Service) GetBankAccount(ctx context.Context, uuid string) (domain.BankAccount, error) {
	return s.repo.GetBankAccount(ctx, uuid)
}

func (s *Service) UpdateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error) {
	return s.repo.UpdateBankAccount(ctx, ba)
}

func (s *Service) DeleteBankAccount(ctx context.Context, uuid string) error {
	return s.repo.DeleteBankAccount(ctx, uuid)
}

func (s *Service) ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error) {
	return s.repo.ListBankAccounts(ctx, leUUID)
}
