package domain

import "time"

type BankAccount struct {
	UUID            string     `json:"uuid"`
	LegalEntityUUID string     `json:"legal_entity_uuid"`
	BIK             string     `json:"bik"`
	Bank            string     `json:"bank"`
	Address         string     `json:"address"`
	CorrAccount     string     `json:"corr_account"`
	Account         string     `json:"account"`
	Currency        string     `json:"currency"`
	Comment         string     `json:"comment"`
	IsPrimary       bool       `json:"is_primary"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at"`
	DeletedAt       *time.Time `json:"deleted_at,omitempty"`
}
