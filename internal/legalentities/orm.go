package legalentities

import (
	"time"

	"gorm.io/gorm"
)

// LegalEntity — ORM-модель таблицы legal_entities.
// jsonb-колонка хранится как []byte (сырой JSON).
type LegalEntity struct {
	UUID             string         `gorm:"type:uuid;primaryKey;column:uuid"`
	Name             string         `gorm:"type:text;not null;column:name"`
	CompanyUUID      *string        `gorm:"type:uuid;column:company_uuid"`
	BankAccountsJSON []byte         `gorm:"type:jsonb;not null;default:'[]'::jsonb;column:bank_accounts"`
	CreatedAt        time.Time      `gorm:"not null;column:created_at"`
	UpdatedAt        time.Time      `gorm:"not null;column:updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (LegalEntity) TableName() string { return "legal_entities" }

// BankAccount — ORM-модель таблицы bank_accounts.
type BankAccount struct {
	UUID            string         `gorm:"type:uuid;primaryKey;default:gen_random_uuid();column:uuid"`
	LegalEntityUUID string         `gorm:"type:uuid;not null;index;column:legal_entity_uuid"`
	BIK             string         `gorm:"type:text;column:bik"`
	Bank            string         `gorm:"type:text;column:bank"`
	Address         string         `gorm:"type:text;column:address"`
	CorrAccount     string         `gorm:"type:text;column:corr_account"`
	Account         string         `gorm:"type:text;not null;column:account"`
	Currency        string         `gorm:"type:text;column:currency"`
	Comment         string         `gorm:"type:text;column:comment"`
	IsPrimary       bool           `gorm:"type:boolean;not null;default:false;column:is_primary"`
	CreatedAt       time.Time      `gorm:"not null;column:created_at"`
	UpdatedAt       time.Time      `gorm:"not null;column:updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index;column:deleted_at"`
}

func (BankAccount) TableName() string { return "bank_accounts" }
