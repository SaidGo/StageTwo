package legalentities

import (
	"example.com/local/Go2part/domain"
)

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Repository интерфейс
type Repository interface {
	// LegalEntity (2.1)
	Create(ctx context.Context, le *domain.LegalEntity) error
	GetAll(ctx context.Context) ([]*domain.LegalEntity, error)
	GetByID(ctx context.Context, id string) (*domain.LegalEntity, error)
	Update(ctx context.Context, le *domain.LegalEntity) error
	Delete(ctx context.Context, id string) error

	// BankAccount (2.2)
	GetAllBankAccounts(ctx context.Context, legalEntityUUID string) ([]*domain.BankAccount, error)
	CreateBankAccount(ctx context.Context, acc *domain.BankAccount) error
	UpdateBankAccount(ctx context.Context, acc *domain.BankAccount) error
	DeleteBankAccount(ctx context.Context, legalEntityUUID string, accountUUID string) error
	ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository { return &repository{db: db} }

// ==========================
// ========== MAPPERS =======
// ==========================

func ormLEToDomain(o *LegalEntity) (*domain.LegalEntity, error) {
	var accs []domain.BankAccount
	if len(o.BankAccountsJSON) > 0 {
		if err := json.Unmarshal(o.BankAccountsJSON, &accs); err != nil {
			return nil, err
		}
	}
	d := &domain.LegalEntity{
		UUID:         o.UUID,
		Name:         o.Name,
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
		BankAccounts: accs,
	}
	if o.CompanyUUID != nil {
		d.CompanyUUID = *o.CompanyUUID
	}
	return d, nil
}

func domainLEToOrm(d *domain.LegalEntity) (*LegalEntity, error) {
	baJSON := []byte("[]")
	if len(d.BankAccounts) > 0 {
		b, err := json.Marshal(d.BankAccounts)
		if err != nil {
			return nil, err
		}
		baJSON = b
	}
	var companyPtr *string
	if d.CompanyUUID != "" {
		v := d.CompanyUUID
		companyPtr = &v
	}
	return &LegalEntity{
		UUID:             d.UUID,
		Name:             d.Name,
		CompanyUUID:      companyPtr,
		BankAccountsJSON: baJSON,
		CreatedAt:        d.CreatedAt,
		UpdatedAt:        d.UpdatedAt,
	}, nil
}

func ormBAtoDomain(o *BankAccount) *domain.BankAccount {
	return &domain.BankAccount{
		UUID:            o.UUID,
		LegalEntityUUID: o.LegalEntityUUID,
		BIK:             o.BIK,
		Bank:            o.Bank,
		Address:         o.Address,
		CorrAccount:     o.CorrAccount,
		Account:         o.Account,
		Currency:        o.Currency,
		Comment:         o.Comment,
		IsPrimary:       o.IsPrimary,
		CreatedAt:       o.CreatedAt,
		UpdatedAt:       o.UpdatedAt,
		// DeletedAt пропускаем наружу как *time.Time по желанию
	}
}

func domainBAtoOrm(d *domain.BankAccount) *BankAccount {
	return &BankAccount{
		UUID:            d.UUID,
		LegalEntityUUID: d.LegalEntityUUID,
		BIK:             d.BIK,
		Bank:            d.Bank,
		Address:         d.Address,
		CorrAccount:     d.CorrAccount,
		Account:         d.Account,
		Currency:        d.Currency,
		Comment:         d.Comment,
		IsPrimary:       d.IsPrimary,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

// ==========================
// ====== LegalEntity =======
// ==========================

func (r *repository) Create(ctx context.Context, le *domain.LegalEntity) error {
	now := time.Now().UTC()
	if le.CreatedAt.IsZero() {
		le.CreatedAt = now
	}
	if le.UpdatedAt.IsZero() {
		le.UpdatedAt = now
	}
	orm, err := domainLEToOrm(le)
	if err != nil {
		return err
	}
	if err := r.db.WithContext(ctx).Create(orm).Error; err != nil {
		return err
	}
	le.CreatedAt = orm.CreatedAt
	le.UpdatedAt = orm.UpdatedAt
	return nil
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.LegalEntity, error) {
	var list []LegalEntity
	if err := r.db.WithContext(ctx).
		Where("deleted_at IS NULL").
		Find(&list).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.LegalEntity, 0, len(list))
	for i := range list {
		d, err := ormLEToDomain(&list[i])
		if err != nil {
			return nil, err
		}
		out = append(out, d)
	}
	return out, nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*domain.LegalEntity, error) {
	var o LegalEntity
	if err := r.db.WithContext(ctx).
		Where("uuid = ? AND deleted_at IS NULL", id).
		First(&o).Error; err != nil {
		return nil, err
	}
	return ormLEToDomain(&o)
}

func (r *repository) Update(ctx context.Context, le *domain.LegalEntity) error {
	baJSON := []byte("[]")
	if len(le.BankAccounts) > 0 {
		b, err := json.Marshal(le.BankAccounts)
		if err != nil {
			return err
		}
		baJSON = b
	}
	updates := map[string]any{
		"name":          le.Name,
		"bank_accounts": baJSON,
		"updated_at":    gorm.Expr("NOW()"),
	}
	// company_uuid nullable
	if le.CompanyUUID == "" {
		updates["company_uuid"] = gorm.Expr("NULL")
	} else {
		updates["company_uuid"] = le.CompanyUUID
	}
	return r.db.WithContext(ctx).
		Model(&LegalEntity{}).
		Where("uuid = ? AND deleted_at IS NULL", le.UUID).
		Updates(updates).Error
}

func (r *repository) Delete(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).
		Model(&LegalEntity{}).
		Where("uuid = ? AND deleted_at IS NULL", id).
		Update("deleted_at", gorm.Expr("NOW()")).Error
}

// ==========================
// ===== Bank Accounts ======
// ==========================

func (r *repository) GetAllBankAccounts(ctx context.Context, legalEntityUUID string) ([]*domain.BankAccount, error) {
	var rows []BankAccount
	if err := r.db.WithContext(ctx).
		Where("legal_entity_uuid = ?", legalEntityUUID).
		Order("is_primary DESC, created_at ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]*domain.BankAccount, 0, len(rows))
	for i := range rows {
		out = append(out, ormBAtoDomain(&rows[i]))
	}
	return out, nil
}

func (r *repository) CreateBankAccount(ctx context.Context, acc *domain.BankAccount) error {
	orm := domainBAtoOrm(acc)
	// защищаем инвариант: у аккаунта должен быть валидный владелец
	if orm.LegalEntityUUID == "" {
		return errors.New("legal_entity_uuid is empty")
	}
	// если is_primary=true — делаем его единственным primary у LE
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if orm.IsPrimary {
			if err := tx.Model(&BankAccount{}).
				Where("legal_entity_uuid = ?", orm.LegalEntityUUID).
				Update("is_primary", false).Error; err != nil {
				return err
			}
		}
		return tx.Create(orm).Error
	})
}

