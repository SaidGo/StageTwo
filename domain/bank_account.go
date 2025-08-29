package domain

import (
	"time"

	"github.com/google/uuid"
)

type BankAccount struct {
	UUID            uuid.UUID
	LegalEntityUUID uuid.UUID // uuid.Nil -> не привязан
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
