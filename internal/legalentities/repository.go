//go:build disable_extras

package legalentities

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"example.com/local/Go2part/domain"
	"example.com/local/Go2part/dto"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	// --- Bank Accounts ---
	CreateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error)
	ListAllBankAccounts(ctx context.Context) ([]domain.BankAccount, error)
	GetBankAccount(ctx context.Context, id string) (domain.BankAccount, error)
	UpdateBankAccount(ctx context.Context, ba *domain.BankAccount) (domain.BankAccount, error)
	DeleteBankAccount(ctx context.Context, id string) error
	ListBankAccounts(ctx context.Context, leUUID string) ([]domain.BankAccount, error)

	// --- Legal Entities (in-memory для фолбэка/демо) ---
	List(ctx context.Context) ([]dto.LegalEntity, error)
	Create(ctx context.Context, in dto.LegalEntityCreate) (dto.LegalEntity, error)
	Get(ctx context.Context, id uuid.UUID) (dto.LegalEntity, error)
	Update(ctx context.Context, id uuid.UUID, in dto.LegalEntityUpdate) (dto.LegalEntity, error)
	Delete(ctx context.Context, id uuid.UUID) error
	ListLegalEntities(ctx context.Context) ([]domain.LegalEntity, error)
	GetLegalEntity(ctx context.Context, id string) (domain.LegalEntity, error)
	UpdateLegalEntity(ctx context.Context, in *domain.LegalEntity) (domain.LegalEntity, error)
	DeleteLegalEntity(ctx context.Context, id string) error
}

type repository struct {
	db *gorm.DB

	leMu sync.RWMutex
	le   map[uuid.UUID]dto.LegalEntity
}

func NewRepository(db *gorm.DB) Repository {
	r := &repository{
		db: db,
		le: make(map[uuid.UUID]dto.LegalEntity),
	}
	r.ensureSchema()
	// ensure legal_entities for SQLite fallback
	if r.db != nil && r.db.Dialector.Name() == "sqlite" {
		if err := r.db.Exec(`CREATE TABLE IF NOT EXISTS legal_entities (
                        uuid TEXT PRIMARY KEY,
                        name TEXT,
                        company_uuid TEXT,
                        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
                        deleted_at DATETIME NULL
                );`).Error; err != nil {
			log.Printf("ensureSchema(sqlite): create table legal_entities failed: %v", err)
		}
	}
	return r
}