func (r *repository) UpdateBankAccount(ctx context.Context, acc *domain.BankAccount) error {
	if acc.UUID == "" {
		return errors.New("bank_account uuid is empty")
	}
	orm := domainBAtoOrm(acc)
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// is_primary: если включаем — выключаем остальные
		if orm.IsPrimary {
			if err := tx.Model(&BankAccount{}).
				Where("legal_entity_uuid = ? AND uuid <> ?", orm.LegalEntityUUID, orm.UUID).
				Update("is_primary", false).Error; err != nil {
				return err
			}
		}
		return tx.Model(&BankAccount{}).
			Where("uuid = ? AND legal_entity_uuid = ? AND deleted_at IS NULL", orm.UUID, orm.LegalEntityUUID).
			Updates(map[string]any{
				"bik":          orm.BIK,
				"bank":         orm.Bank,
				"address":      orm.Address,
				"corr_account": orm.CorrAccount,
				"account":      orm.Account,
				"currency":     orm.Currency,
				"comment":      orm.Comment,
				"is_primary":   orm.IsPrimary,
				"updated_at":   gorm.Expr("NOW()"),
			}).Error
	})
}

func (r *repository) DeleteBankAccount(ctx context.Context, legalEntityUUID string, accountUUID string) error {
	return r.db.WithContext(ctx).
		Model(&BankAccount{}).
		Where("uuid = ? AND legal_entity_uuid = ? AND deleted_at IS NULL", accountUUID, legalEntityUUID).
		Update("deleted_at", gorm.Expr("NOW()")).Error
}

