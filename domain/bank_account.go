package domain

import (
	"time"

	"github.com/google/uuid"
)

type BankAccount struct {
	UUID            uuid.UUID `gorm:"type:char(36);primaryKey"`
	LegalEntityUUID uuid.UUID `gorm:"index"`
	BIK             string
	Bank            string
	Address         string
	CorrAccount     string
	Account         string
	Currency        string
	Comment         string
	IsPrimary       bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}