// ensureSchema — создаёт схему только для SQLite-фолбэка.
// Для Postgres используются миграции (migrations/*.sql), здесь — NOOP.
func (r *repository) ensureSchema() {
	if r.db == nil || r.db.Dialector.Name() != "sqlite" {
		return
	}
	// 1) Нормальный путь — AutoMigrate по модели bankAccountRow
	if err := r.db.AutoMigrate(&bankAccountRow{}); err == nil {
		return
	} else {
		log.Printf("ensureSchema(sqlite): AutoMigrate bank_accounts failed: %v", err)
	}
	// 2) Фолбэк: raw DDL (по одной команде на Exec; некоторые драйверы не принимают мульти-statement)
	if err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS bank_accounts (
	uuid TEXT PRIMARY KEY,
	legal_entity_uuid TEXT NULL,
	bik TEXT,
	bank TEXT,
	address TEXT,
	corr_account TEXT,
	account TEXT,
	currency TEXT,
	comment TEXT,
	is_primary INTEGER DEFAULT 0,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	deleted_at DATETIME NULL
);`).Error; err != nil {
		log.Printf("ensureSchema(sqlite): create table failed: %v", err)
		return
	}
	if err := r.db.Exec(`CREATE INDEX IF NOT EXISTS idx_bank_accounts_le ON bank_accounts(legal_entity_uuid);`).Error; err != nil {
		// ensure table legal_entities for SQLite fallback
		if err := r.db.AutoMigrate(&domain.LegalEntity{}); err != nil {
			// если AutoMigrate не сработал — создадим явным DDL
			if execErr := r.db.Exec(`
CREATE TABLE IF NOT EXISTS legal_entities (
  uuid TEXT PRIMARY KEY,
  name TEXT,
  company_uuid TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  deleted_at DATETIME NULL
);`).Error; execErr != nil {
				log.Printf("ensureSchema(sqlite): create table legal_entities failed: %v (automigrate err: %v)", execErr, err)
			}
		}
		log.Printf("ensureSchema(sqlite): create index failed: %v", err)
	}
}

// ---------------------------
// BankAccount mapping helpers
// ---------------------------

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

// ---------------------------
// BankAccount repository impl
// ---------------------------

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

// -----------------------------------------
// LegalEntity repository impl (in-memory)
// -----------------------------------------

func (r *repository) List(ctx context.Context) ([]dto.LegalEntity, error) {
	r.leMu.RLock()
	defer r.leMu.RUnlock()
	out := make([]dto.LegalEntity, 0, len(r.le))
	for _, v := range r.le {
		out = append(out, v)
	}
	return out, nil
}

func (r *repository) Create(ctx context.Context, in dto.LegalEntityCreate) (dto.LegalEntity, error) {
	id := uuid.New()
	le := dto.LegalEntity{
		UUID:    id.String(),
		Name:    in.Name,
		INN:     in.INN,
		KPP:     in.KPP,
		OGRN:    in.OGRN,
		Address: in.Address,
	}
	r.leMu.Lock()
	r.le[id] = le
	r.leMu.Unlock()
	return le, nil
}

func (r *repository) Get(ctx context.Context, id uuid.UUID) (dto.LegalEntity, error) {
	r.leMu.RLock()
	defer r.leMu.RUnlock()
	if v, ok := r.le[id]; ok {
		return v, nil
	}
	return dto.LegalEntity{}, errors.New("not found")
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, in dto.LegalEntityUpdate) (dto.LegalEntity, error) {
	r.leMu.Lock()
	defer r.leMu.Unlock()
	v, ok := r.le[id]
	if !ok {
		return dto.LegalEntity{}, errors.New("not found")
	}
	if in.Name != nil {
		v.Name = *in.Name
	}
	if in.INN != nil {
		v.INN = *in.INN
	}
	if in.KPP != nil {
		v.KPP = *in.KPP
	}
	if in.OGRN != nil {
		v.OGRN = *in.OGRN
	}
	if in.Address != nil {
		v.Address = *in.Address
	}
	r.le[id] = v
	return v, nil
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	r.leMu.Lock()
	defer r.leMu.Unlock()
	if _, ok := r.le[id]; !ok {
		return errors.New("not found")
	}
	delete(r.le, id)
	return nil
}

// --- autogenerated add: legal entity CRUD

func (r repository) ListLegalEntities(ctx context.Context) ([]domain.LegalEntity, error) {
	type leRow struct {
		UUID      string
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	var rows []leRow
	if err := r.db.WithContext(ctx).Table("legal_entities").Find(&rows).Error; err != nil {
		return nil, err
	}
	out := make([]domain.LegalEntity, 0, len(rows))
	for _, r0 := range rows {
		out = append(out, domain.LegalEntity{
			UUID:      uuid.MustParse(r0.UUID),
			Name:      r0.Name,
			CreatedAt: r0.CreatedAt,
			UpdatedAt: r0.UpdatedAt,
			DeletedAt: r0.DeletedAt,
		})
	}
	return out, nil
}

func (r repository) GetLegalEntity(ctx context.Context, id string) (domain.LegalEntity, error) {
	type leRow struct {
		UUID      string
		Name      string
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time
	}

	var r0 leRow
	tx := r.db.WithContext(ctx).Table("legal_entities").Where("uuid = ?", id).First(&r0)
	if tx.Error != nil {
		return domain.LegalEntity{}, tx.Error
	}
	return domain.LegalEntity{
		UUID:      uuid.MustParse(r0.UUID),
		Name:      r0.Name,
		CreatedAt: r0.CreatedAt,
		UpdatedAt: r0.UpdatedAt,
		DeletedAt: r0.DeletedAt,
	}, nil
}

func (r repository) UpdateLegalEntity(ctx context.Context, in *domain.LegalEntity) (domain.LegalEntity, error) {
	if err := r.db.WithContext(ctx).Table("legal_entities").Where("uuid = ?", in.UUID).Updates(in).Error; err != nil {
		return domain.LegalEntity{}, err
	}
	var out domain.LegalEntity
	if err := r.db.WithContext(ctx).Table("legal_entities").Where("uuid = ?", in.UUID).First(&out).Error; err != nil {
		return domain.LegalEntity{}, err
	}
	return out, nil
}

func (r repository) DeleteLegalEntity(ctx context.Context, id string) error {
	return r.db.WithContext(ctx).Table("legal_entities").Where("uuid = ?", id).Delete(&domain.LegalEntity{}).Error
}
