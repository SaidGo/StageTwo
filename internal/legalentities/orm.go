package legalentities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type bankAccountRow struct {
	UUID            uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	LegalEntityUUID *uuid.UUID `gorm:"type:uuid;index"`

	BIK         string `gorm:"type:varchar(16);not null;default:''"`
	Bank        string `gorm:"type:varchar(255);not null;default:''"`
	Address     string `gorm:"type:varchar(255);not null;default:''"`
	CorrAccount string `gorm:"type:varchar(64);not null;default:''"`
	Account     string `gorm:"type:varchar(64);not null;default:''"`
	Currency    string `gorm:"type:varchar(8);not null;default:''"`
	Comment     string `gorm:"type:text;not null;default:''"`
	IsPrimary   bool   `gorm:"not null;default:false"`

	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `gorm:"not null"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (bankAccountRow) TableName() string { return "bank_accounts" }
