package legalentities

import (
	"context"
	"errors"
	"time"

	"example.com/local/Go2part/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error)
	ListAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error)
	GetBankAccount(ctx context.Context, id string) (domain.BankAccount, error)
	UpdateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error)
	DeleteBankAccount(ctx context.Context, id string) error
	ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error)
}

type repository struct{ db *gorm.DB }

func NewRepository(db *gorm.DB) Repository { return &repository{db: db} }

// mapping

func mapRowToDomain(r bankAccountRow) domain.BankAccount {
	var le uuid.UUID
	if r.LegalEntityUUID != nil {
		le = *r.LegalEntityUUID
	}
	var deletedAt *time.Time
	if r.DeletedAt.Valid {
		t := r.DeletedAt.Time
		deletedAt = &t
	}
	return domain.BankAccount{
		UUID:            r.UUID,
		LegalEntityUUID: le,
		BIK:             r.BIK,
		Bank:            r.Bank,
		Address:         r.Address,
		CorrAccount:     r.CorrAccount,
		Account:         r.Account,
		Currency:        r.Currency,
		Comment:         r.Comment,
		IsPrimary:       r.IsPrimary,
		CreatedAt:       r.CreatedAt,
		UpdatedAt:       r.UpdatedAt,
		DeletedAt:       deletedAt,
	}
}

func mapDomainToRow(d *domain.BankAccount) bankAccountRow {
	var le *uuid.UUID
	if d.LegalEntityUUID != uuid.Nil {
		tmp := d.LegalEntityUUID
		le = &tmp
	}
	return bankAccountRow{
		UUID:            d.UUID,
		LegalEntityUUID: le,
		BIK:             d.BIK,
		Bank:            d.Bank,
		Address:         d.Address,
		CorrAccount:     d.CorrAccount,
		Account:         d.Account,
		Currency:        d.Currency,
		Comment:         d.Comment,
		IsPrimary:       d.IsPrimary,
	}
}

// CRUD

func (r *repository) CreateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error) {
	row := mapDomainToRow(ba)
	if row.UUID == uuid.Nil {
		row.UUID = uuid.New()
	}
	if err := r.db.WithContext(ctx).Create(&row).Error; err != nil {
		return domain.BankAccount{}, err
	}
	var saved bankAccountRow
	if err := r.db.WithContext(ctx).Where("uuid = ?", row.UUID).First(&saved).Error; err != nil {
		return domain.BankAccount{}, err
	}
	return mapRowToDomain(saved), nil
}

func (r *repository) ListAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error) {
	var rows []bankAccountRow
	if err := r.db.WithContext(ctx).Where("deleted_at IS NULL").Order("is_primary DESC, created_at ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]domain.BankAccount, 0, len(rows))
	for _, rr := range rows {
		out = append(out, mapRowToDomain(rr))
	}
	return out, nil
}

func (r *repository) GetBankAccount(ctx context.Context, id string) (domain.BankAccount, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return domain.BankAccount{}, err
	}
	var row bankAccountRow
	if err := r.db.WithContext(ctx).Where("uuid = ? AND deleted_at IS NULL", uid).First(&row).Error; err != nil {
		return domain.BankAccount{}, err
	}
	return mapRowToDomain(row), nil
}

func (r *repository) UpdateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error) {
	var existing bankAccountRow
	if err := r.db.WithContext(ctx).Where("uuid = ? AND deleted_at IS NULL", ba.UUID).First(&existing).Error; err != nil {
		return domain.BankAccount{}, err
	}
	var le *uuid.UUID
	if ba.LegalEntityUUID != uuid.Nil {
		tmp := ba.LegalEntityUUID
		le = &tmp
	}
	updates := map[string]any{
		"legal_entity_uuid": le,
		"bik":               ba.BIK,
		"bank":              ba.Bank,
		"address":           ba.Address,
		"corr_account":      ba.CorrAccount,
		"account":           ba.Account,
		"currency":          ba.Currency,
		"comment":           ba.Comment,
		"is_primary":        ba.IsPrimary,
	}
	if err := r.db.WithContext(ctx).Model(&bankAccountRow{}).Where("uuid = ? AND deleted_at IS NULL", ba.UUID).Updates(updates).Error; err != nil {
		return domain.BankAccount{}, err
	}
	if err := r.db.WithContext(ctx).Where("uuid = ?", ba.UUID).First(&existing).Error; err != nil {
		return domain.BankAccount{}, err
	}
	return mapRowToDomain(existing), nil
}

func (r *repository) DeleteBankAccount(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}
	res := r.db.WithContext(ctx).Where("uuid = ? AND deleted_at IS NULL", uid).Delete(&bankAccountRow{})
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("not found")
	}
	return nil
}

func (r *repository) ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error) {
	le, err := uuid.Parse(leUUID)
	if err != nil {
		return nil, err
	}
	var rows []bankAccountRow
	if err := r.db.WithContext(ctx).Where("legal_entity_uuid = ? AND deleted_at IS NULL", le).Order("is_primary DESC, created_at ASC").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]domain.BankAccount, 0, len(rows))
	for _, rr := range rows {
		out = append(out, mapRowToDomain(rr))
	}
	return out, nil
}