// ==============
// = Индексы ====
// ==============

// EnsureIndexes — необязательная защита от отсутствующих индексов, выполнить один раз при старте (idempotent).
func (r *repository) EnsureIndexes(ctx context.Context) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_bank_accounts_legal_entity_uuid ON bank_accounts(legal_entity_uuid)`).Error; err != nil {
			return err
		}
		if err := tx.Exec(`CREATE INDEX IF NOT EXISTS idx_bank_accounts_deleted_at ON bank_accounts(deleted_at)`).Error; err != nil {
			return err
		}
		// На legal_entities.bank_accounts (jsonb) индекс опциональный, если нужен поиск по json.
		return nil
	})
}

// ==============
// = Preload ====
// ==============

// GetWithAccounts — утилита при необходимости получать LE вместе с аккаунтами (JOIN + Preload)
func (r *repository) GetWithAccounts(ctx context.Context, id string) (*domain.LegalEntity, error) {
	var o LegalEntity
	if err := r.db.WithContext(ctx).
		Preload(clause.Associations).
		Where("uuid = ? AND deleted_at IS NULL", id).
		First(&o).Error; err != nil {
		return nil, err
	}
	return ormLEToDomain(&o)
}

// ListBankAccounts возвращает все живые (не удалённые) счета юрлица из таблицы bank_accounts.
func (r *repository) ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error) {
	var rows []bankAccountRow
	if err := r.db.WithContext(ctx).
		Where("legal_entity_uuid = ?", leUUID).
		Order("is_primary DESC, created_at ASC").
		Find(&rows).Error; err != nil {
		return nil, err
	}

	out := make([]domain.BankAccount, 0, len(rows))
	for _, o := range rows {
		out = append(out, domain.BankAccount{
			UUID:            o.UUID,
			LegalEntityUUID: o.LegalEntityUUID,
			BIK:             o.BIK,
			Bank:            o.Bank,
			Address:         o.Address,
			CorrAccount:     o.CorrAccount,
			Account:         o.Account,
			Currency:        o.Currency,
			Comment:         o.Comment,
			IsPrimary:       o.IsPrimary,
			CreatedAt:       o.CreatedAt,
			UpdatedAt:       o.UpdatedAt,
			DeletedAt: func(d gorm.DeletedAt) *time.Time {
				if d.Valid {
					return &d.Time
				}
				return nil
			}(o.DeletedAt),
		})
	}
	return out, nil
}

// bankAccountRow — локальная ORM-модель таблицы bank_accounts
type bankAccountRow struct {
	UUID            string `gorm:"type:uuid;primaryKey"`
	LegalEntityUUID string `gorm:"type:uuid;not null;index"`
	BIK             string
	Bank            string
	Address         string
	CorrAccount     string
	Account         string `gorm:"not null"`
	Currency        string
	Comment         string
	IsPrimary       bool           `gorm:"not null;default:false"`
	CreatedAt       time.Time      `gorm:"not null"`
	UpdatedAt       time.Time      `gorm:"not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (bankAccountRow) TableName() string { return "bank_accounts" }
